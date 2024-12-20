package services

import (
	"database/sql"
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) CreateProduct(product models.Product) error {
	query := `
        INSERT INTO products (id, name, description, price, stock)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := pdb.db.Exec(query, product.ID, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) GetProductByID(id string) (models.Product, error) {
	query := `
        SELECT id, name, description, price, stock
        FROM products
        WHERE id = $1
    `
	var product models.Product
	err := pdb.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("product with ID %s not found", id)
		}
		return models.Product{}, fmt.Errorf("failed to get product: %w", err)
	}
	return product, nil
}

func (pdb *PostgresDatabase) UpdateProduct(product models.Product) error {
	query := `
        UPDATE products
        SET name = $1, description = $2, price = $3, stock = $4
        WHERE id = $5
    `
	result, err := pdb.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %s not found", product.ID)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteProduct(id string) error {
	query := `
        DELETE FROM products
        WHERE id = $1
    `
	result, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %s not found", id)
	}
	return nil
}

func (pdb *PostgresDatabase) GetAllProducts() ([]models.Product, error) {
	query := `
        SELECT id, name, description, price, stock
        FROM products
    `
	rows, err := pdb.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return products, nil
}
