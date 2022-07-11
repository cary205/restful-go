package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Content string             `bson:"content" json:"content"`
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

func (m *Todo) FindTodoById(docID primitive.ObjectID) (Todo, error) {
	log.Println("FindTodoById called")

	var result Todo
	err := FindOne(db, collection, bson.M{"_id": docID}, nil, &result)
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

func (m *Todo) RemoveTodo(docID primitive.ObjectID) error {
	log.Println("RemoveTodo called")

	return Remove(db, collection, bson.M{"_id": docID})
}
