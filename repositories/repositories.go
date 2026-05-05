package repositories

import (
	"SmartStudyTimer/model"
	"database/sql"
)

// 学習データをinsertする関数
func SaveStudyDate(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

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
