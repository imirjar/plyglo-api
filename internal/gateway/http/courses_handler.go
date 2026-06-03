package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
)

// CoursesHandler handles requests to the courses collection
// @Summary Manage courses collection
// @Description GET - returns list of all courses, POST - creates a new course
// @Tags courses
// @Accept json
// @Produce json
// @Param course body models.Course false "New course data (for POST)"
// @Success 200 {array} models.Course "List of courses"
// @Success 201 {object} models.Course "Created course"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /courses [get]
// @Router /courses [post]
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

// CourseHandler handles requests to a specific course
// @Summary Manage course by ID
// @Description GET - returns course by ID, PUT - updates course, DELETE - removes course
// @Tags courses
// @Accept json
// @Produce json
// @Param course_id path string true "Course ID"
// @Param course body models.Course false "Updated course data (for PUT)"
// @Success 200 {object} models.Course "Course information"
// @Success 201 {object} models.Course "Updated course"
// @Success 204 "Course successfully deleted"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /courses/{course_id} [get]
// @Router /courses/{course_id} [put]
// @Router /courses/{course_id} [delete]
func (srv *HttpServer) CourseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseID := vars["course_id"]

		switch r.Method {
		case "GET":
			courses, err := srv.Service.ReadCourse(r.Context(), courseID)
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
			course.ID = courseID

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
