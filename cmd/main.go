package main

import (
	"log"
	"net/http"

	"payment-options/internal/handler"
	"payment-options/internal/repository"
)

func main() {
	// Initialize database manager
	dbManager := repository.NewDBManager()
	err := dbManager.ConnectDatabases()
	if err != nil {
		log.Println("Warning: Database connection error:", err)
	}
	defer dbManager.Close()

	// Initialize handlers
	paymentHandler := handler.NewPaymentHandler()

	// Set up routes
	http.HandleFunc("/payment/options", paymentHandler.GetPaymentOptions)

	// Start server
	port := ":8081"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
