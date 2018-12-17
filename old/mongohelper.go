package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type mongoInfo struct {
	url      string
	db_name  string
	col_name string
}

var conn *mgo.Session
var col *mgo.Collection

func _initMongoConnection(info mongoInfo) {
	conn, err := mgo.Dial(info.url)
	if err != nil {
		log.Fatalf("Problem connecting to Mongo %s", err)
	}
	col = conn.DB(info.db_name).C(info.col_name)
	log.Println("Connection to the mongo instance OK : \nMongo URL : " + info.url + "\nDatabase Name :" + info.db_name + "\nCollection Name : " + info.col_name)
}

func _getAllPeople() []bson.M {
	var data []bson.M

	col.Find(bson.M{}).All(&data)

	return data
}
