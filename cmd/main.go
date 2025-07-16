package main

import (
	"log"
	"net/http"
	"routing/algorithm/src/api"
	"routing/algorithm/src/dependency"
)

func main() {
	initDependencies()
	initEndpoints()
	startServer()
}

func initDependencies() {
	err := dependency.InitOrderDB()
	if err != nil {
		log.Fatalf("Failed to connect to order database: %v", err)
	}

	err = dependency.InitRouteDB()
	if err != nil {
		log.Fatalf("Failed to connect to route database: %v", err)
	}
}

func initEndpoints() {
	// Order endpoints
	http.HandleFunc("/orders/new", api.CreateOrder)
	http.HandleFunc("/orders/detail", api.GetOrder)

	// Cost endpoints
	http.HandleFunc("/cost/lowest", api.FindLowestCost)

	// Route endpoints
	http.HandleFunc("/routes/new", api.CreateRoute)
	http.HandleFunc("/routes/detail", api.GetRoute)
}

func startServer() {
	port := ":8080"
	log.Printf("Server starting on port %s\n", port)

	err := http.ListenAndServe(port, nil) // The 'nil' argument means use the default ServeMux
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
