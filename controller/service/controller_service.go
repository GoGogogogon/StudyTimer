package service

import (
	"SmartStudyTimer/model"
)

type MyStudyServicer interface {
	//コントローラ―→サービス層　までのインタフェース
	//コントローラ―→サービス→リポジトリだと
	//本当のDB操作までしてしまう（テストが面倒）からやらない

	PostStudyLogService(log model.StudyLog) (model.StudyLog, error)
	UpdateStudyLogService(log model.StudyLog) (model.StudyLog, error)
	DeleteStudyLogService(log model.StudyLog) (model.StudyLog, error)
	SelectStudyLogService(log model.StudyLog) (model.StudyLog, error)
	SelectAllStudylogService(log []model.StudyLog, limit int) ([]model.StudyLog, error)
}
