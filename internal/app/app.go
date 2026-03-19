package app

import (
	"context"
	"log"

	"github.com/imirjar/poliglotim-api/config"
	srv "github.com/imirjar/poliglotim-api/internal/app/http"
	"github.com/imirjar/poliglotim-api/internal/service"
	"github.com/imirjar/poliglotim-api/internal/storage"
)

func Start(ctx context.Context) error {
	config := config.New()

	storage := storage.New(ctx)
	storage.Connect(ctx, config.DBConn)
	defer storage.Disconnect(ctx)

	service := service.New()
	srv := srv.New(config.Port)

	service.Storage = storage
	srv.Service = service

	log.Printf("Starting server on the port %s... \n", config.Port)
	return srv.Run()

}
