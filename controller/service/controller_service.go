package service

import "SmartStudyTimer/model"

type MyStudyServicer interface {
	//コントローラ―→サービス層　までのインタフェース
	//コントローラ―→サービス→リポジトリだと
	//本当のDB操作までしてしまう（テストが面倒）からやらない

	PostStudyLogService(log model.StudyLog) (model.StudyLog, error)
}
