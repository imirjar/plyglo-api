package study_service

import (
	"context"
	"log"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *StudyService) ReadLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	return s.Storage.SelectLessons(ctx, chapterID)
}

func (s *StudyService) ReadLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	return s.Storage.SelectLesson(ctx, lessonID)
}

func (s *StudyService) ReadNextLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	return s.Storage.SelectLesson(ctx, lessonID)
}

func (s *StudyService) CreateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	log.Print(lesson)
	return s.Storage.InsertLesson(ctx, lesson)
}

func (s *StudyService) UpdateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	log.Print(lesson)
	return s.Storage.UpdateLesson(ctx, lesson)
}
func (s *StudyService) DeleteLesson(ctx context.Context, lessonID string) error {
	return s.Storage.DeleteLesson(ctx, lessonID)
}
