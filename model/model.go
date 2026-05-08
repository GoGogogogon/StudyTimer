package model

import (
	"time"
)

type StudyLog struct {
	ID        int       `json:"study_id"`
	Title     string    `json:"subject"`
	Time      int       `json:"study_time"`
	CreatedAt time.Time `json:"created_at"`
}
