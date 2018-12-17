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
	log.Println("Connection to the mongo instance OK : \nMongo URL : " + info.url + "\nDatabase Name :" + info.db_name + "\nCollection Name : " + info.col_name)
}

func _getAll() string {
	var data []bson.M

	err := col.Find(bson.M{}).All(&data)
	if err != nil {
		log.Println("Error getting passports", err)
		panic(err)
	}
	log.Println(data)
	res_json, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting bson to byte[] ", err)
		panic(err)
	}

	return string(res_json[:len(res_json)])
}
