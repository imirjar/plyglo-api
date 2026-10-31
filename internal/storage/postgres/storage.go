package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	psql *pgxpool.Pool
}

func New(ctx context.Context, dbconn string) (*Storage, func(context.Context) error) {
	storage := &Storage{}
	pool, err := pgxpool.New(ctx, dbconn)
	if err != nil {
		return nil, nil
	}
	// defer dbConn.Close()

	if err = pool.Ping(ctx); err != nil {
		return nil, nil
	}

	storage.psql = pool
	return storage, func(ctx context.Context) error {
		return storage.Close(ctx)
	}
}

func (s *Storage) Close(ctx context.Context) error {
	s.psql.Close()
	return nil
}
