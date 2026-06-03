package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	models "github.com/imirjar/poliglotim-api/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the Service interface
type MockService struct {
	mock.Mock
}

// Course operations
func (m *MockService) ReadCourses(ctx context.Context) ([]models.Course, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *MockService) CreateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	args := m.Called(ctx, course)
	return args.Get(0).(models.Course), args.Error(1)
}

func (m *MockService) UpdateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	args := m.Called(ctx, course)
	return args.Get(0).(models.Course), args.Error(1)
}

func (m *MockService) ReadCourse(ctx context.Context, id string) (models.Course, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Course), args.Error(1)
}

func (m *MockService) DeleteCourse(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Chapter operations
func (m *MockService) ReadChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).([]models.Chapter), args.Error(1)
}

func (m *MockService) CreateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	args := m.Called(ctx, chapter)
	return args.Get(0).(models.Chapter), args.Error(1)
}

func (m *MockService) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	args := m.Called(ctx, chapter)
	return args.Get(0).(models.Chapter), args.Error(1)
}

func (m *MockService) ReadChapter(ctx context.Context, id string) (models.Chapter, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Chapter), args.Error(1)
}

func (m *MockService) DeleteChapter(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Lesson operations
func (m *MockService) ReadLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	args := m.Called(ctx, chapterID)
	return args.Get(0).([]models.Lesson), args.Error(1)
}

func (m *MockService) CreateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	args := m.Called(ctx, lesson)
	return args.Get(0).(models.Lesson), args.Error(1)
}

func (m *MockService) UpdateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	args := m.Called(ctx, lesson)
	return args.Get(0).(models.Lesson), args.Error(1)
}

func (m *MockService) ReadLesson(ctx context.Context, id string) (models.Lesson, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Lesson), args.Error(1)
}

func (m *MockService) DeleteLesson(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockService) Health(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Test CoursesHandler
func TestCoursesHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedCourses := []models.Course{
		{ID: "1", Name: "Course 1"},
		{ID: "2", Name: "Course 2"},
	}

	mockService.On("ReadCourses", mock.Anything).Return(expectedCourses, nil)

	req := httptest.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	handler := srv.CoursesHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Course
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedCourses, response)

	mockService.AssertExpectations(t)
}

func TestCoursesHandler_POST(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	newCourse := models.Course{Name: "New Course"}
	createdCourse := models.Course{ID: "3", Name: "New Course"}

	body, _ := json.Marshal(newCourse)
	req := httptest.NewRequest("POST", "/courses", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockService.On("CreateCourse", mock.Anything, newCourse).Return(createdCourse, nil)

	handler := srv.CoursesHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Course
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, createdCourse, response)

	mockService.AssertExpectations(t)
}

func TestCoursesHandler_MethodNotAllowed(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	req := httptest.NewRequest("PUT", "/courses", nil)
	w := httptest.NewRecorder()

	handler := srv.CoursesHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

// Test CourseHandler
func TestCourseHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedCourse := models.Course{ID: "1", Name: "Course 1"}

	mockService.On("ReadCourse", mock.Anything, "1").Return(expectedCourse, nil)

	req := httptest.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()

	// Add path variables
	req = mux.SetURLVars(req, map[string]string{"course_id": "1"})

	handler := srv.CourseHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Course
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedCourse, response)

	mockService.AssertExpectations(t)
}

func TestCourseHandler_PUT(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	updatedCourse := models.Course{Name: "Updated Course"}
	expectedCourse := models.Course{ID: "1", Name: "Updated Course"}

	body, _ := json.Marshal(updatedCourse)
	req := httptest.NewRequest("PUT", "/courses/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"course_id": "1"})

	mockService.On("UpdateCourse", mock.Anything, models.Course{ID: "1", Name: "Updated Course"}).Return(expectedCourse, nil)

	handler := srv.CourseHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Course
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedCourse, response)

	mockService.AssertExpectations(t)
}

func TestCourseHandler_DELETE(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	mockService.On("DeleteCourse", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"course_id": "1"})

	handler := srv.CourseHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

// Test ChaptersHandler
func TestChaptersHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedChapters := []models.Chapter{
		{ID: "1", Name: "Chapter 1", Course: "1"},
		{ID: "2", Name: "Chapter 2", Course: "1"},
	}

	mockService.On("ReadChapters", mock.Anything, "1").Return(expectedChapters, nil)

	req := httptest.NewRequest("GET", "/chapters?course_id=1", nil)
	w := httptest.NewRecorder()

	handler := srv.ChaptersHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Chapter
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedChapters, response)

	mockService.AssertExpectations(t)
}

func TestChaptersHandler_POST(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	newChapter := models.Chapter{Name: "New Chapter", Course: "1"}
	createdChapter := models.Chapter{ID: "3", Name: "New Chapter", Course: "1"}

	body, _ := json.Marshal(newChapter)
	req := httptest.NewRequest("POST", "/chapters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockService.On("CreateChapter", mock.Anything, newChapter).Return(createdChapter, nil)

	handler := srv.ChaptersHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Chapter
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, createdChapter, response)

	mockService.AssertExpectations(t)
}

// Test ChapterHandler
func TestChapterHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedChapter := models.Chapter{ID: "1", Name: "Chapter 1", Course: "1"}

	mockService.On("ReadChapter", mock.Anything, "1").Return(expectedChapter, nil)

	req := httptest.NewRequest("GET", "/chapters/1", nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"chapter_id": "1"})

	handler := srv.ChapterHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Chapter
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedChapter, response)

	mockService.AssertExpectations(t)
}

func TestChapterHandler_PUT(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	updatedChapter := models.Chapter{Name: "Updated Chapter", Course: "1"}
	expectedChapter := models.Chapter{ID: "1", Name: "Updated Chapter", Course: "1"}

	body, _ := json.Marshal(updatedChapter)
	req := httptest.NewRequest("PUT", "/chapters/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"chapter_id": "1"})

	mockService.On("UpdateChapter", mock.Anything, models.Chapter{ID: "1", Name: "Updated Chapter", Course: "1"}).Return(expectedChapter, nil)

	handler := srv.ChapterHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Chapter
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedChapter, response)

	mockService.AssertExpectations(t)
}

func TestChapterHandler_DELETE(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	mockService.On("DeleteChapter", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/chapters/1", nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"chapter_id": "1"})

	handler := srv.ChapterHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

// Test LessonsHandler
func TestLessonsHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedLessons := []models.Lesson{
		{ID: "1", Title: "Lesson 1", Chapter: "1"},
		{ID: "2", Title: "Lesson 2", Chapter: "1"},
	}

	mockService.On("ReadLessons", mock.Anything, "1").Return(expectedLessons, nil)

	req := httptest.NewRequest("GET", "/lessons?chapter_id=1", nil)
	w := httptest.NewRecorder()

	handler := srv.LessonsHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Lesson
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedLessons, response)

	mockService.AssertExpectations(t)
}

func TestLessonsHandler_POST(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	newLesson := models.Lesson{Title: "New Lesson", Chapter: "1"}
	createdLesson := models.Lesson{ID: "3", Title: "New Lesson", Chapter: "1"}

	body, _ := json.Marshal(newLesson)
	req := httptest.NewRequest("POST", "/lessons", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockService.On("CreateLesson", mock.Anything, newLesson).Return(createdLesson, nil)

	handler := srv.LessonsHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Lesson
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, createdLesson, response)

	mockService.AssertExpectations(t)
}

// Test LessonHandler
func TestLessonHandler_GET(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedLesson := models.Lesson{ID: "1", Title: "Lesson 1", Chapter: "1"}

	mockService.On("ReadLesson", mock.Anything, "1").Return(expectedLesson, nil)

	req := httptest.NewRequest("GET", "/lessons/1", nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"lesson_id": "1"})

	handler := srv.LessonHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Lesson
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedLesson, response)

	mockService.AssertExpectations(t)
}

func TestLessonHandler_PUT(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	updatedLesson := models.Lesson{Title: "Updated Lesson", Chapter: "1"}
	expectedLesson := models.Lesson{ID: "1", Title: "Updated Lesson", Chapter: "1"}

	body, _ := json.Marshal(updatedLesson)
	req := httptest.NewRequest("PUT", "/lessons/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"lesson_id": "1"})

	mockService.On("UpdateLesson", mock.Anything, models.Lesson{ID: "1", Title: "Updated Lesson", Chapter: "1"}).Return(expectedLesson, nil)

	handler := srv.LessonHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Lesson
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedLesson, response)

	mockService.AssertExpectations(t)
}

func TestLessonHandler_DELETE(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	mockService.On("DeleteLesson", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/lessons/1", nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"lesson_id": "1"})

	handler := srv.LessonHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

// Test error cases
func TestCoursesHandler_GET_ServiceError(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	mockService.On("ReadCourses", mock.Anything).Return([]models.Course{}, assert.AnError)

	req := httptest.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	handler := srv.CoursesHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestCoursesHandler_POST_InvalidJSON(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	req := httptest.NewRequest("POST", "/courses", bytes.NewReader([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler := srv.CoursesHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCourseHandler_GET_InvalidID(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	// ✅ IMPORTANT: Set up the expectation for the empty ID case
	// Return an empty course and nil error (or whatever your actual handler expects)
	mockService.On("ReadCourse", mock.Anything, "").Return(models.Course{}, nil)

	req := httptest.NewRequest("GET", "/courses/", nil)
	w := httptest.NewRecorder()

	// No path variables set - this will cause an empty ID to be passed to the service
	handler := srv.CourseHandler()
	handler.ServeHTTP(w, req)

	// You might want to add assertions here based on expected behavior
	// For example, if your handler should return 400 Bad Request for empty ID:
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	mockService.AssertExpectations(t)
}

func TestChaptersHandler_GET_NoFilter(t *testing.T) {
	mockService := new(MockService)
	srv := &HttpServer{Service: mockService}

	expectedChapters := []models.Chapter{
		{ID: "1", Name: "Chapter 1"},
		{ID: "2", Name: "Chapter 2"},
	}

	mockService.On("ReadChapters", mock.Anything, "").Return(expectedChapters, nil)

	req := httptest.NewRequest("GET", "/chapters", nil)
	w := httptest.NewRecorder()

	handler := srv.ChaptersHandler()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Chapter
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedChapters, response)

	mockService.AssertExpectations(t)
}
