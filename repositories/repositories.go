package repositories

import (
	"SmartStudyTimer/model"
	"database/sql"
)

func SaveStudyDate(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

	// 学習データをinsertする関数

	const InsertSQL = `
	insert into study_logs (subject,study_time,created_time) values
	(?,?,now())
	`

	result, err := db.Exec(InsertSQL, log.Title, log.Time)

	//sql.resultはIDのこと

	if err != nil {
		return model.StudyLog{}, err
	}

	Id, _ := result.LastInsertId()

	log.ID = int(Id)
	//IDの自動確保
	return log, nil
}

func UpdateStudyData(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

	const updatesql = `
	update study_logs set subject = ? ,study_time = ?
	where study_id = ? 
	`

	_, err := db.Exec(updatesql, log.Title, log.Time, log.ID)

	if err != nil {
		return model.StudyLog{}, err
	}

	return log, nil
}

func DeleteStudyData(db *sql.DB, log model.StudyLog) (model.StudyLog, error) {

	const deletesql = `
	delete from study_logs 
	where study_id = ?
	 `

	_, err := db.Exec(deletesql, log.ID)

	if err != nil {
		return model.StudyLog{}, err
	}

	return log, nil
}
