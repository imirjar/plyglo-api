package mongo

import (
	"context"
	"fmt" // Добавили для форматирования ошибок

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	client *mongo.Client
}

func New(ctx context.Context, dbconn string) (*Storage, func(context.Context) error) {
	// 1. Подключаемся к клиенту
	client, err := mongo.Connect(options.Client().ApplyURI(dbconn))
	if err != nil {
		// Вместо тихого возврата nil выбрасываем панику с деталями
		panic(fmt.Sprintf("failed to connect to mongo: %v", err))
	}

	// 2. Проверяем доступность базы данных
	if err = client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		panic(fmt.Sprintf("failed to ping mongo: %v", err))
	}

	storage := &Storage{
		client: client,
	}

	// 3. Безопасно возвращаем функцию закрытия напрямую от клиента
	closeFn := func(closeCtx context.Context) error {
		return client.Disconnect(closeCtx)
	}

	return storage, closeFn
}

func (s *Storage) Close(ctx context.Context) error {
	if s.client == nil {
		return nil
	}
	return s.client.Disconnect(ctx)
}
