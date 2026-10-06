package models

import "time"

type Progress struct {
	UserID       string     `json:"user_id"`
	LessonID     string     `json:"lesson_id"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	LastViewedAt *time.Time `json:"last_viewed_at,omitempty"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	Score        *int       `json:"score,omitempty"`
}
