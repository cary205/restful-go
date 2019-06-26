package models

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

var globalS *mgo.Session

func init() {
	db_url := os.Getenv("MONGODB_URL")
	if len(db_url) == 0 {
		log.Fatalln("no MONGODB_URL provided")
	}

	s, err := mgo.Dial(db_url)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalS = s

	log.Println("db inited")
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
