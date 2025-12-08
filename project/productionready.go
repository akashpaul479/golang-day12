package project

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Getuser(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "Akash", Email: "Akashpaul479@gmail.com"}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func Setuprouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", Getuser).Methods("GET")
	return r
}
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}
func Productionready() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", Getuser)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
