package day12

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Id   int
	Name string
	Age  int
}

var users []user
var lastID = 0

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u user
	json.NewDecoder(r.Body).Decode(&u)

	lastID++
	u.Id = lastID
	users = append(users, u)

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func Bodyhandling() {
	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUser).Methods("POST")

	fmt.Println("server running on:8080")
	http.ListenAndServe(":8080", r)
}
