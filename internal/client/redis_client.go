package client

import (
	"context"
	"ecom/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisClient struct {
	rds *redis.Client
}

func newRedisClient(
	ctx context.Context,
	cfg *config.RedisConfig,
) (*RedisClient, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	if err := rds.Ping(ctx).Err(); err != nil {
		zap.L().Error(
			"failed to connect to redis",
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info("successfully connected to redis")

	return &RedisClient{rds: rds}, nil
}

func (c *RedisClient) Set(
	ctx context.Context,
	key string,
	value string,
	expiration time.Duration,
) error {
	status := c.rds.Set(ctx, key, value, expiration)

	if err := status.Err(); err != nil {
		zap.L().Error(
			"failed to set value in redis",
			zap.String("key", key),
			zap.Error(err),
		)
		return err
	}

	zap.L().Info(
		"successfully set value in redis",
		zap.String("key", key),
	)

	return nil
}

func (c *RedisClient) Get(
	ctx context.Context,
	key string,
) (string, error) {
	cmd := c.rds.Get(ctx, key)

	if err := cmd.Err(); err != nil {
		zap.L().Error(
			"failed to get value from redis",
			zap.String("key", key),
			zap.Error(err),
		)
		return "", err
	}

	zap.L().Info(
		"successfully got value from redis",
		zap.String("key", key),
	)

	return cmd.Val(), nil
}
