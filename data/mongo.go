package data

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	// DBName nombre de la base de datos
	DBName = "golang"
	// CName nombre de la colleccion de mongodb
	CName = "users"
)

var session *mgo.Session

func createDBSession() {
	var err error
	// session,err:=mgo.DialWithInfo(&mgo.DialInfo{
	// 	Addrs:[]string{"",},
	// 	Username:"",
	// 	Password:"",
	// 	Timeout:
	// })

	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
}

func getSession() *mgo.Session {
	if session == nil {
		createDBSession()
	}
	return session
}

func addIndex() {
	userIndex := mgo.Index{
		Key:        []string{"username", "email"},
		Unique:     true,
		Background: true,
	}

	session := getSession().Copy()
	defer session.Close()
	userC := session.DB(DBName).C(CName)
	err := userC.EnsureIndex(userIndex)
	if err != nil {
		log.Fatal(err)
	}
}

// InitData inicia la sesion en mongodb
func InitData() {
	createDBSession()
	addIndex()
}
