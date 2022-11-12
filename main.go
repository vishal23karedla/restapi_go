package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main funcion is running")

	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/", postHome).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func postHome(response http.ResponseWriter, request *http.Request) {
	//Handle the JSON object

	// fmt.Println(request.Body)

	response.Write([]byte("This is the post route"))
}

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hey there! Welcome Home"))
}
