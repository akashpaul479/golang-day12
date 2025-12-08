package day12

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request :%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Finished procressing request: %s %s", r.Method, r.URL.Path)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to home page!"))
}

func Middleware() {
	router := mux.NewRouter()

	router.Use(Loggingmiddleware)
	router.HandleFunc("/", HomeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
