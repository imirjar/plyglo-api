package models

import "time"

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LogoPath    *string   `json:"logo_path,omitempty"`
	IsPublished bool      `json:"is_published"`
	Updated     time.Time `json:"updated,omitempty"`
}

type Chapter struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Position    int       `json:"position"`
	Updated     time.Time `json:"updated,omitempty"`
}

type Lesson struct {
	ID        string    `json:"id"`
	ChapterID string    `json:"chapter_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text,omitempty"`
	Position  int       `json:"position"`
	Updated   time.Time `json:"updated,omitempty"`
}
