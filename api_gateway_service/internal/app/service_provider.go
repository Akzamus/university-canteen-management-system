package app

import (
	"fmt"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/config"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	userService "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service/user"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler"
	userHandler "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler/user"
)

type serviceProvider struct {
	appConfig   config.AppConfig
	userService service.UserService
	userHandler handler.UserHandler
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
		s.userHandler = userHandler.NewHandler(s.UserService())
	}
	return s.userHandler
}
