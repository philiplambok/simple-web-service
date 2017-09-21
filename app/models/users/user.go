package users

import (
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
}
