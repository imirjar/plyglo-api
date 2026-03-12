package storage

import (
	"context"
	"fmt"
	"log"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) SelectChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	var rows pgx.Rows
	var err error

	// Базовый запрос
	query := `
        SELECT 
            c.id, 
            c.name, 
            c.description, 
            c.course_id,
            c.position,
            c.updated
        FROM 
            chapters c
    `

	// Добавляем WHERE если нужно
	if courseID != "" {
		query += " WHERE course_id = $1"
	}

	// Добавляем сортировку
	query += " ORDER BY c.position ASC"

	// Выполняем запрос
	if courseID != "" {
		rows, err = s.psql.Query(ctx, query, courseID)
	} else {
		rows, err = s.psql.Query(ctx, query)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query chapters: %w", err)
	}
	defer rows.Close()

	var chapters []models.Chapter
	for rows.Next() {
		var chapter models.Chapter
		err := rows.Scan(
			&chapter.ID,
			&chapter.Name,
			&chapter.Description,
			&chapter.Course,
			&chapter.Position,
			&chapter.Updated,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chapter: %w", err)
		}
		chapters = append(chapters, chapter)
	}

	// Проверяем ошибки после итерации
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating chapters: %w", err)
	}

	return chapters, nil
}
func (s *Storage) SelectChapter(ctx context.Context, chapterID string) (models.Chapter, error) {

	query := `
		SELECT 
			c.id,
			c.name,
			c.description,
			c.course_id,
			c.position, 
			c.updated
		FROM 
			chapters c
		WHERE 
			c.id = $1
	`

	row := s.psql.QueryRow(ctx, query, chapterID)

	var chapter models.Chapter
	err := row.Scan(
		&chapter.ID,
		&chapter.Name,
		&chapter.Description,
		&chapter.Course,
		&chapter.Position,
		&chapter.Updated,
	)

	log.Print(chapter)
	return chapter, err
}
func (s *Storage) InsertChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	query := `
        INSERT INTO chapters (name, description, course_id, position)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, description, course_id, position, updated
    `

	// Выполняем запрос и получаем сгенерированные поля
	row := s.psql.QueryRow(
		ctx,
		query,
		chapter.Name,
		chapter.Description,
		chapter.Course,
		chapter.Position,
	)

	var createdChapter models.Chapter
	err := row.Scan(
		&createdChapter.ID,
		&createdChapter.Name,
		&createdChapter.Description,
		&createdChapter.Course,
		&createdChapter.Position,
		&createdChapter.Updated,
	)
	if err != nil {
		return createdChapter, fmt.Errorf("failed to create chapter: %w", err)
	}

	return createdChapter, nil
}
func (s *Storage) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	// log.Print(chapter)
	log.Print("OOOOK")
	query := `
        UPDATE chapters 
        SET name = $1, 
            description = $2, 
            position = $3,
			course_id = $4,
			updated = $5
        WHERE id = $6
        RETURNING id, name, description, position, course_id, updated
    `

	var updatedChapter models.Chapter
	err := s.psql.QueryRow(
		ctx,
		query,
		chapter.Name,
		chapter.Description,
		chapter.Position,
		chapter.Course,
		chapter.Updated,
		chapter.ID,
	).Scan(
		&updatedChapter.ID,
		&updatedChapter.Name,
		&updatedChapter.Description,
		&updatedChapter.Position,
		&updatedChapter.Course,
		&updatedChapter.Updated,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Chapter{}, fmt.Errorf("chapter with id %s not found", chapter.ID)
		}
		return models.Chapter{}, fmt.Errorf("failed to update chapter: %w", err)
	}

	return updatedChapter, nil
}
func (s *Storage) DeleteChapter(ctx context.Context, chapterID string) error {
	// Вариант 1: Простое удаление
	query := `DELETE FROM chapters WHERE id = $1`

	result, err := s.psql.Exec(ctx, query, chapterID)
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
