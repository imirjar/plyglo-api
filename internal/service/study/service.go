package service

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func New(ctx context.Context, opts ...func(*Service)) *Service {
	service := &Service{}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

type Service struct {
	Storage Storage
}

func WithStorage(storage Storage) func(*Service) {
	return func(s *Service) {
		s.Storage = storage
	}
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

	// Ping(context.Context) error
}

// func (s *Service) Health(ctx context.Context) error {
// 	if s.Storage == nil {
// 		return fmt.Errorf("storage not initialized")
// 	}
// 	return s.Storage.Ping(ctx)
// }
