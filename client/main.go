package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// Get Request for "/test" route

	fmt.Println("\nSending GET request to route: /test ")

	resp, _ := http.Get("http://localhost:7500/test")
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response from Server: ", string(body)) // print response body from server

	// send request to server with data - Post Request for "/testres" route
	fmt.Println("\nSending POST request to route: /testres ")

	var payload User
	payload.Name = "John Doe"
	payload.Age = 10

	json_data, _ := json.Marshal(payload) // Encode structure to JSON

	_, err := http.Post("http://localhost:7500/testres", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("error sending request")
	} else {
		fmt.Println("Request Sent Successfully")
	}

}
