package services

import (
	"fluffy-shop-api/internal/models"
)

type Database interface {
	GetProducts() (map[string]models.Product, error)
	CreateProduct(product models.Product) error
	GetProduct(id string) (models.Product, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(id string) error

	GetOrders() (map[string]models.Order, error)
	CreateOrder(order models.Order) error
	GetOrder(id string) (models.Order, error)
	UpdateOrder(order models.Order) error
	DeleteOrder(id string) error

	GetCustomers() (map[string]models.Customer, error)
	CreateCustomer(customer models.Customer) error
	GetCustomer(id string) (models.Customer, error)
	UpdateCustomer(customer models.Customer) error
	DeleteCustomer(id string) error
}
