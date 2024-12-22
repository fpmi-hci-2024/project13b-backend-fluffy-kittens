package services

import (
	"database/sql"
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) CreateFavorites(favorites models.Favorites) error {
	query := `
        INSERT INTO favorites (id, customer_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, favorites.ID, favorites.CustomerID)
	if err != nil {
		return fmt.Errorf("failed to create favorites: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) GetFavoritesByID(id string) (models.Favorites, error) {
	query := `
        SELECT id, customer_id
        FROM favorites
        WHERE id = $1
    `
	var favorites models.Favorites
	err := pdb.db.QueryRow(query, id).Scan(&favorites.ID, &favorites.CustomerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Favorites{}, fmt.Errorf("favorites with ID %s not found", id)
		}
		return models.Favorites{}, fmt.Errorf("failed to get favorites: %w", err)
	}
	return favorites, nil
}

func (pdb *PostgresDatabase) UpdateFavorites(favorites models.Favorites) error {
	query := `
        UPDATE favorites
        SET customer_id = $1
        WHERE id = $2
    `
	result, err := pdb.db.Exec(query, favorites.CustomerID, favorites.ID)
	if err != nil {
		return fmt.Errorf("failed to update favorites: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("favorites with ID %s not found", favorites.ID)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteFavorites(id string) error {
	query := `
        DELETE FROM favorites
        WHERE id = $1
    `
	result, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete favorites: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("customer with ID %s not found", id)
	}
	return nil
}

func (pdb *PostgresDatabase) AddProductToFavorites(favoriteID string, productID string) error {
	query := `
        INSERT INTO favorites_products (favorite_id, product_id)
        VALUES ($1, $2)
    `
	_, err := pdb.db.Exec(query, favoriteID, productID)
	if err != nil {
		return fmt.Errorf("failed to add product to favorites: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) RemoveProductFromFavorites(favoriteID string, productID string) error {
	query := `
        DELETE FROM favorites_products
        WHERE favorite_id = $1 AND product_id = $2
    `
	result, err := pdb.db.Exec(query, favoriteID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove product from favorites: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %s not found in favorites with ID %s", productID, favoriteID)
	}
	return nil
}
