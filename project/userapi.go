package project

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    int
	Name  string
	Email string
}

var users = []user{
	{ID: 1, Name: "Akash", Email: "akashpaul479@gmail.com"},
	{ID: 2, Name: "kushal", Email: "kushal6447@gmail.com"},
}

func Getuser1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	if query != "" {
		var filtered []user
		for _, u := range users {
			if u.Name == query {
				filtered = append(filtered, u)
			}
		}
		json.NewEncoder(w).Encode(filtered)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func GetuserbyID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, u := range users {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
func Setuprouter1() *mux.Router {
	r := mux.NewRouter()
	r.Use(Logging)
	r.HandleFunc("/users", Getuser).Methods("GET")
	r.HandleFunc("/users/{id}", GetuserbyID).Methods("GET")
	return r
}
func Logging1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func UseApi() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", Getuser1)
	fmt.Println("server running on:8080")
	http.ListenAndServe(":8080", mux)
}
