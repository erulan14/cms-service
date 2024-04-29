package pgxdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, host, port, username, password, database string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", username, password, host, port, database))

	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
