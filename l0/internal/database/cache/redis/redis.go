package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"wb-tech-l0/internal/database/models"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(addr string, ttl time.Duration) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return &RedisCache{
		client: client,
		ttl:    ttl,
	}
}

func (c *RedisCache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *RedisCache) SetOrder(ctx context.Context, order *models.Order) error {
	key := fmt.Sprintf("order:%s", order.UID)
	jsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, jsonData, c.ttl).Err()
}

func (c *RedisCache) GetOrder(ctx context.Context, uid string) (*models.Order, error) {
	key := fmt.Sprintf("order:%s", uid)
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := json.Unmarshal(val, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (c *RedisCache) LoadAllOrders(ctx context.Context, getAllFunc func() ([]*models.Order, error)) error {
	orders, err := getAllFunc()
	if err != nil {
		return err
	}

	for _, order := range orders {
		if err := c.SetOrder(ctx, order); err != nil {
			return err
		}
	}
	return nil
}
