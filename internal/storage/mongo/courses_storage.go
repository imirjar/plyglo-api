package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *Storage) SelectCourses(ctx context.Context) ([]models.Course, error) {
	// bson.D для сортировки — сохраняет порядок полей
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: 1}})

	cursor, err := s.client.Database("courses").Collection("courses").Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to query courses: %w", err)
	}
	defer cursor.Close(ctx)

	var courses []models.Course
	if err = cursor.All(ctx, &courses); err != nil {
		return nil, fmt.Errorf("failed to decode courses: %w", err)
	}

	return courses, nil
}

func (s *Storage) SelectCourse(ctx context.Context, courseID string) (models.Course, error) {
	var course models.Course

	err := s.client.Database("courses").Collection("courses").
		FindOne(ctx, bson.M{"_id": courseID}).
		Decode(&course)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Course{}, fmt.Errorf("course with id %s not found", courseID)
		}
		return models.Course{}, fmt.Errorf("failed to select course: %w", err)
	}

	log.Print(course)
	return course, nil
}

func (s *Storage) InsertCourse(ctx context.Context, course models.Course) (models.Course, error) {
	// Mongo не сгенерирует string _id — генерируем сами
	if course.ID == "" {
		course.ID = bson.NewObjectID().Hex()
	}
	// Аналог DEFAULT now() в postgres
	if course.Updated.IsZero() {
		course.Updated = time.Now()
	}

	_, err := s.client.Database("courses").Collection("courses").InsertOne(ctx, course)
	if err != nil {
		return models.Course{}, fmt.Errorf("failed to create course: %w", err)
	}

	return course, nil
}

func (s *Storage) UpdateCourse(ctx context.Context, course models.Course) (models.Course, error) {
	update := bson.M{
		"$set": bson.M{
			"name":         course.Name,
			"description":  course.Description,
			"logo_path":    course.LogoPath,
			"is_published": course.IsPublished,
			// Аналог CURRENT_TIMESTAMP
			"updated": time.Now(),
		},
	}

	// SetReturnDocument(options.After) — аналог RETURNING
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedCourse models.Course
	err := s.client.Database("courses").Collection("courses").
		FindOneAndUpdate(ctx, bson.M{"_id": course.ID}, update, opts).
		Decode(&updatedCourse)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Course{}, fmt.Errorf("course with id %s not found", course.ID)
		}
		return models.Course{}, fmt.Errorf("failed to update course: %w", err)
	}

	return updatedCourse, nil
}

func (s *Storage) DeleteCourse(ctx context.Context, courseID string) error {
	result, err := s.client.Database("courses").Collection("courses").DeleteOne(ctx, bson.M{"_id": courseID})
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
