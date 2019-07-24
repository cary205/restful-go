package routes

import (
	"log"
	"net/http"

	"github.com/cary205/restful-go/controllers"
	"github.com/gorilla/mux"
)

type myRoute struct {
	method  string
	pattern string
	//type HandlerFunc func(ResponseWriter, *Request)
	handlerFunc http.HandlerFunc
}

var myRoutes []myRoute

func init() {
	register("GET", "/todos", controllers.AllTodos)
	register("GET", "/todos/{id}", controllers.FindTodo)
	register("POST", "/todos", controllers.CreateTodo)
	register("PUT", "/todos", controllers.UpdateTodo)
	register("DELETE", "/todos/{id}", controllers.DeleteTodo)

	log.Println("routes inited")

}

func NewRouter() *mux.Router {
	//func NewRouter() *Router
	r := mux.NewRouter()

	for _, route := range myRoutes {
		//func (*Router) HandleFunc
		r.HandleFunc(route.pattern, route.handlerFunc).Methods(route.method)
	}

	return r
}

func register(method, pattern string, handlerFunc http.HandlerFunc) {
	myRoutes = append(myRoutes, myRoute{method, pattern, handlerFunc})
}
