package database

import (
	"context"
	"log"
	"time"

	"wb-tech-l0/internal/database/cache/redis"
	"wb-tech-l0/internal/database/models"
	"wb-tech-l0/internal/database/postgres"
)

type Storage struct {
	cache *redis.RedisCache
	db    *postgres.PostgresRepo
}

func New(cache *redis.RedisCache, db *postgres.PostgresRepo) *Storage {
	return &Storage{
		cache: cache,
		db:    db,
	}
}

func (s *Storage) GetOrderByUID(ctx context.Context, uid string) (*models.Order, error) {
	start := time.Now()

	if order, err := s.cache.GetOrder(ctx, uid); err == nil {
		log.Printf("[Cache] Order %s retrieved from cache in %v", uid, time.Since(start))
		return order, nil
	}

	order, err := s.db.GetOrderByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if err := s.cache.SetOrder(ctx, order); err != nil {
		log.Printf("Failed to cache order %s: %v", uid, err)
	}

	log.Printf("[DB] Order %s retrieved from database in %v", uid, time.Since(start))

	return order, nil
}
