package storage

import (
	"context"
	"fmt"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

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
func (s *Storage) InsertLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	query := `
        INSERT INTO lessons (title, text, chapter)
        VALUES ($1, $2, $3, $4)
    `

	// Выполняем запрос и получаем сгенерированные поля
	row := s.psql.QueryRow(
		ctx,
		query,
		lesson.Title,
		lesson.Text,
		lesson.Chapter,
	)

	var createdLesson models.Lesson
	err := row.Scan(
		&createdLesson.ID,
		&createdLesson.Title,
		&createdLesson.Text,
		&createdLesson.Chapter,
		&createdLesson.Updated,
	)
	if err != nil {
		return createdLesson, fmt.Errorf("failed to create course: %w", err)
	}

	return createdLesson, nil
}
func (s *Storage) UpdateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	query := `
        UPDATE lessons 
        SET title = $1, 
            text = $2, 
            chapter_id = $3,
            updated = CURRENT_TIMESTAMP
        WHERE id = $5
        RETURNING id, title, text, chapter_id, updated
    `

	var updatedLesson models.Lesson
	err := s.psql.QueryRow(
		ctx,
		query,
		lesson.Title,
		lesson.Text,
		lesson.Chapter,
	).Scan(
		&updatedLesson.ID,
		&updatedLesson.Title,
		&updatedLesson.Text,
		&updatedLesson.Chapter,
		&updatedLesson.Updated,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Lesson{}, fmt.Errorf("lesson with id %s not found", lesson.ID)
		}
		return models.Lesson{}, fmt.Errorf("failed to update lesson: %w", err)
	}

	return updatedLesson, nil
}
func (s *Storage) DeleteLesson(ctx context.Context, lessonID string) error {
	// Вариант 1: Простое удаление
	query := `DELETE FROM lessons WHERE id = $1`

	result, err := s.psql.Exec(ctx, query, lessonID)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	// Проверяем, был ли удален хотя бы один ряд
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return pgx.ErrNoRows // или sql.ErrNoRows
	}

	return nil
}
