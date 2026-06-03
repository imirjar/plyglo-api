package http

import (
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
func (srv *HttpServer) LessonsHandler() http.HandlerFunc {
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
func (srv *HttpServer) LessonHandler() http.HandlerFunc {
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
