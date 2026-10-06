package study_service

import (
	"context"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *StudyService) ReadCourses(ctx context.Context) ([]models.Course, error) {
	return s.Storage.SelectCourses(ctx)
}
func (s *StudyService) CreateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	return s.Storage.InsertCourse(ctx, course)
}

func (s *StudyService) ReadCourse(ctx context.Context, courseID string) (models.Course, error) {
	return s.Storage.SelectCourse(ctx, courseID)
}
func (s *StudyService) UpdateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	return s.Storage.UpdateCourse(ctx, course)
}
func (s *StudyService) DeleteCourse(ctx context.Context, courseID string) error {
	return s.Storage.DeleteCourse(ctx, courseID)
}
