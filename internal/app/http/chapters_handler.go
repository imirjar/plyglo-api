package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
)

// ChaptersHandler handles requests to the chapters collection
// @Summary Manage chapters collection
// @Description GET - returns list of chapters (filtered by course_id), POST - creates a new chapter
// @Tags chapters
// @Accept json
// @Produce json
// @Param course_id query string false "Course ID to filter chapters"
// @Param chapter body models.Chapter false "New chapter data (for POST)"
// @Success 200 {array} models.Chapter "List of chapters"
// @Success 201 {object} models.Chapter "Created chapter"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /chapters [get]
// @Router /chapters [post]
func (srv *HttpServer) ChaptersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			queryParams := r.URL.Query()
			chapterID := queryParams.Get("course_id")

			chapters, err := srv.Service.ReadChapters(r.Context(), chapterID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(chapters); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			var chapter models.Chapter

			if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			createdChapter, err := srv.Service.CreateChapter(r.Context(), chapter)
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

// ChapterHandler handles requests to a specific chapter
// @Summary Manage chapter by ID
// @Description GET - returns chapter by ID, PUT - updates chapter, DELETE - removes chapter
// @Tags chapters
// @Accept json
// @Produce json
// @Param chapter_id path string true "Chapter ID"
// @Param chapter body models.Chapter false "Updated chapter data (for PUT)"
// @Success 200 {object} models.Chapter "Chapter information"
// @Success 201 {object} models.Chapter "Updated chapter"
// @Success 204 "Chapter successfully deleted"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /chapters/{chapter_id} [get]
// @Router /chapters/{chapter_id} [put]
// @Router /chapters/{chapter_id} [delete]
func (srv *HttpServer) ChapterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		chapterID := vars["chapter_id"]

		switch r.Method {
		case "GET":
			chapters, err := srv.Service.ReadChapter(r.Context(), chapterID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(chapters); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "PUT":
			var chapter models.Chapter

			if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			chapter.ID = chapterID

			createdChapter, err := srv.Service.UpdateChapter(r.Context(), chapter)
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
		case "DELETE":

			if err := srv.Service.DeleteChapter(r.Context(), chapterID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
