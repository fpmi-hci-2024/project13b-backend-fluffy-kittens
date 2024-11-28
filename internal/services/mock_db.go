package services

import (
	"fluffy-shop-api/internal/models"
)

type MockDatabase struct {
	Products  map[string]models.Product
	Orders    map[string]models.Order
	Customers map[string]models.Customer
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		Products:  make(map[string]models.Product),
		Orders:    make(map[string]models.Order),
		Customers: make(map[string]models.Customer),
	}
}

func (m *MockDatabase) GetProducts() (map[string]models.Product, error) {
	return m.Products, nil
}

func (m *MockDatabase) CreateProduct(product models.Product) error {
	m.Products[product.ID] = product
	return nil
}

func (m *MockDatabase) GetProduct(id string) (models.Product, error) {
	product, exists := m.Products[id]
	if !exists {
		return models.Product{}, nil
	}
	return product, nil
}

func (m *MockDatabase) UpdateProduct(product models.Product) error {
	m.Products[product.ID] = product
	return nil
}

func (m *MockDatabase) DeleteProduct(id string) error {
	delete(m.Products, id)
	return nil
}

func (m *MockDatabase) GetOrders() (map[string]models.Order, error) {
	return m.Orders, nil
}

func (m *MockDatabase) CreateOrder(order models.Order) error {
	m.Orders[order.ID] = order
	return nil
}

func (m *MockDatabase) GetOrder(id string) (models.Order, error) {
	order, exists := m.Orders[id]
	if !exists {
		return models.Order{}, nil
	}
	return order, nil
}

func (m *MockDatabase) UpdateOrder(order models.Order) error {
	m.Orders[order.ID] = order
	return nil
}

func (m *MockDatabase) DeleteOrder(id string) error {
	delete(m.Orders, id)
	return nil
}

func (m *MockDatabase) GetCustomers() (map[string]models.Customer, error) {
	return m.Customers, nil
}

func (m *MockDatabase) CreateCustomer(customer models.Customer) error {
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockDatabase) GetCustomer(id string) (models.Customer, error) {
	customer, exists := m.Customers[id]
	if !exists {
		return models.Customer{}, nil
	}
	return customer, nil
}

func (m *MockDatabase) UpdateCustomer(customer models.Customer) error {
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockDatabase) DeleteCustomer(id string) error {
	delete(m.Customers, id)
	return nil
}
