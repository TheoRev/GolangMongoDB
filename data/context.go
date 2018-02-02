package data

import (
	"gopkg.in/mgo.v2"
)

// Context estructura de sesion de mongodb
type Context struct {
	Session *mgo.Session
}

// Close cierra la sesion de coneccion a mongodb
func (c *Context) Close() {
	c.Session.Close()
}

// DBCollection devuelve la colleccion de mongodb
func (c *Context) DBCollection(name string) *mgo.Collection {
	return c.Session.DB(DBName).C(CName)
}

// NewContext crea un objeto de context
func NewContext() *Context {
	session := getSession().Copy()
	c := &Context{
		Session: session,
	}
	return c
}
