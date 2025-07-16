package api

import (
	"encoding/json"
	"net/http"
	"routing/algorithm/src/dependency"
	"routing/algorithm/src/internal"
	"routing/algorithm/utils"
)

func FindLowestCost(w http.ResponseWriter, r *http.Request) {
	if err := validateRoute(w, r, http.MethodGet); err != nil {
		return
	}

	orderIdsInput := r.URL.Query().Get("order_ids")
	warehouseLongitude, err := utils.StringToFloat64(r.URL.Query().Get("warehouse_longitude"))
	if err != nil {
		http.Error(w, "Invalid warehouse longitude", http.StatusBadRequest)
		return
	}
	warehouseLatitude, err := utils.StringToFloat64(r.URL.Query().Get("warehouse_latitude"))
	if err != nil {
		http.Error(w, "Invalid warehouse latitude", http.StatusBadRequest)
		return
	}

	route, err := internal.FindMinimumCost(&dependency.Route{
		OrderIds:               orderIdsInput,
		WarehouseAddrLongitude: warehouseLongitude,
		WarehouseAddrLatitude:  warehouseLatitude,
	})
	if err != nil {
		http.Error(w, "Failed to find minimum cost: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(route)
	if err != nil {
		http.Error(w, "Failed to encode route", http.StatusInternalServerError)
	}

	http.Error(w, "", http.StatusOK)
}
