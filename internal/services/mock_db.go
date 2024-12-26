package services

import (
	"fluffy-shop-api/internal/models"
	"fmt"
)

type MockDatabase struct {
	Products  map[string]models.Product
	Orders    map[string]models.Order
	Customers map[string]models.Customer
	Carts     map[string]models.Cart
	Favorites map[string]models.Favorites
}

func NewMockDatabase() *MockDatabase {
	db := &MockDatabase{
		Products:  make(map[string]models.Product),
		Orders:    make(map[string]models.Order),
		Customers: make(map[string]models.Customer),
		Carts:     make(map[string]models.Cart),
		Favorites: make(map[string]models.Favorites),
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

	// Add sample carts
	db.Carts = map[string]models.Cart{
		"c1": {
			CustomerID: "c1",
			ProductIDs: []string{"p1", "p2"},
		},
		"c2": {
			CustomerID: "c2",
			ProductIDs: []string{"p3"},
		},
	}

	// Add sample favorites
	db.Favorites = map[string]models.Favorites{
		"c1": {
			CustomerID: "c1",
			ProductIDs: []string{"p1", "p3"},
		},
		"c2": {
			CustomerID: "c2",
			ProductIDs: []string{"p2"},
		},
	}

	return db
}

// Implement ProductDB interface
func (m *MockDatabase) CreateProduct(product models.Product) error {
	m.Products[product.ID] = product
	return nil
}

func (m *MockDatabase) GetProductByID(id string) (models.Product, error) {
	product, exists := m.Products[id]
	if !exists {
		return models.Product{}, fmt.Errorf("product with ID %s not found", id)
	}
	return product, nil
}

func (m *MockDatabase) UpdateProduct(product models.Product) error {
	_, exists := m.Products[product.ID]
	if !exists {
		return fmt.Errorf("product with ID %s not found", product.ID)
	}
	m.Products[product.ID] = product
	return nil
}

func (m *MockDatabase) DeleteProduct(id string) error {
	_, exists := m.Products[id]
	if !exists {
		return fmt.Errorf("product with ID %s not found", id)
	}
	delete(m.Products, id)
	return nil
}

func (m *MockDatabase) GetAllProducts() (map[string]models.Product, error) {
	return m.Products, nil
}

// Implement CustomerDB interface
func (m *MockDatabase) CreateCustomer(customer models.Customer) error {
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockDatabase) GetCustomerByID(id string) (models.Customer, error) {
	customer, exists := m.Customers[id]
	if !exists {
		return models.Customer{}, fmt.Errorf("customer with ID %s not found", id)
	}
	return customer, nil
}

func (m *MockDatabase) UpdateCustomer(customer models.Customer) error {
	_, exists := m.Customers[customer.ID]
	if !exists {
		return fmt.Errorf("customer with ID %s not found", customer.ID)
	}
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockDatabase) DeleteCustomer(id string) error {
	_, exists := m.Customers[id]
	if !exists {
		return fmt.Errorf("customer with ID %s not found", id)
	}
	delete(m.Customers, id)
	return nil
}

// Implement OrderDB interface
func (m *MockDatabase) CreateOrder(order models.Order) error {
	m.Orders[order.ID] = order
	return nil
}

func (m *MockDatabase) GetOrderByID(id string) (models.Order, error) {
	order, exists := m.Orders[id]
	if !exists {
		return models.Order{}, fmt.Errorf("order with ID %s not found", id)
	}
	return order, nil
}

func (m *MockDatabase) UpdateOrder(order models.Order) error {
	_, exists := m.Orders[order.ID]
	if !exists {
		return fmt.Errorf("order with ID %s not found", order.ID)
	}
	m.Orders[order.ID] = order
	return nil
}

func (m *MockDatabase) DeleteOrder(id string) error {
	_, exists := m.Orders[id]
	if !exists {
		return fmt.Errorf("order with ID %s not found", id)
	}
	delete(m.Orders, id)
	return nil
}

func (m *MockDatabase) AddProductToOrder(orderID string, productID string) error {
	order, exists := m.Orders[orderID]
	if !exists {
		return fmt.Errorf("order with ID %s not found", orderID)
	}
	order.ProductIDs = append(order.ProductIDs, productID)
	m.Orders[orderID] = order
	return nil
}

func (m *MockDatabase) RemoveProductFromOrder(orderID string, productID string) error {
	order, exists := m.Orders[orderID]
	if !exists {
		return fmt.Errorf("order with ID %s not found", orderID)
	}
	for i, pid := range order.ProductIDs {
		if pid == productID {
			order.ProductIDs = append(order.ProductIDs[:i], order.ProductIDs[i+1:]...)
			m.Orders[orderID] = order
			return nil
		}
	}
	return fmt.Errorf("product with ID %s not found in order %s", productID, orderID)
}

func (m *MockDatabase) GetCartByUserID(id string) (models.Cart, error) {
	cart, exists := m.Carts[id]
	if !exists {
		return models.Cart{}, fmt.Errorf("cart with ID %s not found", id)
	}
	return cart, nil
}

func (m *MockDatabase) AddProductToCart(cartID string, productID string) error {
	cart, exists := m.Carts[cartID]
	if !exists {
		return fmt.Errorf("cart with ID %s not found", cartID)
	}
	cart.ProductIDs = append(cart.ProductIDs, productID)
	m.Carts[cartID] = cart
	return nil
}

func (m *MockDatabase) RemoveProductFromCart(cartID string, productID string) error {
	cart, exists := m.Carts[cartID]
	if !exists {
		return fmt.Errorf("cart with ID %s not found", cartID)
	}
	for i, pid := range cart.ProductIDs {
		if pid == productID {
			cart.ProductIDs = append(cart.ProductIDs[:i], cart.ProductIDs[i+1:]...)
			m.Carts[cartID] = cart
			return nil
		}
	}
	return fmt.Errorf("product with ID %s not found in cart %s", productID, cartID)
}

func (m *MockDatabase) GetFavoritesByUserID(id string) (models.Favorites, error) {
	favorites, exists := m.Favorites[id]
	if !exists {
		return models.Favorites{}, fmt.Errorf("favorites with ID %s not found", id)
	}
	return favorites, nil
}

func (m *MockDatabase) AddProductToFavorites(favoriteID string, productID string) error {
	favorites, exists := m.Favorites[favoriteID]
	if !exists {
		return fmt.Errorf("favorites with ID %s not found", favoriteID)
	}
	favorites.ProductIDs = append(favorites.ProductIDs, productID)
	m.Favorites[favoriteID] = favorites
	return nil
}

func (m *MockDatabase) RemoveProductFromFavorites(favoriteID string, productID string) error {
	favorites, exists := m.Favorites[favoriteID]
	if !exists {
		return fmt.Errorf("favorites with ID %s not found", favoriteID)
	}
	for i, pid := range favorites.ProductIDs {
		if pid == productID {
			favorites.ProductIDs = append(favorites.ProductIDs[:i], favorites.ProductIDs[i+1:]...)
			m.Favorites[favoriteID] = favorites
			return nil
		}
	}
	return fmt.Errorf("product with ID %s not found in favorites %s", productID, favoriteID)
}
