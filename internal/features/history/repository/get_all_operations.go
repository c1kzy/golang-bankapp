package core_history_repository

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *HistoryRepository) GetAllOperations(
	ctx context.Context,
	userID int,
) ([]domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, r.db.OpTimeOut)
	defer cancel()

	query := `
	SELECT id, version, amount, transaction_type, transaction_status, created_at, author_user_id, receiver_id
	FROM bankapp.transactions
	WHERE author_user_id=$1
	OR receiver_id=$1
	ORDER BY created_at DESC;
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("query row error: %w", err)
	}

	var historyModels []HistoryModel

	for rows.Next() {
		var historyModel HistoryModel
		err := rows.Scan(
			&historyModel.ID,
			&historyModel.Version,
			&historyModel.Amount,
			&historyModel.Transaction_type,
			&historyModel.Transaction_status,
			&historyModel.Created_at,
			&historyModel.AuthorUserID,
			&historyModel.ReceiverID,
		)
		if err != nil {
			return []domain.Transaction{}, fmt.Errorf("row scan error: %w", err)
		}

		historyModels = append(historyModels, historyModel)
	}

	transactionsDomain := historyModelsToDomain(historyModels)

	return transactionsDomain, nil
}
