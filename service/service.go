package service

import (
	"SmartStudyTimer/model"
	"SmartStudyTimer/repositories"
	"database/sql"
)

type MyStudyService struct {
	db *sql.DB
}

func NewMyStudyService(db *sql.DB) *MyStudyService {
	return &MyStudyService{db: db}
}

func (s *MyStudyService) PostStudyLogService(log model.StudyLog) (model.StudyLog, error) {
	newLog, err := repositories.SaveStudyDate(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}

	return newLog, nil
}

func (s *MyStudyService) UpdateStudyLogService(log model.StudyLog) (model.StudyLog, error) {
	newLog, err := repositories.UpdateStudyData(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newLog, err
}
