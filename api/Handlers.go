package api

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/ComputePractice2017/kanban-server/model"
)

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world,Beach")
}
func getAllIssuesHandler(w http.ResponseWriter, r *http.Request) {

	

	issues, err := model.GetIssues()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err = json.NewEncoder(w).Encode(issues); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}


func newIssueHandler(w http.ResponseWriter, r *http.Request)
{
		var issue model.Issue
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil
		 {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
				return
		}
	if err := r.Body.Close(); err != nil 
	{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &issue); err != nil 
	{
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil 
		{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
		issue, err = model.NewIssue(issue)
		if err != nil 
	{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(issue); err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
}