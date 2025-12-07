package day12

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Getuser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Fprintf(w, "Users ID = %s", id)
}
func Pathparameters() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", Getuser).Methods(("Get"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from mux!")
	})
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)

}
