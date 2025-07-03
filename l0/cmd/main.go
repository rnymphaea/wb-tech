package main

import (
	"log"
	"net/http"

	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/router"
	"wb-tech-l0/internal/database/postgres"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded")
	databaseURL := cfg.GetDatabaseURL()
	log.Printf("DB URL: %s", databaseURL)
	
	storage, _ := postgres.New(databaseURL)
	mux := router.NewRouter(storage)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
