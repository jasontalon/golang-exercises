package main

import (
	"fmt"
	"net/http"

	//"encoding/json"

	"github.com/gorilla/mux"
	"github.com/jasontalon/golang-exercises/handlers"
)

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/exercises").Subrouter()
	s.HandleFunc("/hello-world", handlers.HelloWorld).Methods("GET")
	s.HandleFunc("/acronym", handlers.Acronym).Methods("POST")

	port := ":8080"

	fmt.Printf("listens to port %s\n", port)
	http.ListenAndServe(port, r)
}
