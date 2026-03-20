package service

import (
	"context"
	"fmt"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func New(ctx context.Context) *Service {
	return &Service{}
}

type Service struct {
	Storage Storage
}

type Storage interface {
	SelectCourses(context.Context) ([]models.Course, error)
	SelectCourse(context.Context, string) (models.Course, error)
	InsertCourse(context.Context, models.Course) (models.Course, error)
	UpdateCourse(context.Context, models.Course) (models.Course, error)
	DeleteCourse(context.Context, string) error

	SelectChapters(context.Context, string) ([]models.Chapter, error)
	SelectChapter(context.Context, string) (models.Chapter, error)
	InsertChapter(context.Context, models.Chapter) (models.Chapter, error)
	UpdateChapter(context.Context, models.Chapter) (models.Chapter, error)
	DeleteChapter(context.Context, string) error

	SelectLessons(context.Context, string) ([]models.Lesson, error)
	SelectLesson(context.Context, string) (models.Lesson, error)
	InsertLesson(context.Context, models.Lesson) (models.Lesson, error)
	UpdateLesson(context.Context, models.Lesson) (models.Lesson, error)
	DeleteLesson(context.Context, string) error

	Ping(context.Context) error
}

func (s *Service) Health(ctx context.Context) error {
	if s.Storage == nil {
		return fmt.Errorf("storage not initialized")
	}
	return s.Storage.Ping(ctx)
}
