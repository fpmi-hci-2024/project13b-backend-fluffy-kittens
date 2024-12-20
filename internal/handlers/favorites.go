package handlers

import (
	"encoding/json"
	"net/http"

	"fluffy-shop-api/internal/models"
	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

type FavoritesHandler struct {
	db services.FavoritesDB
}

func NewFavoritesHandler(db services.FavoritesDB) *FavoritesHandler {
	return &FavoritesHandler{db: db}
}

func (h *FavoritesHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get all favorites
}

func (h *FavoritesHandler) CreateFavorites(w http.ResponseWriter, r *http.Request) {
	var favorites models.Favorites
	if err := json.NewDecoder(r.Body).Decode(&favorites); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.CreateFavorites(favorites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FavoritesHandler) GetFavoritesByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	favoriteID := params["favoriteId"]
	favorites, err := h.db.GetFavoritesByID(favoriteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(favorites)
}

func (h *FavoritesHandler) UpdateFavorites(w http.ResponseWriter, r *http.Request) {
	var favorites models.Favorites
	if err := json.NewDecoder(r.Body).Decode(&favorites); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.UpdateFavorites(favorites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FavoritesHandler) DeleteFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	favoriteID := params["favoriteId"]
	if err := h.db.DeleteFavorites(favoriteID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *FavoritesHandler) AddProductToFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	favoriteID := params["favoriteId"]
	productID := params["productId"]
	if err := h.db.AddProductToFavorites(favoriteID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FavoritesHandler) RemoveProductFromFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	favoriteID := params["favoriteId"]
	productID := params["productId"]
	if err := h.db.RemoveProductFromFavorites(favoriteID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
