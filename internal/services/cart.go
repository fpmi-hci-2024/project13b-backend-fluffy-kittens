package services

import (
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) GetCartByUserID(customerID string) (models.Cart, error) {
	query := `
        SELECT product_id
        FROM cart_products
        WHERE customer_id = $1
    `
	rows, err := pdb.db.Query(query, customerID)
	if err != nil {
		return models.Cart{}, fmt.Errorf("failed to get cart: %w", err)
	}
	defer rows.Close()

	var productIDs []string
	for rows.Next() {
		var productID string
		if err := rows.Scan(&productID); err != nil {
			return models.Cart{}, fmt.Errorf("failed to scan product ID: %w", err)
		}
		productIDs = append(productIDs, productID)
	}

	if err := rows.Err(); err != nil {
		return models.Cart{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	return models.Cart{
		CustomerID: customerID,
		ProductIDs: productIDs,
	}, nil
}

func (pdb *PostgresDatabase) AddProductToCart(customerID string, productID string) error {
	query := `
        INSERT INTO cart_products (customer_id, product_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, customerID, productID)
	if err != nil {
		return fmt.Errorf("failed to add product to cart: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) RemoveProductFromCart(customerID string, productID string) error {
	query := `
        DELETE FROM cart_products
        WHERE customer_id = $1 AND product_id = $2
    `
	result, err := pdb.db.Exec(query, customerID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove product from cart: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product %s not found in cart for customer %s", productID, customerID)
	}
	return nil
}
