package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"routing/algorithm/src/dependency"
	"routing/algorithm/src/internal"
	"strconv"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if err := validate(w, r, &Args{
		methodType: http.MethodPost,
	}, validateHttpMethod); err != nil {
		return
	}

	var newOrder *dependency.Order
	if err := json.NewDecoder(r.Body).Decode(newOrder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validate(w, r, &Args{
		methodType: http.MethodPost,
		latitude:   newOrder.DeliveryLatitude,
		longitude:  newOrder.DeliveryLongitude,
	}, validateHttpMethod, validateLongitudeAndLatitude); err != nil {
		return
	}

	if err := internal.CreateNewOrder(newOrder); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %v", err), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusCreated)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if err := validate(w, r, &Args{methodType: http.MethodGet}, validateHttpMethod); err != nil {
		return
	}

	orderIdInput := r.URL.Query().Get("order_id")
	orderId, err := strconv.ParseUint(orderIdInput, 10, 64)
	if err != nil || orderId <= 0 {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := internal.GetOrder(orderId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, "Failed to encode order", http.StatusInternalServerError)
	}

	http.Error(w, "", http.StatusOK)
}
