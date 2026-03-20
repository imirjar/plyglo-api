package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
)

// LessonsHandler handles requests to the lessons collection
// @Summary Manage lessons collection
// @Description GET - returns list of lessons (filtered by chapter_id), POST - creates a new lesson
// @Tags lessons
// @Accept json
// @Produce json
// @Param chapter_id query string false "Chapter ID to filter lessons"
// @Param lesson body models.Lesson false "New lesson data (for POST)"
// @Success 200 {array} models.Lesson "List of lessons"
// @Success 201 {object} models.Lesson "Created lesson"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /lessons [get]
// @Router /lessons [post]
func (srv *HttpServer) LessonsHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// vars := mux.Vars(r)
			// lessonID := vars["lesson_id"]
			queryParams := r.URL.Query()
			chapterID := queryParams.Get("chapter_id")

			lesson, err := srv.Service.ReadLessons(r.Context(), chapterID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(lesson); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			var lesson models.Lesson

			if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			createdChapter, err := srv.Service.CreateLesson(r.Context(), lesson)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusCreated)
			if err = json.NewEncoder(w).Encode(createdChapter); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// LessonHandler handles requests to a specific lesson
// @Summary Manage lesson by ID
// @Description GET - returns lesson by ID, PUT - updates lesson, DELETE - removes lesson
// @Tags lessons
// @Accept json
// @Produce json
// @Param lesson_id path string true "Lesson ID"
// @Param lesson body models.Lesson false "Updated lesson data (for PUT)"
// @Success 200 {object} models.Lesson "Lesson information"
// @Success 201 {object} models.Lesson "Updated lesson"
// @Success 204 "Lesson successfully deleted"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 404 {object} ErrorResponse "Lesson not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /lessons/{lesson_id} [get]
// @Router /lessons/{lesson_id} [put]
// @Router /lessons/{lesson_id} [delete]
func (srv *HttpServer) LessonHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lessonID := vars["lesson_id"]
		switch r.Method {
		case "GET":

			lesson, err := srv.Service.ReadLesson(r.Context(), lessonID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(lesson); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "PUT":
			var lesson models.Lesson

			if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			lesson.ID = lessonID

			createdCourse, err := srv.Service.UpdateLesson(r.Context(), lesson)
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

			if err := srv.Service.DeleteLesson(r.Context(), lessonID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
