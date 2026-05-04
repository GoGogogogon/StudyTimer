package model

//学習データ構造体
type StudyLog struct {
	Title string `json:"subject"`
	Time  int    `json:"study_time"`
}
