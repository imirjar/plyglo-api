package main

import (
	"context"
	"log"

	"github.com/imirjar/poliglotim-api/internal/app"

	"github.com/imirjar/poliglotim-api/docs"
)

// @title Poliglotim API
// @version 1.0
// @description API for the Poliglotim educational platform

// @contact.name API Support
// @contact.email support@poliglotim.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	docs.SwaggerInfo.Title = "Poliglotim API"

	ctx := context.Background()
	log.Fatal(app.Start(ctx))
}
