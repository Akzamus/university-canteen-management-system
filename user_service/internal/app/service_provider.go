package app

import (
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/config"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository/psql"
	userRepository "github.com/Akzamus/university-canteen-management-system/user_service/internal/repository/psql/user"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/service"
	userService "github.com/Akzamus/university-canteen-management-system/user_service/internal/service/user"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/transport/http/handler"
	userHandler "github.com/Akzamus/university-canteen-management-system/user_service/internal/transport/http/handler/user"
	"github.com/jmoiron/sqlx"
	"log"
)

type serviceProvider struct {
	appConfig      config.AppConfig
	db             *sqlx.DB
	userRepository repository.UserRepository
	userService    service.UserService
	userHandler    handler.UserHandler
}

func newServiceProvider(cfg config.AppConfig) *serviceProvider {
	return &serviceProvider{appConfig: cfg}
}

func (s *serviceProvider) AppConfig() config.AppConfig {
	return s.appConfig
}

func (s *serviceProvider) Database() *sqlx.DB {
	if s.db == nil {
		db, err := psql.NewClient(s.AppConfig().Database)
		if err != nil {
			log.Fatalf("failed to create database client: %v", err)
		}
		s.db = db
	}
	return s.db
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.Database())
	}
	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository())
	}
	return s.userService
}

func (s *serviceProvider) UserHandler() handler.UserHandler {
	if s.userHandler == nil {
		s.userHandler = userHandler.NewHandler(s.UserService())
	}
	return s.userHandler
}
