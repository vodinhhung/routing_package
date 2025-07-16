package main

import (
	"net/http"
	"routing/algorithm/src/api"
)

func routes() {
	// Order endpoints
	http.HandleFunc("/orders/new", api.OrderSubmit)
	http.HandleFunc("/orders/detail", api.OrderSubmit)

	// Cost endpoints
	http.HandleFunc("/cost/lowest", api.FindLowestCost)

	// Route endpoints
	http.HandleFunc("/routes/new", api.RetrieveRouteInfo)
	http.HandleFunc("/routes/detail", api.RetrieveRouteInfo)
}
