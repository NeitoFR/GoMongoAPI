package main

import (
	"encoding/json"
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
		panic(err)
	}
	col = conn.DB(info.db_name).C(info.col_name)
	log.Println("Connection to the mongo instance OK : \nMongo URL : " + info.url + "\nDatabase Name : " + info.db_name + "\nCollection Name : " + info.col_name)
}

func _getPassports(name string) (string, error) {
	var data bson.M
	var err error = nil
	log.Println("Getting passport for model : " + name)

	err = col.Find(bson.M{"model_Name": name}).One(&data)

	if err != nil {
		// log.Println("Error getting passport", err)
		return "", err
	}
	toByte, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting bson to byte[] ", err)
		return "", err
	}
	log.Println("MongoDB Query result : " + string(toByte[:len(toByte)]))
	return string(toByte[:len(toByte)]), err
}
