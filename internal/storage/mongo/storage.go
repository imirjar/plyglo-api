package mongo

import (
	"context"
	"fmt"

	// Добавили для форматирования ошибок
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	uri    string
	client *mongo.Client
	db     *mongo.Database
}

func New(dbconn string) *Storage {
	return &Storage{uri: dbconn}
}

func (s *Storage) Conn(ctx context.Context) error {
	var err error
	s.client, err = mongo.Connect(options.Client().ApplyURI(s.uri))
	if err != nil {
		return fmt.Errorf("mongo connect: %w", err)
	}
	if err := s.client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("mongo ping: %w", err)
	}
	s.db = s.client.Database("plyglo")
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	if s.client == nil {
		return nil
	}
	return s.client.Disconnect(ctx)
}
