package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func New(path string) (*PostgresRepo, error) {
	pool, err := pgxpool.New(context.Background(), path)
	if err != nil {
		return nil, err
	}

	return &PostgresRepo{pool: pool}, nil
}

func (p *PostgresRepo) Close() {
	p.pool.Close()
}
