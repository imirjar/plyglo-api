package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	connString string
	psql       *pgxpool.Pool
}

func New(ctx context.Context, opts ...func(*Storage)) *Storage {
	storage := &Storage{}

	for _, opt := range opts {
		opt(storage)
	}

	return storage
}

func WithDB(psqlConn string) func(*Storage) {
	return func(s *Storage) {
		s.connString = psqlConn
	}
}

func (s *Storage) Conn(ctx context.Context) error {
	dbConn, err := pgxpool.New(ctx, s.connString)
	if err != nil {
		return err
	}
	// defer dbConn.Close()

	if err = dbConn.Ping(ctx); err != nil {
		return err
	}

	s.psql = dbConn
	return nil
}

func (s *Storage) Close(ctx context.Context) {
	if s.psql != nil {
		s.psql.Close()
		log.Println("Database connection closed")
	}
}
