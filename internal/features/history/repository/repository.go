package core_history_repository

import core_pgx_pool "github.com/c1kzy/golang-bankapp/internal/core/postgres"

type HistoryRepository struct {
	db *core_pgx_pool.Pool
}

func NewHistoryRepository(db *core_pgx_pool.Pool) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}
