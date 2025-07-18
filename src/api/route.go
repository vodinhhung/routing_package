package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"routing/algorithm/src/dependency"
	"routing/algorithm/src/internal"
	"strconv"
)

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	if err := validate(w, r, &Args{methodType: http.MethodPost}, validateHttpMethod); err != nil {
		return
	}

	var newRoute *dependency.Route
	if err := json.NewDecoder(r.Body).Decode(newRoute); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := internal.CreateNewRoute(newRoute); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create route: %v", err), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusCreated)
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	if err := validate(w, r, &Args{methodType: http.MethodGet}, validateHttpMethod); err != nil {
		return
	}

	routeIdInput := r.URL.Query().Get("route_id")
	routeId, err := strconv.ParseUint(routeIdInput, 10, 64)
	if err != nil || routeId <= 0 {
		http.Error(w, "Invalid route ID", http.StatusBadRequest)
		return
	}

	route, err := internal.GetRoute(routeId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete route: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(route)
	if err != nil {
		http.Error(w, "Failed to encode route", http.StatusInternalServerError)
	}

	http.Error(w, "", http.StatusOK)
}
