package users

import (
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}
