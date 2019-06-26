package main

import (
	"net/http"

	"github.com/cary205/restful-go/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":8080", r)
}
