package day12

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Muxserver() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from mux!")

	})
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
