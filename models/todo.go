package models

import (
	"errors"
	"log"

	"github.com/globalsign/mgo/bson"
)

type Todo struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Content string        `bson:"content" json:"content"`
}

const (
	db         = "todo_data"
	collection = "todo"
)

func (m *Todo) FindAllTodos() ([]Todo, error) {
	log.Println("FindAllTodos called")

	var result []Todo
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (m *Todo) FindTodoById(id string) (Todo, error) {
	log.Println("FindTodoById called")

	var result Todo

	if !bson.IsObjectIdHex(id) {
		return result, errors.New("Invalid OID")
	}

	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (m *Todo) InsertTodo(todo Todo) error {
	log.Println("InsertTodo called")

	return Insert(db, collection, todo)
}

func (m *Todo) UpdateTodo(todo Todo) error {
	log.Println("UpdateTodo called")

	return Update(db, collection, bson.M{"_id": todo.Id}, todo)
}

func (m *Todo) RemoveTodo(id string) error {
	log.Println("RemoveTodo called")

	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid OID")
	}

	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
