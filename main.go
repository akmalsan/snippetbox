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

func main() {
	// Create a new router
	mux := http.NewServeMux()

	// Handle the root route
	mux.HandleFunc("/", home)

	// Start the server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
