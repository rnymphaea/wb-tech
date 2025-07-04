package main

import (
	"log"
	"net/http"
	"context"
	"time"

	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/router"
	"wb-tech-l0/internal/database/postgres"
	"wb-tech-l0/internal/database/cache/redis"
	"wb-tech-l0/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded")
	databaseURL := cfg.GetDatabaseURL()
	
	pgrepo, err := postgres.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer pgrepo.Close()

	cache := redis.NewRedisCache(cfg.Redis.Addr, cfg.GetRedisTTL())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := cache.Ping(ctx); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
	
	storage := database.New(cache, pgrepo)

	mux := router.NewRouter(storage)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
