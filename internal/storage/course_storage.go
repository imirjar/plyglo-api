package storage

import (
	"context"
	"fmt"
	"log"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

// COURSES
func (s *Storage) SelectCourses(ctx context.Context) ([]models.Course, error) {
	var courses []models.Course

	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
			c.updated, 
			c.logo_path,
			c.is_published
		FROM 
			courses c
		ORDER BY 
			c.name;
	`

	rows, err := s.psql.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var course models.Course

		err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.Updated,
			&course.LogoPath,
			&course.IsPublished,
		)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
func (s *Storage) SelectCourseByID(ctx context.Context, courseID string) (models.Course, error) {

	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
			c.updated, 
			c.logo_path,
			c.is_published
		FROM 
			courses c
		WHERE 
			c.id = $1
	`

	row := s.psql.QueryRow(ctx, query, courseID)

	var course models.Course
	err := row.Scan(
		&course.ID,
		&course.Name,
		&course.Description,
		&course.Updated,
		&course.LogoPath,
		&course.IsPublished,
	)

	log.Print(course)
	return course, err
}
func (s *Storage) InsertCourse(ctx context.Context, course models.Course) (models.Course, error) {
	query := `
        INSERT INTO courses (name, description, logo_path)
        VALUES ($1, $2, $3)
        RETURNING id, name, description, updated, logo_path, is_published
    `

	// Выполняем запрос и получаем сгенерированные поля
	row := s.psql.QueryRow(
		ctx,
		query,
		course.Name,
		course.Description,
		course.LogoPath,
	)

	var createdCourse models.Course
	err := row.Scan(
		&createdCourse.ID,
		&createdCourse.Name,
		&createdCourse.Description,
		&createdCourse.Updated,
		&createdCourse.LogoPath,
		&createdCourse.IsPublished,
	)
	if err != nil {
		return createdCourse, fmt.Errorf("failed to create course: %w", err)
	}

	return createdCourse, nil
}
func (s *Storage) UpdateCourse(ctx context.Context, course models.Course) (models.Course, error) {
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

	var updatedCourse models.Course
	err := s.psql.QueryRow(
		ctx,
		query,
		course.Name,
		course.Description,
		course.LogoPath,
		course.IsPublished,
		course.ID,
	).Scan(
		&updatedCourse.ID,
		&updatedCourse.Name,
		&updatedCourse.Description,
		&updatedCourse.Updated,
		&updatedCourse.LogoPath,
		&updatedCourse.IsPublished,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Course{}, fmt.Errorf("course with id %s not found", course.ID)
		}
		return models.Course{}, fmt.Errorf("failed to update course: %w", err)
	}

	return updatedCourse, nil
}
func (s *Storage) DeleteCourse(ctx context.Context, courseID string) error {
	// Вариант 1: Простое удаление
	query := `DELETE FROM courses WHERE id = $1`

	result, err := s.psql.Exec(ctx, query, courseID)
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
