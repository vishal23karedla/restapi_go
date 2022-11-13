package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("\n Main funcion is running")

	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/", postHome).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func postHome(response http.ResponseWriter, request *http.Request) {

	//get query param
	target := request.URL.Query().Get("target")

	//parse body
	reqBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	//convert the json data into a map
	var mapData map[string]interface{}
	json.Unmarshal(reqBody, &mapData)

	//update the jsonData
	updateData(mapData, target)

	//Convert map data to json
	updatedJsonData, error := json.Marshal(mapData)
	if error != nil {
		panic(error)
	}

	fmt.Println("All the instances of", target, "are deleted :) \n")

	//return the updated JSON
	response.Header().Set("Content-Type", "json")
	response.Write(updatedJsonData)
}

func updateData(mapData map[string]interface{}, target string) {

	//delete the target in current level
	delete(mapData, target)

	//in the case of nested JSON data
	for _, value := range mapData {

		switch value.(type) {

		// when the value is an array
		case []interface{}:
			for _, element := range value.([]interface{}) {
				handleJsonArray(element, target)
			}

		//when the value is a json object in itself
		case map[string]interface{}:
			updateData(value.(map[string]interface{}), target)

		default:
			// Do nothing :P
		}

	}
}

// Each element of an array is passed to this function since the element can be of any type
func handleJsonArray(element interface{}, target string) {

	switch element.(type) {

	//when the element is again an array
	case []interface{}:
		for _, ele := range element.([]interface{}) {
			handleJsonArray(ele, target)
		}

	//when the element is a json object
	case map[string]interface{}:
		updateData(element.(map[string]interface{}), target)

	default:
		// Do nothing :P
	}

}

// Prints the elements(key,value) of the map and their types
func printMap(mapData map[string]interface{}) {

	for key, value := range mapData {
		fmt.Printf("Key is %v, Value is %v and Type of value %T \n", key, value, value)
	}
}

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hey there! I am a newbie GOpher"))
}
