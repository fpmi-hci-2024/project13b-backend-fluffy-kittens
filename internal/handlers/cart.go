package handlers

import (
	"encoding/json"
	"net/http"

	"fluffy-shop-api/internal/models"
	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

type CartHandler struct {
	db services.CartDB
}

func NewCartHandler(db services.CartDB) *CartHandler {
	return &CartHandler{db: db}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.CreateCart(cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	cart, err := h.db.GetCartByID(cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.UpdateCart(cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	if err := h.db.DeleteCart(cartID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	productID := params["productId"]
	if err := h.db.AddProductToCart(cartID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) RemoveProductFromCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["cartId"]
	productID := params["productId"]
	if err := h.db.RemoveProductFromCart(cartID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
