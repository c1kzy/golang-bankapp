package core_pgx_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	opTimeOut time.Duration
}

func NewPostgresPool(ctx context.Context, config Config) (*Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return &Pool{}, err
	}

	return &Pool{
		Pool:      pool,
		opTimeOut: config.Timeout,
	}, nil
}
