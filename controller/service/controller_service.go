package service

import "SmartStudyTimer/model"

type MyStudyServicer interface {
	//サービス層までのインタフェース
	//リポジトリは本当のDB操作までしてしまうからやらない
	PostStudyLogService(log model.StudyLog) (model.StudyLog, error)
}
