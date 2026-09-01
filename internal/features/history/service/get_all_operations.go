package core_history_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

func (s *HistoryService) GetAllOperations(
	ctx context.Context,
	userID int,
) ([]domain.Transaction, error) {
	transactions, err := s.historyRepository.GetAllOperations(ctx, userID)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("unable to get transactions: %w", err)
	}

	if len(transactions) == 0 {
		return []domain.Transaction{}, fmt.Errorf("no transactions yet:%w", core_errors.ErrNotFound)
	}

	return transactions, nil
}
