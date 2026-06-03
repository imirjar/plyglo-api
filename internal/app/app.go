package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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

	serverErrors := make(chan error, 1)

	go func() {
		if err := app.server.Run(); err != nil && err != http.ErrServerClosed {
			serverErrors <- fmt.Errorf("server error: %w", err)
		}
	}()

	select {
	case err := <-serverErrors:
		app.storage.Close()
		return fmt.Errorf("server failed: %w", err)

	case <-ctx.Done():
		log.Println("Starting graceful shutdown...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		log.Println("Step 1: Stopping HTTP server...")
		if err := app.server.Stop(shutdownCtx); err != nil {
			app.storage.Close()
			return fmt.Errorf("stop server: %w", err)
		}

		log.Println("Step 2: Closing storage connections...")
		app.storage.Close()
		log.Println("Application gracefully stopped")
		return nil
	}
}
