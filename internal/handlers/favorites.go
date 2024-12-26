package handlers

import (
	"encoding/json"
	"net/http"

	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

type FavoritesHandler struct {
	db services.FavoritesDB
}

func NewFavoritesHandler(db services.FavoritesDB) *FavoritesHandler {
	return &FavoritesHandler{db: db}
}

// GetFavoritesByUserID retrieves the favorites for a specific user.
func (h *FavoritesHandler) GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	favorites, err := h.db.GetFavoritesByUserID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(favorites)
}

// AddProductToFavorites adds a product to the user's favorites.
func (h *FavoritesHandler) AddProductToFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	productID := params["productId"]
	if err := h.db.AddProductToFavorites(customerID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// RemoveProductFromFavorites removes a product from the user's favorites.
func (h *FavoritesHandler) RemoveProductFromFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["customerId"]
	productID := params["productId"]
	if err := h.db.RemoveProductFromFavorites(customerID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
