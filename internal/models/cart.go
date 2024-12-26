package models

type Cart struct {
	CustomerID string   `json:"customerId"`
	ProductIDs []string `json:"productIds"`
}
