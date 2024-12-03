package models

type Order struct {
	ID         string   `json:"id"`
	CustomerID string   `json:"customerId"`
	ProductIDs []string `json:"productIds"`
	Total      float64  `json:"total"`
	Status     string   `json:"status"`
}
