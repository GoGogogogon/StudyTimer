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

func DeleteStudyData(db *sql.DB, id int) (model.StudyLog, error) {

	const deletesql = `
	delete from study_logs 
	where study_id = ?
	 `

	var log model.StudyLog
	_, err := db.Exec(deletesql, id)

	if err != nil {
		return model.StudyLog{}, err
	}

	return log, nil
}

func SelectStudyData(db *sql.DB, id int) (model.StudyLog, error) {
	//指定したIDのデータを返す関数
	const selectsql = `
	select * from study_logs
	where study_id = ?
	`

	//指定したIDの一行データ
	row := db.QueryRow(selectsql, id)

	var log model.StudyLog
	var createdtime sql.NullTime

	if err := row.Scan(&log.ID, &log.Title, &log.Time, &createdtime); err != nil {
		return model.StudyLog{}, err
	}

	if createdtime.Valid {
		log.CreatedAt = createdtime.Time
	}

	return log, nil
}

func AllSelectStudyData(db *sql.DB, limit int) ([]model.StudyLog, error) {
	//limitまでの学習データを取得
	const allselectsql = `
	select * from study_logs
	limit ?
	`

	rows, err := db.Query(allselectsql, limit)

	if err != nil {
		return []model.StudyLog{}, err
	}

	defer rows.Close()

	var loglist []model.StudyLog
	// rowが終わるまでloglistにnewlogを追加
	for rows.Next() {
		var newlog model.StudyLog
		var createdtime sql.NullTime
		if err := rows.Scan(&newlog.ID, &newlog.Title, &newlog.Time, &createdtime); err != nil {
			return []model.StudyLog{}, err
		}

		loglist = append(loglist, newlog)
	}

	return loglist, nil
}
