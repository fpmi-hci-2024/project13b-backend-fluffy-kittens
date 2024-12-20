package models

type Cart struct {
	ID         string   `json:"id"`
	CustomerID string   `json:"customerId"`
	ProductIDs []string `json:"productIds"`
}
