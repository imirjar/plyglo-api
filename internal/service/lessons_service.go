package service

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func (s *Service) ReadLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	return s.Storage.SelectLessons(ctx, chapterID)
}

func (s *Service) ReadLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	return s.Storage.SelectLesson(ctx, lessonID)
}
