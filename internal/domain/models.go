package models

import (
	"time"
)

type Course struct {
	ID          string    `bson:"_id"          json:"id"`
	Name        string    `bson:"name"         json:"name"`
	Description string    `bson:"description"  json:"description"`
	Updated     time.Time `bson:"updated"      json:"updated"`
	LogoPath    *string   `bson:"logo_path"    json:"logo_path,omitempty"`
	IsPublished bool      `bson:"is_published" json:"is_published,omitempty"`
}

type Chapter struct {
	ID          string    `bson:"_id"          json:"id"`
	Name        string    `bson:"name"         json:"name"`
	Description string    `bson:"description"  json:"description"`
	Updated     time.Time `bson:"updated"      json:"updated"`
	Course      string    `bson:"course_id"    json:"course_id,omitempty"`
	Position    int       `bson:"position"     json:"position,omitempty"`
}

type Lesson struct {
	ID       string    `bson:"_id"          json:"id"`
	Chapter  string    `bson:"chapter_id"   json:"chapter_id,omitempty"`
	Title    string    `bson:"title"        json:"title"`
	Text     string    `bson:"text"         json:"text,omitempty"`
	Updated  time.Time `bson:"updated"      json:"updated"`
	Position int       `bson:"position"     json:"position,omitempty"`
}
