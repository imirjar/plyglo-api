package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	models "github.com/imirjar/poliglotim-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *Storage) SelectLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	filter := bson.M{}
	if chapterID != "" {
		filter["chapter_id"] = chapterID
	}

	// bson.D для сортировки — сохраняет порядок полей
	opts := options.Find().SetSort(bson.D{{Key: "position", Value: 1}})

	cursor, err := s.db.Collection("lessons").Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to query lessons: %w", err)
	}
	defer cursor.Close(ctx)

	var lessons []models.Lesson
	if err = cursor.All(ctx, &lessons); err != nil {
		return nil, fmt.Errorf("failed to decode lessons: %w", err)
	}

	return lessons, nil
}

func (s *Storage) SelectLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	var lesson models.Lesson

	err := s.db.Collection("lessons").
		FindOne(ctx, bson.M{"_id": lessonID}).
		Decode(&lesson)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Lesson{}, fmt.Errorf("lesson with id %s not found", lessonID)
		}
		return models.Lesson{}, fmt.Errorf("failed to select lesson: %w", err)
	}

	return lesson, nil
}

func (s *Storage) InsertLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	// Mongo не сгенерирует string _id — генерируем сами
	if lesson.ID == "" {
		lesson.ID = bson.NewObjectID().Hex()
	}
	// Аналог DEFAULT now() в postgres
	if lesson.Updated.IsZero() {
		lesson.Updated = time.Now()
	}

	_, err := s.db.Collection("lessons").InsertOne(ctx, lesson)
	if err != nil {
		return models.Lesson{}, fmt.Errorf("failed to create lesson: %w", err)
	}

	return lesson, nil
}

func (s *Storage) UpdateLesson(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	update := bson.M{
		"$set": bson.M{
			"title":      lesson.Title,
			"text":       lesson.Text,
			"chapter_id": lesson.Chapter,
			"position":   lesson.Position,
			// Аналог CURRENT_TIMESTAMP
			"updated": time.Now(),
		},
	}

	// SetReturnDocument(options.After) — аналог RETURNING
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedLesson models.Lesson
	err := s.db.Collection("lessons").
		FindOneAndUpdate(ctx, bson.M{"_id": lesson.ID}, update, opts).
		Decode(&updatedLesson)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Lesson{}, fmt.Errorf("lesson with id %s not found", lesson.ID)
		}
		return models.Lesson{}, fmt.Errorf("failed to update lesson: %w", err)
	}

	return updatedLesson, nil
}

func (s *Storage) DeleteLesson(ctx context.Context, lessonID string) error {
	result, err := s.db.Collection("lessons").DeleteOne(ctx, bson.M{"_id": lessonID})
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
