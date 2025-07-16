package main

import (
	"log"
	"net/http"
)

func main() {
	routes()
	startServer()
}

func startServer() {
	port := ":8080"
	log.Printf("Server starting on port %s\n", port)

	err := http.ListenAndServe(port, nil) // The 'nil' argument means use the default ServeMux
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
