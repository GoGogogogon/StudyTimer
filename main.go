package main

import (
	"SmartStudyTimer/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbname = os.Getenv("DATABASE")

	dbUser = os.Getenv("USER_NAME")

	dbpass = os.Getenv("USERPASS")

	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.3306)%s?parseTime=true", dbname, dbUser, dbpass)
)

func main() {

	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Println("データベースに接続できません")
	}

	defer db.Close()

	if err != nil {
		log.Fatal("接続失敗", err)
	}

	r := api.NewRouter(db)
	// _, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("テーブル作成失敗:", err)
	}

	// fmt.Println("テーブル作成成功（もしくはすでにある)")

	// http.HandleFunc("/api/save", saveHandler)

	// fileServer := http.FileServer(http.Dir("./front"))
	// http.Handle("/", fileServer)

	// fmt.Println("サーバーを起動しました！ http://localhost:8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("サーバー起動失敗:", err)
	}
	log.Fatal(err)
}
