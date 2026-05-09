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

	//データを保存

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode((&reqdata)); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	studyLog, err := c.service.PostStudyLogService(reqdata)

	if err != nil {
		http.Error(w, "データ内部に入れませんでした", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studyLog)

}

func (c *StudyTimerController) UpdateController(w http.ResponseWriter, r *http.Request) {
	//誤ってデータを保存したときに値を修正する用

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	studylog, err := c.service.UpdateStudyLogService(reqdata)

	if err != nil {
		http.Error(w, "データ内部に入れませんでした", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studylog)
}

func (c *StudyTimerController) DeleteController(w http.ResponseWriter, r *http.Request) {

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	studylog, err := c.service.DeleteStudyLogService(reqdata)

	if err != nil {
		http.Error(w, "データ内部に入れませんでした", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studylog)
}
