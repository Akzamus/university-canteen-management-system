package app

import (
	"context"
	"fmt"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

type App struct {
	serviceProvider *serviceProvider
	config          config.AppConfig
	httpRouter      *chi.Mux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	var err error
	a.config, err = config.NewAppConfig()
	return err
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.config)
	return nil
}

func (a *App) initHttpServer(_ context.Context) error {
	a.httpRouter = chi.NewRouter()
	timeoutSeconds := a.serviceProvider.AppConfig().Server.HttpTimeoutSeconds

	a.httpRouter.Use(middleware.RequestID)
	a.httpRouter.Use(middleware.RealIP)
	a.httpRouter.Use(middleware.Logger)
	a.httpRouter.Use(middleware.Recoverer)
	a.httpRouter.Use(middleware.Timeout(time.Duration(timeoutSeconds) * time.Second))

	a.serviceProvider.UserHandler().RegisterRoutes(a.httpRouter)

	return nil
}

func (a *App) runHttpServer() error {
	port := a.serviceProvider.AppConfig().Server.HttpPort
	log.Printf("Starting http server on port: %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a.httpRouter)
}
