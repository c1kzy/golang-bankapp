package core_transactions_repository

import core_pgx_pool "github.com/c1kzy/golang-bankapp/internal/core/postgres"

type TransactionsRepository struct {
	db *core_pgx_pool.Pool
}

func NewTransactionsRepository(
	db *core_pgx_pool.Pool,
	) *TransactionsRepository {
		return &TransactionsRepository{
			db: db,
		}
}
