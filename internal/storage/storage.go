package storage

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	mongo "github.com/imirjar/poliglotim-api/internal/storage/mongo"
	postgres "github.com/imirjar/poliglotim-api/internal/storage/postgres"
)

func parseConnSchema(dbConn string) string {
	parsedURL, err := url.Parse(strings.TrimSpace(dbConn))
	if err != nil {
		return ""
	}

	// in upper case to match URL scheme format
	return strings.ToLower(parsedURL.Scheme)
}
func New(dbConn string) *Storage {
	storage := &Storage{}

	schema := parseConnSchema(dbConn)
	switch schema {
	case "mongodb":
		store, closeFn := mongo.New(context.Background(), dbConn)
		storage.Store = store
		storage.closeFunc = closeFn

	case "postgres":
		store, closeFn := postgres.New(context.Background(), dbConn)
		storage.Store = store
		storage.closeFunc = closeFn
	// case "s3":
	// 	return storage
	default:
		panic(fmt.Sprintf("unsupported schema: %s", schema))
	}
	return storage

}

type Storage struct {
	Store
	closeFunc func(context.Context) error
}

func (s *Storage) Close(ctx context.Context) error {
	return s.closeFunc(ctx)
}

type Store interface {

	// Courses
	SelectCourses(context.Context) ([]models.Course, error)
	SelectCourse(context.Context, string) (models.Course, error)
	InsertCourse(context.Context, models.Course) (models.Course, error)
	UpdateCourse(context.Context, models.Course) (models.Course, error)
	DeleteCourse(context.Context, string) error

	// Chapters
	SelectChapters(context.Context, string) ([]models.Chapter, error)
	SelectChapter(context.Context, string) (models.Chapter, error)
	InsertChapter(context.Context, models.Chapter) (models.Chapter, error)
	UpdateChapter(context.Context, models.Chapter) (models.Chapter, error)
	DeleteChapter(context.Context, string) error

	// Lessons
	SelectLessons(context.Context, string) ([]models.Lesson, error)
	SelectLesson(context.Context, string) (models.Lesson, error)
	InsertLesson(context.Context, models.Lesson) (models.Lesson, error)
	UpdateLesson(context.Context, models.Lesson) (models.Lesson, error)
	DeleteLesson(context.Context, string) error
}
