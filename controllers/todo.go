package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cary205/restful-go/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	todo.Id = primitive.NewObjectID()
	if err := dao.InsertTodo(todo); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateTodo called")

	defer r.Body.Close()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateTodo(todo); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTodo called")

	vars := mux.Vars(r)
	id := vars["id"]
	if err := dao.RemoveTodo(id); err != nil {
		responseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
