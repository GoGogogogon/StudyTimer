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
	newlog, err := repositories.SaveStudyDate(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}

	return newlog, nil
}

func (s *MyStudyService) UpdateStudyLogService(log model.StudyLog) (model.StudyLog, error) {
	newlog, err := repositories.UpdateStudyData(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newlog, nil
}

func (s *MyStudyService) DeleteStudyLogService(id int) (model.StudyLog, error) {
	newlog, err := repositories.DeleteStudyData(s.db, id)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newlog, nil
}

func (s *MyStudyService) SelectStudyLogService(id int) (model.StudyLog, error) {
	newlog, err := repositories.SelectStudyData(s.db, id)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newlog, nil
}

func (s *MyStudyService) SelectAllStudylogService(limit int) ([]model.StudyLog, error) {
	loglist, err := repositories.AllSelectStudyData(s.db, limit)
	if err != nil {
		return []model.StudyLog{}, err
	}
	return loglist, nil
}
