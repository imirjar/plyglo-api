package storage

import (
	"context"
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
func New(dbConn string) Storage {
	schema := parseConnSchema(dbConn)
	switch schema {
	case "mongo":
		return mongo.New(context.Background(), dbConn)
	case "postgrs":
		return postgres.New(context.Background(), dbConn)
	// case "s3":
	// 	return storage
	default:
		return nil
	}

}

type Storage interface {
	// Connect/Close
	Close(context.Context) error

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
