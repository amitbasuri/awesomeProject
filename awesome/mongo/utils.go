package mongo

import (
	"gopkg.in/mgo.v2"
)

func NewMongoConnection(mongoConnStr string) *mgo.Session {
	session, err := mgo.Dial(mongoConnStr)
	if err != nil {
		panic(err)
	}
	return session
}
