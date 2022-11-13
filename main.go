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

	// fmt.Println("Before update:")
	// printMap(mapData)

	//update the jsonData - pending!!
	updateData(mapData, target)

	// fmt.Println("\n After update:")
	// printMap(mapData)

	//Convert map data to json
	returnJsonData, error := json.Marshal(mapData)
	if error != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "json")
	response.Write(returnJsonData)
}

func updateData(mapData map[string]interface{}, target string) {

	//delete in the first level
	delete(mapData, target)

	//in the case of nested JSON data
	for _, value := range mapData {

		switch q := value.(type) {

		case []interface{}:
			fmt.Println("value is of array type")
			for _, element := range value.([]interface{}) {
				handleJsonArray(element, target)
			}

		case map[string]interface{}:
			fmt.Println("value is of map type")
			updateData(value.(map[string]interface{}), target)

		default:
			fmt.Printf("value is of type: %T\n", q)
		}

	}

}

func handleJsonArray(element interface{}, target string) {

	switch q := element.(type) {

	case []interface{}:
		fmt.Println("value is of array type")
		for _, ele := range element.([]interface{}) {
			handleJsonArray(ele, target)
		}

	case map[string]interface{}:
		updateData(element.(map[string]interface{}), target)

	default:
		fmt.Printf("value is of type: %T\n", q)

	}
}

func printMap(mapData map[string]interface{}) {

	for key, value := range mapData {
		fmt.Printf("Key is %v, Value is %v and type of value %T \n", key, value, value)
	}
}

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hey there! Welcome Home"))
}
