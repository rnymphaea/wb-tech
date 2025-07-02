package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(path string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), path)
	if err != nil {
		return nil, err
	}

	return &Storage{pool: pool}, nil
}
