package db

import (
	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}

	return session, nil
}
