package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/kanban-server/model"
	"github.com/gorilla/mux"
)

// Run starting the server//
func Run() {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(model.InitSession())
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", helloworldHandler).Methods("GET")
	r.HandleFunc("/issues", getAllIssueHandler).Methods("GET")
	r.HandleFunc("/issues", newIssueHandler).Methods("POST")
	r.HandleFunc("/issues/(guid)", editIssueHandler).Methods("PUT")
	r.HandleFunc("/issues/{guid}", deleteIssueHandler).Methods("DELETE")

	r.HandleFunc("/stages", getAllStagesHandler).Methods("GET")
	r.HandleFunc("/stages", newStageHandler).Methods("POST")
	r.HandleFunc("/stages/(guid)", editStageHandler).Methods("PUT")
	r.HandleFunc("/stages/{guid}", deleteStageHandler).Methods("DELETE")



	http.ListenAndServe(":8000", r)
//	db.CreateDBIfNotExist()
//	db.CreateTableIfNotExistStage()
//	db.CreateTableIfNotExistIssue()
}
