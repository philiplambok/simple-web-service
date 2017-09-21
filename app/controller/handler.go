package controller

import (
	"net/http"

	"github.com/philiplambok/restApp/app/utils/json"
	"github.com/philiplambok/simple-go-webservice/app/models/posts"
	"github.com/philiplambok/simple-go-webservice/app/models/users"
)

func Welcome(response http.ResponseWriter, request *http.Request) {

	result := struct {
		Posts []posts.Posts `json:"posts"`
		Users []users.Users `json:"users"`
	}{}
	var err error
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
