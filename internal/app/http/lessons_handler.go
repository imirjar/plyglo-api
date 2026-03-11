package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// getLesson возвращает информацию об уроке
// @Summary Получить урок по ID
// @Description Возвращает информацию об уроке по временной ссылке
// @Tags lessons
// @Accept json
// @Produce json
// @Param lesson_id path string true "ID урока"
// @Success 200 {object} models.Lesson "Информация об уроке"
// @Failure 400 {object} ErrorResponse "Неверный ID урока"
// @Failure 403 {object} ErrorResponse "Доступ запрещен"
// @Failure 404 {object} ErrorResponse "Урок не найден"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /lesson/{lesson_id} [get]
func (srv *HttpServer) getLessons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func (srv *HttpServer) LessonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lessonID := vars["lesson_id"]

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
	}
}
