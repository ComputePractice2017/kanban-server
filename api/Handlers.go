package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ComputePractice2017/kanban-server/model"
	"github.com/gorilla/mux"
)

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world,Beach")
}
func getAllIssueHandler(w http.ResponseWriter, r *http.Request) {

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

func newIssueHandler(w http.ResponseWriter, r *http.Request) {
	var issue model.Issue
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &issue); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	issue, err = model.NewIssue(issue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(issue); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
func editIssueHandler(w http.ResponseWriter, r *http.Request) {

	var issue model.Issue

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err := r.Body.Close(); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err := json.Unmarshal(body, &issue); err != nil {

		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(err); err != nil {

			w.WriteHeader(http.StatusInternalServerError)

			log.Println(err)

			return

		}

	}

	err = model.EditIssue(issue)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(issue); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

	}

}
func deleteIssueHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := model.DeleteIssue(vars["guid"])

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.WriteHeader(http.StatusOK)

}




func getAllStagesHandler(w http.ResponseWriter, r *http.Request) {

	stages, err := model.GetStages()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err = json.NewEncoder(w).Encode(stages); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func newStageHandler(w http.ResponseWriter, r *http.Request) {
	var stage model.Stage
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &stage); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	stage, err = model.NewStage(stage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(stage); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
func editStageHandler(w http.ResponseWriter, r *http.Request) {

	var stage model.Stage
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err := r.Body.Close(); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	if err := json.Unmarshal(body, &stage); err != nil {

		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(err); err != nil {

			w.WriteHeader(http.StatusInternalServerError)

			log.Println(err)

			return

		}

	}

	err = model.EditStage(stage)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(stage); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

	}

}
func deleteStageHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := model.DeleteStage(vars["guid"])

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		return

	}

	w.WriteHeader(http.StatusOK)

}
