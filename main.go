package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// struct field name has json tag but is not exported.

// if json tags are used for struct fields.. then the field has to be exported ie, field name to be started with capital letter)
type User struct {
	Name string `json:"name"` // name -> Name
	Age  int    `json:"age"`  // age -> Age
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)
	router.HandleFunc("/testres", testres)

	http.ListenAndServe(":7500", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nRequest received from client at route: /test")

	user := User{Name: "xyz", Age: 3}
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Sending Response to Client...")

	json.NewEncoder(w).Encode(user)

}

func testres(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nRequest received from client at route: /testres")

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("Decoded Message from Client:", user)
}
