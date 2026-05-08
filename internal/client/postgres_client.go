package client

import (
	"context"
	"ecom/internal/config"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(ctx context.Context, cfg *config.PostgresConfig) (*PostgresClient, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		zap.L().Error(
			"failed to connect to postgres",
			zap.Error(err),
		)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		zap.L().Error(
			"failed to ping postgres",
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info("successfully connected to postgres")

	return &PostgresClient{pool: pool}, nil
}

func (c *PostgresClient) QueryAll(
	ctx context.Context,
	query string,
	params pgx.NamedArgs,
	result any,
) error {
	rows, err := c.pool.Query(ctx, query, params)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := pgxscan.ScanAll(result, rows); err != nil {
		return err
	}

	return nil
}

func (c *PostgresClient) QueryOne(
	ctx context.Context,
	query string,
	params pgx.NamedArgs,
	result any,
) error {
	rows, err := c.pool.Query(ctx, query, params)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := pgxscan.ScanOne(result, rows); err != nil {
		return err
	}

	return nil
}

func (c *PostgresClient) Exec(
	ctx context.Context,
	query string,
	params pgx.NamedArgs,
) error {
	cmdTg, err := c.pool.Exec(ctx, query, params)
	if err != nil {
		zap.L().Error(
			"failed to execute query",
			zap.String("query", query),
			zap.Any("params", params),
			zap.Error(err),
		)
		return err
	}

	if cmdTg.RowsAffected() == 0 {
		zap.L().Warn(
			"query executed but no rows were affected",
			zap.String("query", query),
			zap.Any("params", params),
			zap.Int64("rows_affected", cmdTg.RowsAffected()))
	}

	return nil
}
