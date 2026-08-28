package core_transactions_repository

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *TransactionsRepository) Deposit(
	ctx context.Context,
	transaction domain.Transaction,
) (domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, r.db.OpTimeOut)
	defer cancel()

	query := `
		INSERT INTO bankapp.transactions (amount, transaction_type, transaction_status, created_at, author_user_id, receiver_id)
		VALUES($1,$2,$3,$4,$5,$6)
		RETURNING id, version, amount, transaction_type, transaction_status, created_at, author_user_id, receiver_id;
	`

	row := r.db.QueryRow(
		ctx,
		query,
		transaction.Amount,
		transaction.TransactionType,
		transaction.TransactionStatus,
		transaction.CreatedAt,
		transaction.AuthorUserID,
		transaction.ReceiverID,
	)

	var transactionModel TransactionModel

	err := row.Scan(
		&transactionModel.ID,
		&transactionModel.Version,
		&transactionModel.Amount,
		&transactionModel.Transaction_type,
		&transactionModel.Transaction_status,
		&transactionModel.Created_at,
		&transactionModel.AuthorUserID,
		&transactionModel.ReceiverID,
	)

	if err != nil {
		return domain.Transaction{}, fmt.Errorf("transcation row scan error: %w", err)
	}

	transactionDomain := modelToDomain(transactionModel)

	return transactionDomain, nil

}
