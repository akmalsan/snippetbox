package main

import (
	"fmt"
	"html/template"
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
	// Create a slice containing the paths to the two files
	// home.page.tmpl must be the first file in the slice
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	// Use the template.Parsefiles function to read the the template
	// file into a template set. In case of error, log the error
	// message and send a generic response to the client
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Execute the template set and write the response
	ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
