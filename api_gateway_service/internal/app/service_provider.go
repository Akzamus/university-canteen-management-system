package app

import (
	"fmt"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/config"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	authService "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service/auth"
	userService "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service/user"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler"
	authHandler "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler/auth"
	userHandler "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler/user"
	"github.com/go-chi/jwtauth"
)

type serviceProvider struct {
	appConfig   config.AppConfig
	userService service.UserService
	userHandler handler.UserHandler
	authService service.AuthService
	authHandler handler.AuthHandler
	jwtAuth     *jwtauth.JWTAuth
}

func newServiceProvider(cfg config.AppConfig) *serviceProvider {
	return &serviceProvider{appConfig: cfg}
}

func (s *serviceProvider) AppConfig() config.AppConfig {
	return s.appConfig
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		userServiceBaseUrl := fmt.Sprintf(
			"http://%s:%d",
			s.AppConfig().UserService.Host,
			s.AppConfig().UserService.Port,
		)
		s.userService = userService.NewService(userServiceBaseUrl)
	}
	return s.userService
}

func (s *serviceProvider) UserHandler() handler.UserHandler {
	if s.userHandler == nil {
		s.userHandler = userHandler.NewHandler(s.UserService(), s.JwtAuth())
	}
	return s.userHandler
}

func (s *serviceProvider) JwtAuth() *jwtauth.JWTAuth {
	if s.jwtAuth == nil {
		s.jwtAuth = jwtauth.New("HS256", []byte(s.AppConfig().JwtSecretKey), nil)
	}
	return s.jwtAuth
}

func (s *serviceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.JwtAuth(), s.UserService())
	}
	return s.authService
}

func (s *serviceProvider) AuthHandler() handler.AuthHandler {
	if s.authHandler == nil {
		s.authHandler = authHandler.NewHandler(s.AuthService())
	}
	return s.authHandler
}
