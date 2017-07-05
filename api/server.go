package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run starting the server//
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloworldHandler).Methods("GET")
	log.Println("start server in port 8000....")
	http.ListenAndServe(":8000", r)

}
