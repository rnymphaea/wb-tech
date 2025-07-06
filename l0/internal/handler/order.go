package handler

import (
	"context"
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
		http.Redirect(w, r, "/?error=Order+UID+is+required", http.StatusSeeOther)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	order, err := h.storage.GetOrderByUID(ctx, uid)
	if err != nil {
		log.Printf("Error getting order: %v", err)
		http.Redirect(w, r, "/?error=Internal+server+error", http.StatusSeeOther)
		return
	}
	if order == nil {
		http.Redirect(w, r, "/?error=Order+not+found", http.StatusSeeOther)
		return
	}
	data := map[string]interface{}{
    "Title": "Order Details",
		"Order": order,
  }

  err = templates.ExecuteTemplate(w, "order.html", data)
  if err != nil {
    log.Printf("Error rendering order page: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
  }

	w.Header().Set("Content-Type", "application/json")
}
