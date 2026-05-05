package model

import (
	"database/sql"
	"time"
)

type StudyLog struct {
	ID        int       `json:"study_id"`
	Title     string    `json:"subject"`
	Time      int       `json:"study_time"`
	CreatedAt time.Time `json:"created_at"`
}

type TimerService struct {
	db *sql.DB
}

func NewTimerService(db *sql.DB) *TimerService {
	return &TimerService{db: db}
}
