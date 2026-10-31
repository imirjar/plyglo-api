package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	psql *pgxpool.Pool
}

func New(ctx context.Context, dbconn string) *Storage {
	storage := &Storage{}
	pool, err := pgxpool.New(ctx, dbconn)
	if err != nil {
		return nil
	}
	// defer dbConn.Close()

	if err = pool.Ping(ctx); err != nil {
		return nil
	}

	storage.psql = pool
	return storage
}

func (s *Storage) Close(ctx context.Context) error {
	s.psql.Close()
	return nil
}
