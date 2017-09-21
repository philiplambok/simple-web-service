package controller

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/philiplambok/simple-go-webservice/app/models"
	"github.com/philiplambok/simple-go-webservice/app/tools/json"
)

func Welcome(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		message := Message{}
		message.Set("Method Not Allowed!")

		json.WriteJson(response, message)
		return
	}

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
	if request.Method != "GET" {
		message := Message{}
		message.Set("Method Not Allowed!")

		json.WriteJson(response, message)
		return
	}

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

func CreatePost(response http.ResponseWriter, request *http.Request) {
	// read request body
	body, _ := ioutil.ReadAll(request.Body)

	newPost := struct {
		Post     models.Posts `json:"post"`
		Username string       `json:"username"`
	}{}
	// unmarhal json to object
	json.ReadJson(body, &newPost)

	newPost.Post.SetUserID(newPost.Username)
	newPost.Post.Save()

	message := Message{}
	message.Set("Post Telah Disimpan.")
	json.WriteJson(response, message)

	return
}

func Info(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		message := Message{
			Body: "Method not allowed!",
		}

		json.WriteJson(response, message)
		return
	}

	splited := strings.Split(request.URL.Path, "/")
	var post models.Posts
	post.Title = splited[2]

	result := struct {
		Post models.Posts `json:"post"`
		User models.Users `json:"user"`
	}{}

	var err error
	result.Post, err = post.Find()
	if err != nil {
		return
	}
	result.User, err = result.Post.GetUser()
	if err != nil {
		return
	}

	json.WriteJson(response, result)
	return
}
