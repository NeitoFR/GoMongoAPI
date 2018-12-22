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
	err := r.ParseForm()

	if err != nil {
		log.Println("Error parsing URL", err)
		panic(err)
	}
	if r.Form.Get("model_Name") == "" {
		log.Println("Error, no model name")
		fmt.Fprintln(w, "{\"error\": \"model_Name required in query parameters\"}")
		return
	} else {
		// log.Println("Getting passport for model : " + r.Form.Get("model_Name"))
		res, err := _getPassports(r.Form.Get("model_Name"))
		if err != nil {
			log.Println("Error getting passport for model name : "+r.Form.Get("model_Name"), err)
			fmt.Fprintln(w, "{\"error\": \"Error getting passport for model name : "+r.Form.Get("model_Name")+"\"}")
			return
		}
		log.Println("Passport found for model : " + r.Form.Get("model_Name"))
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
	_initMongoConnection(mongoInfo{os.Getenv("mongo_url"), os.Getenv("mongo_db"), os.Getenv("mongo_col")})

	router := mux.NewRouter()
	router.HandleFunc("/passports", GetAllRoute).Methods("GET")
	router.HandleFunc("/test", test).Methods("GET")

	log.Println("Listening on port : " + os.Getenv("port"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), router))
}
