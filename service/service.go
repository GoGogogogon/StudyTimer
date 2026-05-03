package service

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func saveHandler(w http.ResponseWriter, r *http.Request) {

// 	// 【課題】ここに「保存処理」を書く
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "POST以外は受け取りません", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var a StudyLog

// 	err := json.NewDecoder(r.Body).Decode((&a))
// 	if err != nil {
// 		log.Print("Jsonの読み込みエラー")
// 		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
// 		return
// 	}

// 	const insertSQL = `INSERT INTO study_logs (subject,study_time) VALUES(?,?)`

// 	_, err = db.Exec(insertSQL, a.Title, a.Time)

// 	if err != nil {
// 		log.Print("dbに格納できません")
// 		http.Error(w, "データが挿入できません", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Println("データ挿入成功")
// 	w.Write([]byte("保存しました!"))
// 	fmt.Println("JSからアクセスが来ました") // 確認用ログ
// }
//
