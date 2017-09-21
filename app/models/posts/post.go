package posts

import (
	"github.com/philiplambok/simple-go-webservice/app/models/users"
	"github.com/philiplambok/simple-go-webservice/app/tools/db"
	"gopkg.in/mgo.v2/bson"
)

type Posts struct {
	Title  string        `bson:"title" json:"title"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`
}

func All() ([]Posts, error) {
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

func GetUsers(posts []Posts) ([]users.Users, error) {
	session, err := db.Connect()

	if err != nil {
		return nil, err
	}

	usersCollection := session.DB("relation").C("users")

	var dump users.Users
	var users []users.Users

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
