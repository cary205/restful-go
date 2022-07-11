package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/cary205/restful-go/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmptyResult struct {
}

type ErrorResult struct {
	Error string `json:"error"`
}

var (
	dao = models.Todo{}
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Cross-Origin Resource Sharing (CORS)
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

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get and check docID
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, ErrorResult{Error: "Invalid OID"})
		return
	}

	// DO
	result, err := dao.FindTodoById(docID)

	// error check
	if err != nil {
		errorResult := ErrorResult{Error: err.Error()}
		if err.Error() == "mongo: no documents in result" {
			responseWithJson(w, http.StatusNotFound, errorResult)
		} else {
			responseWithJson(w, http.StatusInternalServerError, errorResult)
		}
		return
	}

	// ok result
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

	vars := mux.Vars(r)
	id := vars["id"]
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid OID")
		return
	}
	todo.Id = docID

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

	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get and check docID
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseWithJson(w, http.StatusBadRequest, ErrorResult{Error: "Invalid OID"})
		return
	}

	// DO & error check
	if err := dao.RemoveTodo(docID); err != nil {
		errorResult := ErrorResult{Error: err.Error()}
		if err.Error() == "No document found" {
			responseWithJson(w, http.StatusNotFound, errorResult)
		} else {
			responseWithJson(w, http.StatusInternalServerError, errorResult)
		}
		return
	}

	// ok result
	responseWithJson(w, http.StatusOK, EmptyResult{})
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("Test called")

	tmpl := template.Must(template.ParseFiles("page/test.html"))
	tmpl.Execute(w, nil)
}

type Tasks []Task
type Task struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func TaskIndex(w http.ResponseWriter, r *http.Request) {

	tasks := Tasks{
		Task{Name: "Write presentation"},
		Task{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)

	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

}
