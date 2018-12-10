package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetAllPeoplesRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all peoples")
	res := _getAllPeople()
	log.Println("Got " + strconv.Itoa(len(res)) + " people")
	// json_res := json.Marshal(res[0])
	fmt.Fprintln(w, res)
}

func GetPeopleByName() {

}
func main() {
	godotenv.Load("./mongo.env")
	_initMongoConnection(mongoInfo{os.Getenv("mongo_url"), os.Getenv("mongo_db"), os.Getenv("mongo_col")})

	router := mux.NewRouter()
	router.HandleFunc("/peoples", GetAllPeoplesRoute).Methods("GET")

	log.Println("Listening on port : " + os.Getenv("port"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), router))
}
