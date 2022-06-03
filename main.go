package main

import (
	"log"
	"net/http"
)

// Define a home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// Write a response to the client
	w.Write([]byte("Hello from Snippetbox!"))
}

// Define a showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Write a response to the client
	w.Write([]byte("Show a specific snippet!"))
}

// Define a createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Write a response to the client
	w.Write([]byte("Create a new snippet!"))
}

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
