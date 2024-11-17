package handler

import (
	"billing-app/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.FetchUsersFromDatabase()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users data", http.StatusInternalServerError)
	}

	// TODO: ankit.pinge add logging here
}

func GetProductsByShopID(w http.ResponseWriter, r *http.Request) {
	shopIDStr := r.URL.Query().Get("shop_id")
	if shopIDStr == "" {
		http.Error(w, "Missing shop_id query parameter", http.StatusBadRequest)
		return
	}

	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		http.Error(w, "Invalid shop_id parameter", http.StatusBadRequest)
		return
	}

	products, err := database.FetchProductsByShopID(shopID)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products data", http.StatusInternalServerError)
	}
}
