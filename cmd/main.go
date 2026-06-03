package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/imirjar/poliglotim-api/internal/app"
)

// @title Language course API
// @version 1.0
// @description API for the educational platform

// @contact.name Artem Zadorov
// @contact.email azadorov1234@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := app.New()

	serverErrors := make(chan error, 1)

	go func() {
		if err := app.Start(ctx); err != nil {
			serverErrors <- fmt.Errorf("start app: %w", err)
		}
	}()

	select {
	case err := <-serverErrors:
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if stopErr := app.Stop(shutdownCtx); stopErr != nil {
			log.Printf("stop app after error: %v", stopErr)
		}

		log.Fatal(err)

	case <-ctx.Done():
		log.Println("Starting graceful shutdown...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.Stop(shutdownCtx); err != nil {
			log.Fatalf("stop app: %v", err)
		}

		log.Println("Application gracefully stopped")
	}
}
