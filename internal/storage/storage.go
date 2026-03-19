package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	psql *pgxpool.Pool
}

func New(ctx context.Context) *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, psqlConn string) error {
	dbConn, err := pgxpool.New(ctx, psqlConn)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err = dbConn.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.psql = dbConn

	if err = s.up(ctx); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func (s *Storage) Disconnect(ctx context.Context) error {
	if s.psql != nil {
		s.psql.Close()
	}
	return nil
}

func (s *Storage) up(ctx context.Context) error {

	_, err := s.psql.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS public.courses (
			id uuid DEFAULT gen_random_uuid() NOT NULL,
			name character varying(255) NOT NULL,
			description text,
			updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
			logo_path character varying(512),
			is_published boolean DEFAULT false NOT NULL,
			PRIMARY KEY (id)
		);

		CREATE TABLE IF NOT EXISTS public.chapters (
			id uuid DEFAULT gen_random_uuid() NOT NULL,
			course_id uuid NOT NULL,
			name character varying(255) NOT NULL,
			description text,
			updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
			"position" integer DEFAULT 0 NOT NULL,
			PRIMARY KEY (id),
			CONSTRAINT chapters_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE
		);

		CREATE INDEX idx_chapters_course_id ON public.chapters USING btree (course_id);

		CREATE TABLE IF NOT EXISTS public.lessons (
			id uuid DEFAULT gen_random_uuid() NOT NULL,
			chapter_id uuid,
			title text NOT NULL,
			text text NOT NULL,
			updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			CONSTRAINT lessons_chapter_id_fkey FOREIGN KEY (chapter_id) REFERENCES public.chapters(id)
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	// log.Print("Migrations completed successfully")
	return nil
}

func (s *Storage) GetPool() *pgxpool.Pool {
	return s.psql
}
