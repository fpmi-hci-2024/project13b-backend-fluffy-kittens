package services

import (
	"database/sql"
	"fluffy-shop-api/internal/models"
	"fmt"
)

func (pdb *PostgresDatabase) CreateCustomer(customer models.Customer) error {
	query := `
        INSERT INTO customers (id, name, email, phone)
        VALUES ($1, $2, $3, $4)
    `
	_, err := pdb.db.Exec(query, customer.ID, customer.Name, customer.Email, customer.Phone)
	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}

func (pdb *PostgresDatabase) GetCustomerByID(id string) (models.Customer, error) {
	query := `
        SELECT id, name, email, phone
        FROM customers
        WHERE id = $1
    `
	var customer models.Customer
	err := pdb.db.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Customer{}, fmt.Errorf("customer with ID %s not found", id)
		}
		return models.Customer{}, fmt.Errorf("failed to get customer: %w", err)
	}
	return customer, nil
}

func (pdb *PostgresDatabase) UpdateCustomer(customer models.Customer) error {
	query := `
        UPDATE customers
        SET name = $1, email = $2, phone = $3
        WHERE id = $4
    `
	result, err := pdb.db.Exec(query, customer.Name, customer.Email, customer.Phone, customer.ID)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("customer with ID %s not found", customer.ID)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteCustomer(id string) error {
	query := `
        DELETE FROM customers
        WHERE id = $1
    `
	result, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
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
