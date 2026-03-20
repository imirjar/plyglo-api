package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Language course API
// @version 1.0
// @description API for the educational platform

// @contact.name Artem Zadorov
// @contact.email azadorov1234@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
type HttpServer struct {
	Port    string
	Service Service
}

// New creates a new HTTP server instance
func New(ctx context.Context, port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

// Run starts the HTTP server and configures all routes
func (srv *HttpServer) Run(ctx context.Context) error {
	router := mux.NewRouter()

	// Course routes
	router.Handle("/courses", srv.CoursesHandler(ctx)).Methods("GET", "POST")
	router.Handle("/courses/{course_id}", srv.CourseHandler(ctx)).Methods("GET", "PUT", "DELETE")

	// Chapter routes
	router.Handle("/chapters", srv.ChaptersHandler(ctx)).Methods("GET", "POST")
	router.Handle("/chapters/{chapter_id}", srv.ChapterHandler(ctx)).Methods("GET", "PUT", "DELETE")

	// Lesson routes
	router.Handle("/lessons", srv.LessonsHandler(ctx)).Methods("GET", "POST")
	router.Handle("/lessons/{lesson_id}", srv.LessonHandler(ctx)).Methods("GET", "PUT", "DELETE")

	// Swagger documentation endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/health", srv.HealthHandler())

	return http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), router)
}

// Service defines the interface for business logic operations
// All methods handle CRUD operations for courses, chapters, and lessons
type Service interface {
	// Course operations
	// @Description Get all courses
	ReadCourses(context.Context) ([]models.Course, error)

	// @Description Create a new course
	// @Success 201
	CreateCourse(context.Context, models.Course) (models.Course, error)

	// @Description Update an existing course
	UpdateCourse(context.Context, models.Course) (models.Course, error)

	// @Description Get a specific course by ID
	ReadCourse(context.Context, string) (models.Course, error)

	// @Description Delete a course by ID
	// @Success 204
	DeleteCourse(context.Context, string) error

	// Chapter operations
	// @Description Get all chapters (optionally filtered by course_id)
	ReadChapters(context.Context, string) ([]models.Chapter, error)

	// @Description Create a new chapter
	CreateChapter(context.Context, models.Chapter) (models.Chapter, error)

	// @Description Update an existing chapter
	UpdateChapter(context.Context, models.Chapter) (models.Chapter, error)

	// @Description Get a specific chapter by ID
	ReadChapter(context.Context, string) (models.Chapter, error)

	// @Description Delete a chapter by ID
	// @Success 204
	DeleteChapter(context.Context, string) error

	// Lesson operations
	// @Description Get all lessons (optionally filtered by chapter_id)
	ReadLessons(context.Context, string) ([]models.Lesson, error)

	// @Description Create a new lesson
	CreateLesson(context.Context, models.Lesson) (models.Lesson, error)

	// @Description Update an existing lesson
	UpdateLesson(context.Context, models.Lesson) (models.Lesson, error)

	// @Description Get a specific lesson by ID
	ReadLesson(context.Context, string) (models.Lesson, error)

	// @Description Delete a lesson by ID
	// @Success 204
	DeleteLesson(context.Context, string) error

	Health(ctx context.Context) error
}

// ErrorResponse represents an error message returned to the client
type ErrorResponse struct {
	Error string `json:"error" example:"error description"`
	Code  int    `json:"code,omitempty" example:"400"`
}

func (srv *HttpServer) HealthHandler() http.HandlerFunc {
	h := health.NewChecker(
		health.WithTimeout(5*time.Second),
		health.WithCheck(health.Check{
			Name:    "postgres",
			Timeout: 2 * time.Second,
			Check:   srv.Service.Health,
		}),
	)

	return health.NewHandler(h)
}
