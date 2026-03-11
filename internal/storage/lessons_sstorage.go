package storage

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

// LESSONS
func (s *Storage) SelectLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {

	var lessons []models.Lesson

	var rows pgx.Rows
	var err error

	if chapterID == "" {
		query := `SELECT 
			id, title, updated
		FROM lessons`

		rows, err = s.psql.Query(ctx, query)
	} else {
		query := `
		SELECT 
			id, 
			title, 
			updated
		FROM 
			lessons
		WHERE chapter_id = $1

	`
		rows, err = s.psql.Query(ctx, query, chapterID)

	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Итерируем по результатам

	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.Title,
			&lesson.Updated,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}
func (s *Storage) SelectLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	query := `
		SELECT 
			l.id, 
			l.title, 
			l.text, 
			l.updated
		FROM 
			lessons l
		WHERE
			l.id = $1
	`
	row := s.psql.QueryRow(ctx, query, lessonID)

	var lesson models.Lesson
	err := row.Scan(
		&lesson.ID,
		&lesson.Title,
		&lesson.Text,
		&lesson.Updated,
	)
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}
