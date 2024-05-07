package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/* Item represents an item in the webpage */
type ItemMenu struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var inventory []ItemMenu

/* main is the entry point for the application */
func main() {
	// Add some sample data to the inventory
	inventory = append(inventory, ItemMenu{ID: "1", Name: "Intro"})
	inventory = append(inventory, ItemMenu{ID: "2", Name: "Technologies"})
	inventory = append(inventory, ItemMenu{ID: "3", Name: "Projects"})

	// Handle HTTP requests to the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/chris.html")
	})

	// Handle HTTP requests to /api/items
	http.HandleFunc("/api/items", getItems)

	// Log that the server is starting
	log.Println("Server starting on http://localhost:8080/")

	// Start the server on port 8080 and log any errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/* enableCors sets the CORS headers for the response */
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
}

/* getItems handles HTTP requests to the /api/items path */
func getItems(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	enableCors(&w)

	// Set the content type of the response
	w.Header().Set("Content-Type", "application/json")

	// Encode inventory slice to JSON and write it as the response
	err := json.NewEncoder(w).Encode(inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
