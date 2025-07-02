package router

import (
	"net/http"

	"wb-tech-l0/internal/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomePage)

	return mux
}
