package core_http_respository

import (
	core_pgx_pool "github.com/c1kzy/golang-bankapp/internal/core/postgres"
)

type UserRepository struct {
	db *core_pgx_pool.Pool
}

func NewUserRepository(db *core_pgx_pool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
