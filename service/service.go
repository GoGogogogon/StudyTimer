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

func (s *MyStudyService) DeleteStudyLogService(log model.StudyLog) (model.StudyLog, error) {
	newlog, err := repositories.DeleteStudyData(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newlog, nil
}

func (s *MyStudyService) SelectStudyLogService(log model.StudyLog) (model.StudyLog, error) {
	newlog, err := repositories.SelectStudyData(s.db, log)
	if err != nil {
		return model.StudyLog{}, err
	}
	return newlog, nil
}

func (s *MyStudyService) SelectAllStudylogService(log []model.StudyLog, limit int) ([]model.StudyLog, error) {
	loglist, err := repositories.AllSelectStudyData(s.db, log, limit)
	if err != nil {
		return []model.StudyLog{}, err
	}
	return loglist, nil
}
