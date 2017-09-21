package models

import (
	"github.com/philiplambok/simple-go-webservice/app/tools/db"
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
}

func (u Users) Get() (Users, error) {
	session, err := db.Connect()

	if err != nil {
		return Users{}, err
	}

	// set collection
	usersCollection := session.DB("relation").C("users")

	// get data user
	selector := bson.M{"username": u.Username}
	err = usersCollection.Find(selector).One(&u)

	if err != nil {
		return Users{}, err
	}

	return u, nil
}

func (u Users) GetPosts() ([]Posts, error) {
	session, err := db.Connect()

	if err != nil {
		return nil, err
	}

	postsCollection := session.DB("relation").C("posts")

	// get posts
	var posts []Posts
	selector := bson.M{"user_id": u.ID}
	err = postsCollection.Find(selector).All(&posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (u Users) IsAny() (bool, error) {
	session, err := db.Connect()
	if err != nil {
		return false, err
	}

	usersCollection := session.DB("relation").C("users")

	selector := bson.M{"username": u.Username, "password": u.Password}
	err = usersCollection.Find(selector).One(&u)

	// not auth
	if err != nil {
		return false, nil
	}

	return true, nil
}
