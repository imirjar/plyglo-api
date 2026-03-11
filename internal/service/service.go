package service

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func New() *Service {
	return &Service{}
}

type Service struct {
	Storage Storage
}

type Storage interface {
	SelectCourses(context.Context) ([]models.Course, error)
	SelectCourseByID(context.Context, string) (models.Course, error)
	InsertCourse(context.Context, models.Course) (models.Course, error)
	DeleteCourse(context.Context, string) error

	SelectChapters(context.Context) ([]models.Chapter, error)
	SelectChapterByID(context.Context, string) (models.Chapter, error)
	InsertChapter(context.Context, models.Chapter) (models.Chapter, error)
	DeleteChapter(context.Context, string) error

	SelectLessons(context.Context, string) ([]models.Lesson, error)
	SelectLesson(context.Context, string) (models.Lesson, error)
}
