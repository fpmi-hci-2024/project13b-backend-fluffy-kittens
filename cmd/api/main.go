package main

import (
	"log"
	"net/http"

	"fluffy-shop-api/internal/handlers"
	"fluffy-shop-api/internal/services"
	"fluffy-shop-api/internal/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database (for production, use the actual database connection)
	db, err := services.NewPostgresDatabase("postgresql://fluffy_kitten:htUtmkIjKHoQuwO6cDcbvsNsQZzXuYb7@dpg-ct7mmg56l47c73crse90-a/fluffy_db")

	if err != nil {
		log.Fatal("Could not initialize db connection")
	}

	// Initialize handlers with the database instance
	productHandler := handlers.NewProductHandler(db)
	orderHandler := handlers.NewOrderHandler(db)
	customerHandler := handlers.NewCustomerHandler(db)
	cartHandler := handlers.NewCartHandler(db)           // Initialize CartHandler
	favoritesHandler := handlers.NewFavoritesHandler(db) // Initialize FavoritesHandler

	r := mux.NewRouter()
	r.Use(utils.EnableCORS)

	// Product routes
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{productId}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products/{productId}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{productId}", productHandler.DeleteProduct).Methods("DELETE")

	// Order routes
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{orderId}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{orderId}", orderHandler.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{orderId}", orderHandler.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{orderId}/products/{productId}", orderHandler.AddProductToOrder).Methods("POST")
	r.HandleFunc("/orders/{orderId}/products/{productId}", orderHandler.RemoveProductFromOrder).Methods("DELETE")

	// Customer routes
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{customerId}", customerHandler.GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{customerId}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{customerId}", customerHandler.DeleteCustomer).Methods("DELETE")

	// Cart routes
	r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
	r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/carts/{cartId}", cartHandler.UpdateCart).Methods("PUT")
	r.HandleFunc("/carts/{cartId}", cartHandler.DeleteCart).Methods("DELETE")
	r.HandleFunc("/carts/{cartId}/products/{productId}", cartHandler.AddProductToCart).Methods("POST")
	r.HandleFunc("/carts/{cartId}/products/{productId}", cartHandler.RemoveProductFromCart).Methods("DELETE")

	// Favorites routes
	r.HandleFunc("/favorites", favoritesHandler.CreateFavorites).Methods("POST")
	r.HandleFunc("/favorites/{favoriteId}", favoritesHandler.GetFavoritesByID).Methods("GET")
	r.HandleFunc("/favorites/{favoriteId}", favoritesHandler.UpdateFavorites).Methods("PUT")
	r.HandleFunc("/favorites/{favoriteId}", favoritesHandler.DeleteFavorites).Methods("DELETE")
	r.HandleFunc("/favorites/{favoriteId}/products/{productId}", favoritesHandler.AddProductToFavorites).Methods("POST")
	r.HandleFunc("/favorites/{favoriteId}/products/{productId}", favoritesHandler.RemoveProductFromFavorites).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
