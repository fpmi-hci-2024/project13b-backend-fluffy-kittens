package models

type Favorites struct {
	CustomerID string   `json:"customerId"`
	ProductIDs []string `json:"productIds"`
}
