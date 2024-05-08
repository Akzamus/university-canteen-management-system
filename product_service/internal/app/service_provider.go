package app

import (
	"log"

	"github.com/Akzamus/university-canteen-management-system/product_service/internal/config"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/repository"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/repository/psql"
	ProductRepository "github.com/Akzamus/university-canteen-management-system/product_service/internal/repository/psql/user"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/service"
	ProductService "github.com/Akzamus/university-canteen-management-system/product_service/internal/service/user"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/transport/http/handler"
	ProductHandler "github.com/Akzamus/university-canteen-management-system/product_service/internal/transport/http/handler/user"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	appConfig         config.AppConfig
	db                *sqlx.DB
	ProductRepository repository.ProductRepository
	ProductService    service.ProductService
	ProductHandler    handler.ProductHandler
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

func (s *serviceProvider) ProductRepository() repository.ProductRepository {
	if s.ProductRepository == nil {
		s.ProductRepository = ProductRepository.NewRepository(s.Database())
	}
	return s.ProductRepository
}

func (s *serviceProvider) ProductService() service.ProductService {
	if s.ProductService == nil {
		s.ProductService = ProductService.NewService(s.ProductRepository())
	}
	return s.ProductService
}

func (s *serviceProvider) ProductHandler() handler.ProductHandler {
	if s.ProductHandler == nil {
		s.ProductHandler = ProductHandler.NewHandler(s.ProductService())
	}
	return s.ProductHandler
}
