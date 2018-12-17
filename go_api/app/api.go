package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetAllRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all passports")
	res := _getAll()

	fmt.Fprintf(w, res)
}

func main() {
	godotenv.Load("./mongo.env")
	_initMongoConnection(mongoInfo{os.Getenv("mongo_url"), os.Getenv("mongo_db"), os.Getenv("mongo_col")})

	router := mux.NewRouter()
	router.HandleFunc("/passports", GetAllRoute).Methods("GET")

	log.Println("Listening on port : " + os.Getenv("port"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), router))
}
