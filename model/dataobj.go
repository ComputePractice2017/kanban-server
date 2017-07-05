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
