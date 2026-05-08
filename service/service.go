package service

import "database/sql"

type MyStudyService struct {
	db *sql.DB
}

func NewMyStudyService(db *sql.DB) *MyStudyService {
	return &MyStudyService{db: db}
}
