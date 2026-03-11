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
func (srv *HttpServer) ChaptersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			chapters, err := srv.Service.ReadChapters(r.Context())
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
func (srv *HttpServer) ChapterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		chapterID := vars["chapter_id"]

		switch r.Method {
		case "GET":
			chapters, err := srv.Service.ReadChapterByID(r.Context(), chapterID)
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
