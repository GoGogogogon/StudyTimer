package api

import (
	"SmartStudyTimer/service"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/save", service.SaveHandler).Methods(http.MethodPost)

	return r
}
