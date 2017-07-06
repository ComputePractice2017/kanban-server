package model
import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)
func InitSesson() error

type Issue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var session *r.Session

func InitSesson() error{
	var err error
	Session, err := r.Connect(r.ConnectOpts
	{
		Address: "localhost",
	})
 return err
}

func GetIssues() ([]Issue, error) {
	res, err :=  r.DB("kanban").Table("Issues").Run(session)
	if err != nil {
		return nil,err
	}

	var response []Issue
	err = res.All(&response)
	if err != nil {
		return nil,err
	}
	return response,nil

}
func NewPerson(i Issue) (Issue, error) {

	res, err := r.UUID().Run(session)

	if err != nil {

		return i, err

	}



	var UUID string

	err = res.One(&UUID)

	if err != nil {

		return i, err

	}



	p.ID = UUID



	res, err = r.DB("kanban").Table("issues").Insert(i).Run(session)

	if err != nil {

		return i, err

	}



	return i, nil

}