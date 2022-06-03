package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new router
	mux := http.NewServeMux()
	// Handle the routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	// Start the server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
