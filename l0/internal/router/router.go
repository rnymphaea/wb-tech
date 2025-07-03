package router

import (
	"net/http"

	"wb-tech-l0/internal/handler"
	"wb-tech-l0/internal/database/postgres"
)

func NewRouter(storage *postgres.Storage) *http.ServeMux {
	mux := http.NewServeMux()
	orderHandler := handler.NewOrderHandler(storage)

	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/order/", orderHandler.GetOrder)

	return mux
}
