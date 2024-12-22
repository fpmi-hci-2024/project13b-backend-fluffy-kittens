package services

import (
	"database/sql"
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) CreateOrder(order models.Order) error {
	query := `
        INSERT INTO orders (id, customer_id, total, status)
        VALUES ($1, $2, $3, $4)
    `
	_, err := pdb.db.Exec(query, order.ID, order.CustomerID, order.Total, order.Status)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) GetOrderByID(id string) (models.Order, error) {
	query := `
        SELECT id, customer_id, total, status
        FROM orders
        WHERE id = $1
    `
	var order models.Order
	err := pdb.db.QueryRow(query, id).Scan(&order.ID, &order.CustomerID, &order.Total, &order.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, fmt.Errorf("order with ID %s not found", id)
		}
		return models.Order{}, fmt.Errorf("failed to get order: %w", err)
	}
	return order, nil
}

func (pdb *PostgresDatabase) UpdateOrder(order models.Order) error {
	query := `
        UPDATE orders
        SET total = $1, status = $2
        WHERE id = $3
    `
	result, err := pdb.db.Exec(query, order.Total, order.Status, order.ID)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("order with ID %s not found", order.ID)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteOrder(id string) error {
	query := `
        DELETE FROM orders
        WHERE id = $1
    `
	result, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("order with ID %s not found", id)
	}
	return nil
}

func (pdb *PostgresDatabase) AddProductToOrder(orderID string, productID string) error {
	query := `
        INSERT INTO order_products (order_id, product_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, orderID, productID)
	if err != nil {
		return fmt.Errorf("failed to add product to order: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) RemoveProductFromOrder(orderID string, productID string) error {
	query := `
        DELETE FROM order_products
        WHERE order_id = $1 AND product_id = $2
    `
	result, err := pdb.db.Exec(query, orderID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove product from order: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product %s not found in order %s", productID, orderID)
	}
	return nil
}
