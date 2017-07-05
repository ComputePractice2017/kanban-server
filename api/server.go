package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/kanban-server/model"
	"github.com/gorilla/mux"
)

// Run starting the server//
func Run() {
	log.Primtln("Connecting to rethinkDB on localhost...")
	err := model.InitSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(model.InitSession())
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", helloworldHandler).Methods("GET")
	r.HandleFunc("/issues", getAllIssuesHandler).Methods("GET")

	http.ListenAndServe(":8000", r)
	db.CreateDBIfNotExist()
	db.CreateTableIfNotExistStage()
	db.CreateTableIfNotExistIssue()
}
