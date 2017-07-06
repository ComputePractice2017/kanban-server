package model

import (
	"log"
	"os"

	r "gopkg.in/gorethink/gorethink.v3"
)

type Issue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Stage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var session *r.Session

func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "kanban" {
			return nil
		}
	}
	_, err = r.DBCreate("kanban").Run(session)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableIfNotExist1() error {
	res, err := r.DB("kanban").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "issues" {
			return nil
		}
	}

	_, err = r.DB("kanban").TableCreate("issues", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)

	return err
}
func CreateTableIfNotExist2() error {
	res, err := r.DB("kanban").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "stages" {
			return nil
		}
	}

	_, err = r.DB("kanban").TableCreate("stages", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)

	return err
}
func InitSession() error {

	dbaddress := os.Getenv("RETHINKDB_HOST")

	if dbaddress == "" {

		dbaddress = "localhost"

	}
	log.Printf("RETHINKDB_HOST: %s\n", dbaddress)

	var err error

	session, err = r.Connect(r.ConnectOpts{

		Address: dbaddress,
	})

	if err != nil {

		return err

	}
	err = CreateDBIfNotExist()

	if err != nil {

		return err

	}
	err = CreateTableIfNotExist1()
	if err != nil {

		return err
	}
	err = CreateTableIfNotExist2()
	return err

}
func GetIssues() ([]Issue, error) {
	res, err := r.DB("kanban").Table("issues").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Issue
	err = res.All(&response)
	if err != nil {
		return nil, err
	}
	return response, nil

}
func NewIssue(i Issue) (Issue, error) {

	res, err := r.UUID().Run(session)

	if err != nil {

		return i, err

	}

	var UUID string

	err = res.One(&UUID)

	if err != nil {

		return i, err

	}
	i.ID = UUID

	res, err = r.DB("kanban").Table("issues").Insert(i).Run(session)

	if err != nil {

		return i, err

	}

	return i, nil

}
func EditIssue(i Issue) error {

	_, err := r.DB("kanban").Table("issue").Get(i.ID).Replace(i).Run(session)

	if err != nil {

		return err

	}

	return nil

}
func DeleteIssue(id string) error {

	_, err := r.DB("kanban").Table("issues").Get(id).Delete().Run(session)

	if err != nil {

		return err

	}

	return nil

}

func GetStages() ([]Stage, error) {
	res, err := r.DB("kanban").Table("stages").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Stage
	err = res.All(&response)
	if err != nil {
		return nil, err
	}
	return response, nil

}
func NewStage(s Stage) (Stage, error) {

	res, err := r.UUID().Run(session)

	if err != nil {

		return s, err

	}

	var UUID string

	err = res.One(&UUID)

	if err != nil {

		return s, err

	}
	s.ID = UUID

	res, err = r.DB("kanban").Table("stages").Insert(s).Run(session)

	if err != nil {

		return s, err

	}

	return s, nil

}
func EditStage(s Stage) error {

	_, err := r.DB("kanban").Table("stages").Get(s.ID).Replace(s).Run(session)

	if err != nil {

		return err

	}

	return nil

}
func DeleteStage(id string) error {

	_, err := r.DB("kanban").Table("stages").Get(id).Delete().Run(session)

	if err != nil {

		return err

	}

	return nil

}
