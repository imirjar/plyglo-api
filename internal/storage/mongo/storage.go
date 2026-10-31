package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	client *mongo.Client
}

func New(ctx context.Context, dbconn string) *Storage {
	storage := &Storage{}

	client, err := mongo.Connect(options.Client().ApplyURI(dbconn))
	if err != nil {
		return nil
	}

	if err = client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil
	}

	storage.client = client
	return storage
}

func (s *Storage) Close(ctx context.Context) error {
	if s.client == nil {
		return nil
	}
	return s.client.Disconnect(ctx)
}
