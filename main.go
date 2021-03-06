package main

import (
	"net/http"

	"github.com/philiplambok/simple-go-webservice/app/controller"
)

func main() {
	port := ":9000"

	server := &http.Server{
		Addr:    port,
		Handler: nil,
	}

	// GET /
	http.HandleFunc("/", controller.Welcome)

	// GET /profile/:username
	http.HandleFunc("/profile/", controller.Profile)

	// GET /posts/:title
	http.HandleFunc("/info/", controller.Info)

	// POST /posts/create
	http.HandleFunc("/posts/create/", controller.CreatePost)

	// DELETE /posts/delete/:title
	http.HandleFunc("/posts/delete/", controller.Delete)

	// PUT /posts/update/:title
	http.HandleFunc("/posts/update/", controller.Update)

	server.ListenAndServe()
}
