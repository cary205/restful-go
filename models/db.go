package models

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var GlobalC *mongo.Client

func init() {
	db_url := os.Getenv("MONGODB_URL")
	if len(db_url) == 0 {
		log.Fatalln("no MONGODB_URL provided")
	}

	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_url))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))
	if err != nil {
		log.Fatalln("db Connect error ", err)
	}

	// check db connection
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("db ping error ", err)
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
	coll := GlobalC.Database(db).Collection(collection)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	return cursor.All(context.TODO(), result)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	coll := GlobalC.Database(db).Collection(collection)
	return coll.FindOne(context.TODO(), query).Decode(result)
}

func Insert(db, collection string, docs ...interface{}) error {
	coll := GlobalC.Database(db).Collection(collection)
	_, err := coll.InsertMany(context.TODO(), docs)
	return err
}

func Update(db, collection string, query, update interface{}) error {
	coll := GlobalC.Database(db).Collection(collection)
	result, err := coll.ReplaceOne(context.TODO(), query, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		err = errors.New("No document found")
	}
	return err
}

func Remove(db, collection string, query interface{}) error {
	// ms, c := connect(db, collection)
	// defer ms.Close()
	// return c.Remove(query)
	return nil
}
