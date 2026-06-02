package service

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func (s *Service) ReadChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	return s.Storage.SelectChapters(ctx, courseID)
}
func (s *Service) CreateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	return s.Storage.InsertChapter(ctx, chapter)
}

func (s *Service) ReadChapter(ctx context.Context, chapterID string) (models.Chapter, error) {
	return s.Storage.SelectChapter(ctx, chapterID)
}
func (s *Service) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	return s.Storage.UpdateChapter(ctx, chapter)
}
func (s *Service) DeleteChapter(ctx context.Context, chapterID string) error {
	return s.Storage.DeleteChapter(ctx, chapterID)
}
