package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/testres", testres)
	http.ListenAndServe(":7500", router)

}
func test(w http.ResponseWriter, r *http.Request) {
	user := User{name: "xyz", age: 3}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func testres(w http.ResponseWriter, r *http.Request) {
	user := User{name: "xyz", age: 3}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

}
