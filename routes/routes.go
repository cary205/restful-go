package routes

import (
	"log"
	"net/http"

	"github.com/cary205/restful-go/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

var routes []Route

func init() {
	register("GET", "/todo", controllers.AllTodos)
	register("GET", "/todo/{id}", controllers.FindTodo)
	register("POST", "/todo", controllers.CreateTodo)
	register("PUT", "/todo", controllers.UpdateTodo)
	register("DELETE", "/todo", controllers.DeleteTodo)

	log.Println("routes inited")

}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc) {
	routes = append(routes, Route{method, pattern, handler})
}
