package main

import (
	"log"
	"net/http"

	"github.com/jtreutel/golang-fileserver/internal/handlers"
)

func main() {
	// Set up routes for file handling
	http.HandleFunc("/files", handlers.ListFiles)       // GET
	http.HandleFunc("/files/", handlers.FileOperations) // POST, DELETE

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
