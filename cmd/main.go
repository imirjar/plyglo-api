package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/imirjar/poliglotim-api/internal/app"

	"github.com/imirjar/poliglotim-api/docs"
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
	docs.SwaggerInfo.Title = "Language course API"

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("App started")

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("App stopped gracefully")
}
