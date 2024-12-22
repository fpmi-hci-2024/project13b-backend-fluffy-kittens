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
	CreateCart(cart models.Cart) error
	GetCartByID(id string) (models.Cart, error)
	UpdateCart(cart models.Cart) error
	DeleteCart(id string) error
	AddProductToCart(cartID string, productID string) error
	RemoveProductFromCart(cartID string, productID string) error
}

// FavoritesDB defines database operations for Favorites model
type FavoritesDB interface {
	CreateFavorites(favorites models.Favorites) error
	GetFavoritesByID(id string) (models.Favorites, error)
	UpdateFavorites(favorites models.Favorites) error
	DeleteFavorites(id string) error
	AddProductToFavorites(favoriteID string, productID string) error
	RemoveProductFromFavorites(favoriteID string, productID string) error
}
