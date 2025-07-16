package api

import (
	"net/http"
)

func OrderSubmit(w http.ResponseWriter, r *http.Request) {
	// Call the internal function to handle order submission
	//internal.OrderSubmit(w, r)
}

func FindLowestCost(w http.ResponseWriter, r *http.Request) {
	// Call the internal function to find the lowest cost
	//internal.FindLowestCost(w, r)
}

func RetrieveRouteInfo(w http.ResponseWriter, r *http.Request) {
	// Call the internal function to retrieve route information
	//internal.RetrieveRouteInfo(w, r)
}
