package services

import (
	"database/sql"
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) CreateCart(cart models.Cart) error {
	query := `
        INSERT INTO cart (id, customer_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, cart.ID, cart.CustomerID)
	if err != nil {
		return fmt.Errorf("failed to create cart: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) GetCartByID(id string) (models.Cart, error) {
	query := `
        SELECT id, customer_id
        FROM cart
        WHERE id = $1
    `
	var cart models.Cart
	err := pdb.db.QueryRow(query, id).Scan(&cart.ID, &cart.CustomerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Cart{}, fmt.Errorf("cart with ID %s not found", id)
		}
		return models.Cart{}, fmt.Errorf("failed to get cart: %w", err)
	}
	return cart, nil
}

func (pdb *PostgresDatabase) UpdateCart(cart models.Cart) error {
	query := `
        UPDATE cart
        SET customer_id = $1
        WHERE id = $2
    `
	result, err := pdb.db.Exec(query, cart.CustomerID, cart.ID)
	if err != nil {
		return fmt.Errorf("failed to update cart: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cart with ID %s not found", cart.ID)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteCart(id string) error {
	query := `
        DELETE FROM cart
        WHERE id = $1
    `
	result, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete cart: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cart with ID %s not found", id)
	}
	return nil
}

func (pdb *PostgresDatabase) AddProductToCart(cartID string, productID string) error {
	query := `
        INSERT INTO cart_products (cart_id, product_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, cartID, productID)
	if err != nil {
		return fmt.Errorf("failed to add product to cart: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) RemoveProductFromCart(cartID string, productID string) error {
	query := `
        DELETE FROM cart_products
        WHERE cart_id = $1 AND product_id = $2
    `
	result, err := pdb.db.Exec(query, cartID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove product from cart: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product %s not found in cart %s", productID, cartID)
	}
	return nil
}
