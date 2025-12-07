package day12

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Listuser(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	fmt.Fprintf(w, "page=%s Limit=%s", page, limit)
}
func Queryparameters() {
	r := mux.NewRouter()
	r.HandleFunc("/users", Listuser).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from mux!")
	})
	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", r)
}
