package main

import (
	"billing-app/database"
	"billing-app/handler"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler.GetUsers)
	mux.HandleFunc("/products", handler.GetProductsByShopID)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
