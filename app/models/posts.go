package models

import (
	"github.com/philiplambok/simple-go-webservice/app/tools/db"
	"gopkg.in/mgo.v2/bson"
)

type Posts struct {
	Title  string        `bson:"title" json:"title"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`
}

func (p Posts) All() ([]Posts, error) {
	session, err := db.Connect()

	if err != nil {
		return nil, err
	}

	postsCollection := session.DB("relation").C("posts")

	var posts []Posts
	err = postsCollection.Find(nil).All(&posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p Posts) GetUsers(posts []Posts) ([]Users, error) {
	session, err := db.Connect()

	if err != nil {
		return nil, err
	}

	usersCollection := session.DB("relation").C("users")

	var dump Users
	var users []Users

	for i, _ := range posts {
		selector := bson.M{"_id": posts[i].UserID}
		users = append(users, dump)
		err = usersCollection.Find(selector).One(&users[i])

		if err != nil {
			return nil, err
		}

	}

	return users, nil
}

func (p *Posts) SetUserID(username string) {
	session, err := db.Connect()

	if err != nil {
		return
	}

	userCollection := session.DB("relation").C("users")
	var user Users
	selector := bson.M{"username": username}
	err = userCollection.Find(selector).One(&user)
	p.UserID = user.ID
}

func (p Posts) Save() {
	session, err := db.Connect()

	if err != nil {
		return
	}

	postsCollection := session.DB("relation").C("posts")
	err = postsCollection.Insert(&p)

	if err != nil {
		return
	}

	return
}
