package main

import (
	"goplay/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods("GET")

	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handlers.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books", handlers.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books", handlers.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
