package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"wb-tech-l0/internal/database"
)

type OrderHandler struct {
	storage *database.Storage
}

func NewOrderHandler(storage *database.Storage) *OrderHandler {
	return &OrderHandler{storage: storage}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	uid := strings.TrimPrefix(path, "/order/")
	if uid == "" {
		http.Error(w, "Order UID is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	order, err := h.storage.GetOrderByUID(ctx, uid)
	if err != nil {
		log.Printf("Error getting order: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
