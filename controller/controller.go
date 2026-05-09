package controller

import (
	"SmartStudyTimer/controller/service"
	"SmartStudyTimer/model"
	"encoding/json"
	"net/http"
	"strconv"
)

type StudyTimerController struct {
	service service.MyStudyServicer
}

func NewStudyTimerController(s service.MyStudyServicer) *StudyTimerController {
	return &StudyTimerController{service: s}
}

func (c *StudyTimerController) SaveController(w http.ResponseWriter, r *http.Request) {

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

func (c *StudyTimerController) SelectStudylogController(w http.ResponseWriter, r *http.Request) {

	var reqdata model.StudyLog

	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	studydata, err := c.service.SelectStudyLogService(reqdata)

	if err != nil {
		http.Error(w, "データ内部にアクセスできませんでした", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studydata)

}

func (c *StudyTimerController) SelectAllStudylogController(w http.ResponseWriter, r *http.Request) {

	var reqdata []model.StudyLog
	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")

	if limitStr == "" {
		limitStr = "10"
	}

	limit, _ := strconv.Atoi(limitStr)

	studylog, err := c.service.SelectAllStudylogService(reqdata, limit)

	if err != nil {
		http.Error(w, "データ内部にアクセスできませんでした", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studylog)

}
