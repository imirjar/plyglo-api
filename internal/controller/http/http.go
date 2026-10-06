package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gorilla/mux"
	"github.com/imirjar/poliglotim-api/internal/controller/http/client"
	"github.com/imirjar/poliglotim-api/internal/models"
	"github.com/rs/cors"
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
type HTTP struct {
	server  *http.Server
	client  Client
	Service Service
}

// New creates a new HTTP server instance
func New(opts ...func(*HTTP)) *HTTP {
	server := &HTTP{}

	for _, opt := range opts {
		opt(server)
	}
	return server
}

// Run starts the HTTP server and configures all routes
func (srv *HTTP) Run() error {
	log.Printf("HTTP server is running on %s", srv.server.Addr)
	return srv.server.ListenAndServe()
}

func (srv *HTTP) Stop(ctx context.Context) error {
	log.Println("Stopping HTTP server...")
	return srv.server.Shutdown(ctx)
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

	// @Description Get a specific chapter by ID
	ReadNextChapter(context.Context, string) (models.Chapter, error)

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

	// @Description Get a specific lesson by ID
	ReadNextLesson(context.Context, string) (models.Lesson, error)

	// @Description Delete a lesson by ID
	// @Success 204
	DeleteLesson(context.Context, string) error

	// Health(ctx context.Context) error
}

type Client interface {
	ValidateToken(string) (bool, error)
}

// ErrorResponse represents an error message returned to the client
type ErrorResponse struct {
	Error string `json:"error" example:"error description"`
	Code  int    `json:"code,omitempty" example:"400"`
}

func (srv *HTTP) HealthHandler() http.HandlerFunc {
	h := health.NewChecker(
		health.WithTimeout(5*time.Second),
		health.WithCheck(health.Check{
			Name:    "postgres",
			Timeout: 2 * time.Second,
			// Check:   srv.Service.Health,
		}),
	)

	return health.NewHandler(h)
}

func WithService(service Service) func(*HTTP) {
	return func(s *HTTP) {
		s.Service = service
	}
}

func WithClient() func(*HTTP) {
	return func(srv *HTTP) {
		srv.client = client.New()
	}
}

func WithServer(port string) func(*HTTP) {

	return func(srv *HTTP) {
		router := mux.NewRouter()

		// Course routes
		courses := router.PathPrefix("/courses").Subrouter()
		courses.Handle("", srv.CoursesHandler()).Methods("GET", "POST")
		courses.Handle("/{course_id}", srv.CourseHandler()).Methods("GET", "PUT", "DELETE")

		// Chapter routes
		chapters := router.PathPrefix("/chapters").Subrouter()
		chapters.Use(srv.authMiddleware)
		chapters.Handle("", srv.ChaptersHandler()).Methods("GET", "POST")
		chapters.Handle("/{chapter_id}", srv.ChapterHandler()).Methods("GET", "PUT", "DELETE")
		chapters.Handle("/{chapter_id}/next", srv.NextChapterHandler()).Methods("GET")

		// Lesson routes
		lessons := router.PathPrefix("/lessons").Subrouter()
		lessons.Use(srv.authMiddleware)
		lessons.Handle("", srv.LessonsHandler()).Methods("GET", "POST")
		lessons.Handle("/{lesson_id}", srv.LessonHandler()).Methods("GET", "PUT", "DELETE")
		// lessons.Handle("/{lesson_id}/next", srv.NextLessonHandler()).Methods("GET")

		// Swagger documentation endpoint
		router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		router.Handle("/health", srv.HealthHandler())

		// Настройка CORS
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"https://study.plyglo.com"}, // Укажите ваш фронтенд домен
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           300, // Кэширование preflight запросов на 5 минут
		})

		srv.server = &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: c.Handler(router),
		}
	}
}
