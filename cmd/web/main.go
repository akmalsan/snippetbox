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

	// Create a file server which servers files from the
	// "./ui/static" directory. Path to http.Dir function
	// is relative to the directory root
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Use the mux.Handle function to register the file server
	// as the handler for all URL paths that start with "/static".
	// Strip the "/static" prefix before it reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Start the server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
