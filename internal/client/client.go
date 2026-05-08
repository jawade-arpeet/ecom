package client

import (
	"context"
	"ecom/internal/config"
)

type Client struct {
	Postgres      *PostgresClient
	Redis         *RedisClient
	ElasticSearch *ElasticSearchClient
}

func NewClient(ctx context.Context, cfg *config.Config) (*Client, error) {
	pg, err := newPostgresClient(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	rds, err := newRedisClient(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

	es, err := newElasticSearchClient(cfg.ElasticSearch)
	if err != nil {
		return nil, err
	}

	return &Client{
		Postgres:      pg,
		Redis:         rds,
		ElasticSearch: es,
	}, nil
}