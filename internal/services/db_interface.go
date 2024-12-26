package services

import (
	"fluffy-shop-api/internal/models"
)

// CustomerDB defines database operations for Customer model
type CustomerDB interface {
	CreateCustomer(customer models.Customer) error
	GetCustomerByID(id string) (models.Customer, error)
	UpdateCustomer(customer models.Customer) error
	DeleteCustomer(id string) error
}

// ProductDB defines database operations for Product model
type ProductDB interface {
	CreateProduct(product models.Product) error
	GetProductByID(id string) (models.Product, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(id string) error
	GetAllProducts() (map[string]models.Product, error)
}

// OrderDB defines database operations for Order model
type OrderDB interface {
	CreateOrder(order models.Order) error
	GetOrderByID(id string) (models.Order, error)
	UpdateOrder(order models.Order) error
	DeleteOrder(id string) error
	AddProductToOrder(orderID string, productID string) error
	RemoveProductFromOrder(orderID string, productID string) error
}

// CartDB defines database operations for Cart model
type CartDB interface {
	GetCartByUserID(customerID string) (models.Cart, error)
	AddProductToCart(customerID string, productID string) error
	RemoveProductFromCart(customerID string, productID string) error
}

// FavoritesDB defines database operations for Favorites model
type FavoritesDB interface {
	GetFavoritesByUserID(id string) (models.Favorites, error)
	AddProductToFavorites(favoriteID string, productID string) error
	RemoveProductFromFavorites(favoriteID string, productID string) error
}
