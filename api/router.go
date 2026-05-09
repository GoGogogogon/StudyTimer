package api

import (
	"SmartStudyTimer/controller"
	"SmartStudyTimer/service"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := service.NewMyStudyService(db)
	con := controller.NewStudyTimerController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/api/studylogs", con.SaveController).Methods(http.MethodPost)

	r.HandleFunc("/api/studylogs", con.SelectAllStudylogController).Methods(http.MethodGet)

	r.HandleFunc("/api/studylogs/{id:[0-9]+}", con.UpdateController).Methods(http.MethodPut)

	r.HandleFunc("/api/studylogs/{id:[0-9]+}", con.DeleteController).Methods(http.MethodDelete)

	r.HandleFunc("/api/studylogs/{id:[0-9]+}", con.SelectStudylogController).Methods(http.MethodGet)

	return r
}
