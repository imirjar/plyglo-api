package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Updated     time.Time `json:"updated,omitempty"`
	LogoPath    *string   `json:"logo_path,omitempty"`
	IsPublished bool      `json:"is_published,omitempty"`
}

type Chapter struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Updated     time.Time `json:"updated"`
	Course      string    `json:"course_id,omitempty"`
	Position    int       `json:"position,omitempty"`
}

type Lesson struct {
	ID      string    `json:"id"`
	Chapter string    `json:"chapter_id,omitempty"`
	Title   string    `json:"title"`
	Text    string    `json:"text,omitempty"`
	Updated time.Time `json:"updated"`
}

type File struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Link    string             `json:"link"`
	Updated time.Time          `json:"updated"`
}
