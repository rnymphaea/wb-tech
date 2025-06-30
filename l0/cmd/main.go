package main

import (
	"log"

	"wb-tech-l0/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded")
	databaseURL := cfg.GetDatabaseURL()
	log.Printf("DB URL: %s", databaseURL)
}
