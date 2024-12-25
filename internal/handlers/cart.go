package handlers

import (
	"encoding/json"
	"net/http"

	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

type CartHandler struct {
	db services.CartDB
}

func NewCartHandler(db services.CartDB) *CartHandler {
	return &CartHandler{db: db}
}

// GetCartByUserID retrieves the cart for a specific user.
func (h *CartHandler) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	cart, err := h.db.GetCartByUserID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

// AddProductToCart adds a product to the user's cart.
func (h *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	productID := params["productId"]
	if err := h.db.AddProductToCart(customerID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// RemoveProductFromCart removes a product from the user's cart.
func (h *CartHandler) RemoveProductFromCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	productID := params["productId"]
	if err := h.db.RemoveProductFromCart(customerID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
