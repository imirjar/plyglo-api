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

func (s *Storage) SelectChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	filter := bson.M{}
	if courseID != "" {
		filter["course_id"] = courseID
	}

	// bson.D для сортировки — сохраняет порядок полей
	opts := options.Find().SetSort(bson.D{{Key: "position", Value: 1}})

	cursor, err := s.db.Collection("chapters").Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to query chapters: %w", err)
	}
	defer cursor.Close(ctx)

	var chapters []models.Chapter
	if err = cursor.All(ctx, &chapters); err != nil {
		return nil, fmt.Errorf("failed to decode chapters: %w", err)
	}

	return chapters, nil
}

func (s *Storage) SelectChapter(ctx context.Context, chapterID string) (models.Chapter, error) {
	var chapter models.Chapter

	err := s.db.Collection("chapters").
		FindOne(ctx, bson.M{"_id": chapterID}).
		Decode(&chapter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Chapter{}, fmt.Errorf("chapter with id %s not found", chapterID)
		}
		return models.Chapter{}, fmt.Errorf("failed to select chapter: %w", err)
	}

	log.Print(chapter)
	return chapter, nil
}

func (s *Storage) InsertChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	// Mongo не сгенерирует string _id — генерируем сами
	if chapter.ID == "" {
		chapter.ID = bson.NewObjectID().Hex()
	}
	// Аналог DEFAULT now() в postgres
	if chapter.Updated.IsZero() {
		chapter.Updated = time.Now()
	}

	_, err := s.db.Collection("chapters").InsertOne(ctx, chapter)
	if err != nil {
		return models.Chapter{}, fmt.Errorf("failed to create chapter: %w", err)
	}

	return chapter, nil
}

func (s *Storage) UpdateChapter(ctx context.Context, chapter models.Chapter) (models.Chapter, error) {
	log.Print("OOOOK")

	update := bson.M{
		"$set": bson.M{
			"name":        chapter.Name,
			"description": chapter.Description,
			"position":    chapter.Position,
			"course_id":   chapter.Course,
			"updated":     chapter.Updated,
		},
	}

	// SetReturnDocument(options.After) — аналог RETURNING
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedChapter models.Chapter
	err := s.db.Collection("chapters").
		FindOneAndUpdate(ctx, bson.M{"_id": chapter.ID}, update, opts).
		Decode(&updatedChapter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Chapter{}, fmt.Errorf("chapter with id %s not found", chapter.ID)
		}
		return models.Chapter{}, fmt.Errorf("failed to update chapter: %w", err)
	}

	return updatedChapter, nil
}

func (s *Storage) DeleteChapter(ctx context.Context, chapterID string) error {
	result, err := s.db.Collection("chapters").DeleteOne(ctx, bson.M{"_id": chapterID})
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
