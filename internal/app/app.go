package app

import (
	"context"
	"log"

	// "github.com/alexliesenfeld/health"

	"github.com/imirjar/poliglotim-api/config"
	gw "github.com/imirjar/poliglotim-api/internal/gateway/http"
	"github.com/imirjar/poliglotim-api/internal/service"
	"github.com/imirjar/poliglotim-api/internal/storage"
)

func Start(ctx context.Context) error {
	config := config.New(ctx)

	storage, err := storage.New(ctx, config.DBConn)
	if err != nil {
		panic(err)
	}
	defer storage.Disconnect(ctx)

	service := service.New(ctx)
	srv := gw.New(ctx, config.Port)

	service.Storage = storage
	srv.Service = service

	log.Printf("Starting server on the port %s... \n", config.Port)
	return srv.Run(ctx)

}
