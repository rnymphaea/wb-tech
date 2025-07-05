package main

import (
	"log"
	"net/http"
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"

	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/router"
	"wb-tech-l0/internal/database/postgres"
	"wb-tech-l0/internal/database/cache/redis"
	"wb-tech-l0/internal/database"
	"wb-tech-l0/internal/kafka"
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
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	log.Println("Connected to Postgres")

	defer pgrepo.Close()

	cache := redis.NewRedisCache(cfg.Redis.Addr, cfg.Redis.TTL)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := cache.Ping(ctx); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
	
	storage := database.New(cache, pgrepo)

	kafkaConsumer := kafka.NewConsumer(
		cfg.Kafka.Brokers,
		cfg.Kafka.Topic,
		cfg.Kafka.GroupID,
		storage,
	)

	consumerCtx, consumerCancel := context.WithCancel(context.Background())
	defer consumerCancel()

	go kafkaConsumer.Run(consumerCtx)
	log.Printf("Kafka consumer started for topic: %s", cfg.Kafka.Topic)

	mux := router.NewRouter(storage)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		log.Printf("Starting HTTP server on %v", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
	
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown error: %v", err)
	}
	
	consumerCancel()
	
	log.Println("Server exited gracefully")
}
