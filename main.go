package main

import (
	"fmt"
	"net/http"
	"gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Dispatch map for CRUD operations.
	router.HandleFunc("/", ClichesAll).Methods("GET")
	router.HandleFunc("/cliches", ClichesAll).Methods("GET")
	router.HandleFunc("/cliches/{id:[0-9]+}", ClichesOne).Methods("GET")

	router.HandleFunc("/cliches", ClichesCreate).Methods("POST")
	router.HandleFunc("/cliches/{id:[0-9]+}", ClichesEdit).Methods("PUT")
	router.HandleFunc("/cliches/{id:[0-9]+}", ClichesDelete).Methods("DELETE")

	http.Handle("/", router) // enable the router

	// Start the server.
	port := ":8888"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, router); // mux.Router now in play
}
