package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	_ "github.com/go-sql-driver/mysql"
)



var {
	dbUser := os.Getenv()
}


func main() {


	dsn := "root:root@tcp(127.0.0.1:3306)/study_timer"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("設定エラー：", err)
	}

	defer db.Close()

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Println("待機中")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("接続失敗...:", err)
	}

	fmt.Println("接続成功しました")

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("テーブル作成失敗:", err)
	}

	fmt.Println("テーブル作成成功（もしくはすでにある)")

	http.HandleFunc("/api/save", saveHandler)

	fileServer := http.FileServer(http.Dir("./front"))
	http.Handle("/", fileServer)

	fmt.Println("サーバーを起動しました！ http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバー起動失敗:", err)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

	// 【課題】ここに「保存処理」を書く
	if r.Method != http.MethodPost {
		http.Error(w, "POST以外は受け取りません", http.StatusMethodNotAllowed)
		return
	}

	var a StudyLog

	err := json.NewDecoder(r.Body).Decode((&a))
	if err != nil {
		log.Print("Jsonの読み込みエラー")
		http.Error(w, "読み込みエラーが生じました", http.StatusBadRequest)
		return
	}

	const insertSQL = `INSERT INTO study_logs (subject,study_time) VALUES(?,?)`

	_, err = db.Exec(insertSQL, a.Title, a.Time)

	if err != nil {
		log.Print("dbに格納できません")
		http.Error(w, "データが挿入できません", http.StatusInternalServerError)
		return
	}

	fmt.Println("データ挿入成功")
	w.Write([]byte("保存しました!"))
	fmt.Println("JSからアクセスが来ました") // 確認用ログ
}
