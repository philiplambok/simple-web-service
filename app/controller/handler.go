package controller

import (
	"net/http"
	"strings"

	"github.com/philiplambok/simple-go-webservice/app/models"
	"github.com/philiplambok/simple-go-webservice/app/tools/json"
)

func Welcome(response http.ResponseWriter, request *http.Request) {

	result := struct {
		Posts []models.Posts `json:"posts"`
		Users []models.Users `json:"users"`
	}{}
	var err error
	var posts models.Posts

	result.Posts, err = posts.All()

	if err != nil {
		message := Message{
			Body: err.Error(),
		}

		json.WriteJson(response, message)
		return
	}
	result.Users, err = posts.GetUsers(result.Posts)

	if err != nil {
		message := Message{}
		message.Set(err.Error())
		json.WriteJson(response, message)
		return
	}

	json.WriteJson(response, result)
	return
}

func Profile(response http.ResponseWriter, request *http.Request) {
	result := struct {
		User  models.Users   `json:"user"`
		Posts []models.Posts `json:"posts"`
	}{}
	var user models.Users

	// parse username from url
	splited := strings.Split(request.URL.Path, "/")
	user.Username = splited[2]

	// get data user from username
	var err error
	result.User, err = user.Get()

	if err != nil {
		message := Message{}
		message.Set(err.Error())
		json.WriteJson(response, message)
		return
	}

	// get post
	result.Posts, err = result.User.GetPosts()
	if err != nil {
		message := Message{}
		message.Set(err.Error())
		json.WriteJson(response, message)
		return
	}

	json.WriteJson(response, result)
	return
}
