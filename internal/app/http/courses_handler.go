package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
)

// getCourses возвращает список всех курсов
// @Summary Получить все курсы
// @Description Возвращает список всех доступных курсов
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course "Список курсов"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /courses [get]
func (srv *HttpServer) CoursesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			courses, err := srv.Service.ReadCourses(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(courses); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			var course models.Course

			if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			createdCourse, err := srv.Service.CreateCourse(r.Context(), course)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusCreated)
			if err = json.NewEncoder(w).Encode(createdCourse); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// CourseHandler возвращает список всех курсов
// @Summary Получить все курсы
// @Description Возвращает список всех доступных курсов
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course "Список курсов"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /courses [get]
func (srv *HttpServer) CourseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseID := vars["course_id"]

		switch r.Method {
		case "GET":
			courses, err := srv.Service.ReadCourseByID(r.Context(), courseID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(courses); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "PUT":
			var course models.Course

			if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			createdCourse, err := srv.Service.UpdateCourse(r.Context(), course)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusCreated)
			if err = json.NewEncoder(w).Encode(createdCourse); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "DELETE":

			if err := srv.Service.DeleteCourse(r.Context(), courseID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
