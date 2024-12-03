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
	db := &MockDatabase{
		Products:  make(map[string]models.Product),
		Orders:    make(map[string]models.Order),
		Customers: make(map[string]models.Customer),
	}

	// Add sample products
	db.Products = map[string]models.Product{
		"p1": {
			ID:          "p1",
			Name:        "Persian Kitten",
			Description: "A beautiful white Persian kitten with blue eyes",
			Price:       999.99,
			Stock:       5,
		},
		"p2": {
			ID:          "p2",
			Name:        "Maine Coon Kitten",
			Description: "Large and friendly Maine Coon kitten",
			Price:       1299.99,
			Stock:       3,
		},
		"p3": {
			ID:          "p3",
			Name:        "Siamese Kitten",
			Description: "Elegant Siamese kitten with striking blue eyes",
			Price:       899.99,
			Stock:       7,
		},
	}

	// Add sample customers
	db.Customers = map[string]models.Customer{
		"c1": {
			ID:    "c1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "555-0123",
		},
		"c2": {
			ID:    "c2",
			Name:  "Jane Smith",
			Email: "jane.smith@example.com",
			Phone: "555-0124",
		},
		"c3": {
			ID:    "c3",
			Name:  "Bob Johnson",
			Email: "bob.johnson@example.com",
			Phone: "555-0125",
		},
	}

	// Add sample orders
	db.Orders = map[string]models.Order{
		"o1": {
			ID:         "o1",
			CustomerID: "c1",
			ProductIDs: []string{"p1", "p2"},
			Total:      2299.98,
			Status:     "pending",
		},
		"o2": {
			ID:         "o2",
			CustomerID: "c2",
			ProductIDs: []string{"p3"},
			Total:      899.99,
			Status:     "completed",
		},
		"o3": {
			ID:         "o3",
			CustomerID: "c3",
			ProductIDs: []string{"p1", "p2", "p3"},
			Total:      3199.97,
			Status:     "processing",
		},
	}

	return db
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
