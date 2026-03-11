package service

import (
	"context"

	models "github.com/imirjar/poliglotim-api/internal/domain"
)

func (s *Service) ReadCourses(ctx context.Context) ([]models.Course, error) {
	return s.Storage.SelectCourses(ctx)
}
func (s *Service) CreateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	return s.Storage.InsertCourse(ctx, course)
}

func (s *Service) ReadCourseByID(ctx context.Context, courseID string) (models.Course, error) {
	return s.Storage.SelectCourseByID(ctx, courseID)
}
func (s *Service) UpdateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	return s.Storage.InsertCourse(ctx, course)
}
func (s *Service) DeleteCourse(ctx context.Context, courseID string) error {
	return s.Storage.DeleteCourse(ctx, courseID)
}
