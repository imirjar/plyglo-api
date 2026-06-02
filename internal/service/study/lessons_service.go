package service

import (
	"context"
	"log"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func (s *Service) ReadLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	return s.Storage.SelectLessons(ctx, chapterID)
}

func (s *Service) ReadLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	return s.Storage.SelectLesson(ctx, lessonID)
}

func (s *Service) CreateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	log.Print(lesson)
	return s.Storage.InsertLesson(ctx, lesson)
}

func (s *Service) UpdateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	log.Print(lesson)
	return s.Storage.UpdateLesson(ctx, lesson)
}
func (s *Service) DeleteLesson(ctx context.Context, lessonID string) error {
	return s.Storage.DeleteLesson(ctx, lessonID)
}
