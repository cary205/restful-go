package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cary205/restful-go/models"
	"github.com/cary205/restful-go/routes"
)

func main() {
	defer func() {
		log.Println("!!! defer client.Disconnect")
		if err := models.GlobalC.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := routes.NewRouter()

	//func ListenAndServe(addr string, handler Handler) error
	//Handler <=> *Router ### func (*Router) ServeHTTP ###
	log.Fatal(http.ListenAndServe(":8080", r))
}
