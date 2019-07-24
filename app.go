package main

import (
	"log"
	"net/http"

	"github.com/cary205/restful-go/routes"
)

func main() {
	r := routes.NewRouter()

	//func ListenAndServe(addr string, handler Handler) error
	//Handler <=> *Router ### func (*Router) ServeHTTP ###
	log.Fatal(http.ListenAndServe(":8080", r))
}
