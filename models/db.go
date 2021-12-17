package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GlobalC *mongo.Client

func init() {
	// db_url := os.Getenv("MONGODB_URL")
	db_url := "mongodb://webuser:web987@192.168.247.129:27017/todo_data?authSource=admin"
	if len(db_url) == 0 {
		log.Fatalln("no MONGODB_URL provided")
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_url))
	if err != nil {
		log.Fatalln("db Connect error ", err)
	}
	GlobalC = client

	log.Println("db inited")
}

// func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
// 	s := globalS.Copy()
// 	c := s.DB(db).C(collection)
// 	return s, c
// }

func FindAll(db, collection string, query, selector, result interface{}) error {
	// ms, c := connect(db, collection)
	// defer ms.Close()
	// return c.Find(query).Select(selector).All(result)
	return nil
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	coll := GlobalC.Database(db).Collection(collection)
	return coll.FindOne(context.TODO(), query).Decode(result)
}

func Insert(db, collection string, docs ...interface{}) error {
	// ms, c := connect(db, collection)
	// defer ms.Close()
	// return c.Insert(docs...)
	return nil
}

func Update(db, collection string, query, update interface{}) error {
	// ms, c := connect(db, collection)
	// defer ms.Close()
	// return c.Update(query, update)
	return nil
}

func Remove(db, collection string, query interface{}) error {
	// ms, c := connect(db, collection)
	// defer ms.Close()
	// return c.Remove(query)
	return nil
}
