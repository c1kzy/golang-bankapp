package core_history_service

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type HistoryRepository interface {
	GetAllOperations(ctx context.Context, userID int) ([]domain.Transaction, error)
	GetTransferOperations(ctx context.Context, userID int) ([]domain.Transaction, error)
}

type HistoryService struct {
	historyRepository HistoryRepository
}

func NewHistoryService(historyRepository HistoryRepository) *HistoryService {
	return &HistoryService{
		historyRepository: historyRepository,
	}
}
