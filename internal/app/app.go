package app

import (
	"context"
	"fmt"
	"net/http"

	// "github.com/alexliesenfeld/health"

	"github.com/imirjar/poliglotim-api/config"
	server "github.com/imirjar/poliglotim-api/internal/gateway/http"
	service "github.com/imirjar/poliglotim-api/internal/service/study"
	"github.com/imirjar/poliglotim-api/internal/storage"
)

type App struct {
	server  Server
	storage Storage
}

type Server interface {
	Run() error
	Stop(context.Context) error
}

type Storage interface {
	Conn(context.Context) error
	Close()
}

func New() *App {
	config := config.New()

	storage := storage.New(storage.WithDB(config.DBConn))
	service := service.New(service.WithStorage(storage))
	server := server.New(server.WithServer(config.Port), server.WithService(service))

	return &App{
		server:  server,
		storage: storage,
	}
}

func (app *App) Start(ctx context.Context) error {
	if err := app.storage.Conn(ctx); err != nil {
		return fmt.Errorf("connect storage: %w", err)
	}

	if err := app.server.Run(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	if err := app.server.Stop(ctx); err != nil {
		app.storage.Close()
		return fmt.Errorf("stop server: %w", err)
	}

	app.storage.Close()
	return nil
}
