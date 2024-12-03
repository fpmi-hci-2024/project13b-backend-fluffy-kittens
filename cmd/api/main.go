package main

import (
	"log"
	"net/http"

	"fluffy-shop-api/internal/handlers"
	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database (for production, use the actual database connection)
	db := services.NewMockDatabase()

	// Initialize handlers with the database instance
	productHandler := handlers.NewProductHandler(db)
	orderHandler := handlers.NewOrderHandler(db)
	customerHandler := handlers.NewCustomerHandler(db)

	r := mux.NewRouter()

	// Product routes
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{productId}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products/{productId}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{productId}", productHandler.DeleteProduct).Methods("DELETE")

	// Order routes
	r.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{orderId}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{orderId}", orderHandler.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{orderId}", orderHandler.DeleteOrder).Methods("DELETE")

	// Customer routes
	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{customerId}", customerHandler.GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{customerId}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{customerId}", customerHandler.DeleteCustomer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
