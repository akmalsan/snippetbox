package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define command-line flag and a default value of port
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command-line flags
	flag.Parse()

	// Use log.New to create a logger for writing info messages.
	// Three parameters: destination to write logs to, string
	// prefix for message, and flags to indicate additional
	// information to include
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create a logger for writing error messages but use stderr
	// as the destination and use the log.Lshortfile flag to
	// include the file name and line number in the log message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	// Initialize a new http.Server struct. Set the Addr and Handler
	// fields so the server uses the same network address and routes as
	// before. Set the ErrorLog field so the server uses the custom
	// errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Value returned from the flag.String function is a pointer
	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe method on the new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
