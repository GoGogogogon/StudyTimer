package repositories

import (
	"SmartStudyTimer/model"
	"database/sql"
)

func SaveStudyDate(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

	// 学習データをinsertする関数

	const InsertSQL = `
	insert into study_logs (subject,study_time,created_time) values
	(?,?,now());
	`

	var Newlog model.StudyLog

	Newlog.Title, Newlog.Time = log.Title, log.Time

	result, err := db.Exec(InsertSQL, log.Title, log.Time)

	//sql.resultはIDのこと

	if err != nil {
		return model.StudyLog{}, err
	}

	Id, _ := result.LastInsertId()

	Newlog.ID = int(Id)
	//IDの自動確保
	return Newlog, nil
}

func UpdateStudyData(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

	const updatesql = `
	update study_logs set subject = ? study_time = ?
	where study_id = ? ;`

	result, err := db.Exec(updatesql, log.Title, log.Time)
}
