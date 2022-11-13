package main

import (
	"fmt"
	"io/ioutil"
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

	//get query param
	target := request.URL.Query().Get("target")
	fmt.Println("Target value:", target)

	//handle body
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// fmt.Println(request.Body)
	response.Write([]byte("This is the post route"))
}

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hey there! Welcome Home"))
}
