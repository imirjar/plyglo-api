package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @host localhost:6060
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type HttpServer struct {
	Port    string
	Service Service
}

func New(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (srv *HttpServer) Run() error {
	router := mux.NewRouter()

	// Swagger документация
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/courses", srv.CoursesHandler()).Methods("GET", "POST")
	router.Handle("/courses/{course_id}", srv.CourseHandler()).Methods("GET", "PUT", "DELETE")

	router.Handle("/chapters", srv.ChaptersHandler()).Methods("GET")
	router.Handle("/chapters/{chapter_id}", srv.ChapterHandler()).Methods("GET")

	router.Handle("/lessons", srv.getLessons()).Methods("GET")
	router.Handle("/lessons/{lesson_id}", srv.LessonHandler()).Methods("GET")

	// protected.Handle("/course/{course_id}", srv.getCourse()).Methods("GET")
	// router.Handle("/lesson/{lesson_id}", srv.getLesson()).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), router)
}

// HttpServer представляет HTTP сервер приложения
// @title Poliglotim API Gateway
// @version 1.0
// @description API для образовательной платформы Poliglotim
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@poliglotim.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
type Service interface {
	// Path '/courses'
	ReadCourses(context.Context) ([]models.Course, error)
	CreateCourse(context.Context, models.Course) (models.Course, error)
	UpdateCourse(context.Context, models.Course) (models.Course, error)

	// Path '/courses/{id}'
	ReadCourseByID(context.Context, string) (models.Course, error)
	DeleteCourse(context.Context, string) error

	// Path '/chapters'
	ReadChapters(context.Context) ([]models.Chapter, error)
	CreateChapter(context.Context, models.Chapter) (models.Chapter, error)
	UpdateChapter(context.Context, models.Chapter) (models.Chapter, error)

	// Path '/chapters/{id}'
	ReadChapterByID(context.Context, string) (models.Chapter, error)
	DeleteChapter(context.Context, string) error

	// ReadChapters(context.Context, string) ([]models.Chapter, error)
	ReadLessons(context.Context, string) ([]models.Lesson, error)
	ReadLesson(context.Context, string) (models.Lesson, error)
}
