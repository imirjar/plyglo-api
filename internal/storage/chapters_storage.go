package storage

import (
	"context"
	"fmt"
	"log"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

// COURSES
func (s *Storage) SelectChapters(ctx context.Context) ([]models.Chapter, error) {

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
		ORDER BY 
			c.position
		ASC;
	`

	rows, err := s.psql.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}

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
			return nil, err
		}
		// log.Print(chapter)
		chapters = append(chapters, chapter)
	}

	return chapters, nil
}
func (s *Storage) SelectChapterByID(ctx context.Context, chapterID string) (models.Chapter, error) {

	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
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
		&chapter.Course,
		&chapter.Name,
		&chapter.Description,
		&chapter.Updated,
	)

	log.Print(chapter)
	return chapter, err
}
func (s *Storage) InsertChapter(ctx context.Context, course models.Chapter) (models.Chapter, error) {
	query := `
        INSERT INTO chapters (name, description, course_id)
        VALUES ($1, $2, $3)
        RETURNING id, name, description, updated, logo_path, is_published
    `

	// Выполняем запрос и получаем сгенерированные поля
	row := s.psql.QueryRow(
		ctx,
		query,
		course.Name,
		course.Description,
		course.Course,
	)

	var createdChapter models.Chapter
	err := row.Scan(
		&createdChapter.ID,
		&createdChapter.Course,
		&createdChapter.Name,
		&createdChapter.Description,
		&createdChapter.Updated,
	)
	if err != nil {
		return createdChapter, fmt.Errorf("failed to create chapter: %w", err)
	}

	return createdChapter, nil
}
func (s *Storage) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	query := `
        UPDATE courses 
        SET name = $1, 
            description = $2, 
            logo_path = $3, 
            is_published = $4,
            updated = CURRENT_TIMESTAMP
        WHERE id = $5
        RETURNING id, name, description, updated, logo_path, is_published
    `

	var updatedChapter models.Chapter
	err := s.psql.QueryRow(
		ctx,
		query,
		chapter.Name,
		chapter.Description,
		chapter.Course,
		chapter.ID,
	).Scan(
		&updatedChapter.ID,
		&updatedChapter.Name,
		&updatedChapter.Course,
		&updatedChapter.Description,
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
