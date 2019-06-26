package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cary205/restful-go/models"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var (
	dao = models.Todo{}
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("AllTodos called")

	defer r.Body.Close()
	var todos []models.Todo
	todos, err := dao.FindAllTodos()
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, todos)
}

func FindTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("FindTodo called")

	vars := mux.Vars(r)
	id := vars["id"]
	result, err := dao.FindTodoById(id)
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, result)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateTodo called")

	defer r.Body.Close()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.Id = bson.NewObjectId()
	if err := dao.InsertTodo(todo); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateTodo called")
	fmt.Fprintln(w, "not implemented !")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTodo called")
	fmt.Fprintln(w, "not implemented !")
}
