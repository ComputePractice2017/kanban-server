package api

import "log"
import "net/http"

// Run starting the server//
func Run() {
	log.Println("start server in port 1000....")
	http.ListenAndServe(":1000", nil)

}
