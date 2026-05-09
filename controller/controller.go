package controller

import (
	"SmartStudyTimer/model"
	"SmartStudyTimer/service"
	"encoding/json"
	"net/http"
)

type StudyTimerController struct {
	service service.MyStudyService
}

func NewStudyTimerController(s service.MyStudyService) *StudyTimerController {
	return &StudyTimerController{service: s}
}

func (c *StudyTimerController) SaveContnroller(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "POST以外は受け取りません", http.StatusMethodNotAllowed)
		return
	}

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode((&reqdata)); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
	}
}

//func UpdateHandler(w http.ResponseWriter, r *http.Request)
