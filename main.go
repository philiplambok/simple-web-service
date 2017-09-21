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

	http.HandleFunc("/", controller.Welcome)

	server.ListenAndServe()
}
