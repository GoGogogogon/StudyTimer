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

	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbpass, dbname)
)

func main() {

	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Println("データベースに接続できません")
		return
	}

	defer db.Close()

	r := api.NewRouter(db)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("サーバー起動失敗:", err)
	}

}
