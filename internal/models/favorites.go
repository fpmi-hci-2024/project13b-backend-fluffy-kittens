package models

type Favorites struct {
	ID         string   `json:"id"`
	CustomerID string   `json:"customerId"`
	ProductIDs []string `json:"productIds"`
}
