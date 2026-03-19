package main

import (
	"context"
	"log"

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

	ctx := context.Background()
	log.Fatal(app.Start(ctx))
}
