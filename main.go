package main

import (
	"log"
	"net/http"
)

func main() {
	// Handle HTTP requests to the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/chris.html")
	})

	// Log that the server is starting
	log.Println("Server starting on http://localhost:8080/")

	// Start the server on port 8080 and log any errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
