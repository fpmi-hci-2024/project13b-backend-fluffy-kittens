package handlers

import (
	"encoding/json"
	"net/http"

	"fluffy-shop-api/internal/models"
	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
    db services.Database
}

func NewOrderHandler(db services.Database) *OrderHandler {
    return &OrderHandler{db: db}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
    orders, err := h.db.GetOrders()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err = h.db.CreateOrder(order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderId"]
    order, err := h.db.GetOrder(orderID)
    if err != nil {
        http.Error(w, "Order not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderId"]
    var updatedOrder models.Order
    err := json.NewDecoder(r.Body).Decode(&updatedOrder)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updatedOrder.ID = orderID
    err = h.db.UpdateOrder(updatedOrder)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["orderId"]
    err := h.db.DeleteOrder(orderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
