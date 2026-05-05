package controller

import (
	"SmartStudyTimer/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TimerController struct {
	db *sql.DB
}

func NewTimerService(db *sql.DB) *TimerController {
	return &TimerController{db: db}
}

func (c *TimerController) SaveHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "POST以外は受け取りません", http.StatusMethodNotAllowed)
		return
	}

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode((&reqdata)); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
	}

	//data, err = repositories.SaveStudyDate(db)
}

//func UpdateHandler(w http.ResponseWriter, r *http.Request)
