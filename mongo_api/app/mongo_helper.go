package main

import (
	"encoding/json"
	"log"
	"net/url"

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

func _getPassports(a url.Values) (string, error) {
	var data bson.M
	var err error = nil

	if len(a) > 0 {
		keys := make([]string, len(a))
		for key, _ := range a {
			keys = append(keys, key)
		}
		log.Println(keys, a, len(a))
		// #TODO Parse url.Values with the $and operator and do a filtered Find() query
		err = col.Find(bson.M{}).One(&data)
	} else {
		err = col.Find(bson.M{}).One(&data)
	}
	if err != nil {
		log.Println("Error getting passports", err)
		return "", err
	}
	toByte, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting bson to byte[] ", err)
		return "", err
	}

	return string(toByte[:len(toByte)]), err
}
