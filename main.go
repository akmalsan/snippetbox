package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// Add URL path check
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Write a response to the client
	w.Write([]byte("Hello from Snippetbox!"))
}

// Define a showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the ID parameter from the query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID: %d...", id)
}

// Define a createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Restrict the function to only POST requests
	if r.Method != "POST" {
		// Include Allow POST header to the response
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
