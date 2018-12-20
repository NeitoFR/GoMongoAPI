package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func AskPassportsRoute(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing URL", err)
		panic(err)
	}
	if r.Form.Get("model_Name") == "" {
		log.Println("no name")
		fmt.Fprintln(w, "{\"error\": \"Model Name required in query parameters\"}")
	} else {

		res, err := _askMqttPassports(r.Form.Get("model_Name"))

		if err != nil {
			log.Println("Error getting passport", err)
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, res)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := r.Form
	log.Println(m)

	keys := make([]string, len(r.Form))
	for key, _ := range r.Form {
		keys = append(keys, key)
	}
	log.Println(keys)
}
func main() {
	godotenv.Load("./.env")

	_initMqttConnection(mqttInfo{os.Getenv("mqtt_server"), os.Getenv("mqtt_port")})

	router := mux.NewRouter()
	router.HandleFunc("/passports", AskPassportsRoute).Methods("GET")
	router.HandleFunc("/test", test).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), router))
	log.Println("Listening on port : " + os.Getenv("port"))
}
