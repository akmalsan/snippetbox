package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define command-line flag and a default value of port
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command-line flags
	flag.Parse()

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

	// Value returned from the flag.String function is a pointer
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
