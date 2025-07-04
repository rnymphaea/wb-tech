package database

import (
	"log"
	"context"
	
	"wb-tech-l0/internal/database/postgres"
	"wb-tech-l0/internal/database/cache/redis"
	"wb-tech-l0/internal/database/models"
)

type Storage struct {
	cache *redis.RedisCache
	db *postgres.PostgresRepo
}

func New(cache *redis.RedisCache, db *postgres.PostgresRepo) *Storage {
	return &Storage{
		cache: cache,
		db: db,
	}
}

func (s *Storage) GetOrderByUID(ctx context.Context, uid string) (*models.Order, error) {
	if order, err := s.cache.GetOrder(ctx, uid); err == nil {
		log.Printf("Order %s retrieved from cache", uid)
		return order, nil
	}

	order, err := s.db.GetOrderByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if err := s.cache.SetOrder(ctx, order); err != nil {
		log.Printf("Failed to cache order %s: %v", uid, err)
	}
	
	return order, nil
}
