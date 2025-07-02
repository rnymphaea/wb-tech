package main

import (
	"log"
	"net/http"

	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded")
	databaseURL := cfg.GetDatabaseURL()
	log.Printf("DB URL: %s", databaseURL)

	mux := router.NewRouter()
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
