package study_service

import (
	"context"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *StudyService) ReadChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	return s.Storage.SelectChapters(ctx, courseID)
}
func (s *StudyService) CreateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	return s.Storage.InsertChapter(ctx, chapter)
}

func (s *StudyService) ReadChapter(ctx context.Context, chapterID string) (models.Chapter, error) {
	return s.Storage.SelectChapter(ctx, chapterID)
}
func (s *StudyService) ReadNextChapter(ctx context.Context, chapterID string) (models.Chapter, error) {
	// chapter, err := s.Storage.SelectChapter(ctx, chapterID)
	nextchapter, err := s.Storage.SelectChapter(ctx, chapterID)
	return nextchapter, err
}
func (s *StudyService) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	return s.Storage.UpdateChapter(ctx, chapter)
}
func (s *StudyService) DeleteChapter(ctx context.Context, chapterID string) error {
	return s.Storage.DeleteChapter(ctx, chapterID)
}
