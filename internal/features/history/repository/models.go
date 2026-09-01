package core_history_repository

import (
	"time"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type HistoryModel struct {
	ID                 int
	Version            int
	Amount             int
	Transaction_type   domain.TransactionType
	Transaction_status domain.TransactionStatus

	Created_at *time.Time

	AuthorUserID int
	ReceiverID   *int
}

func historyModelToDomain(historyModel HistoryModel) domain.Transaction {
	return domain.Transaction{
		ID:                historyModel.ID,
		Version:           historyModel.Version,
		Amount:            historyModel.Amount,
		TransactionType:   historyModel.Transaction_type,
		TransactionStatus: historyModel.Transaction_status,
		CreatedAt:         historyModel.Created_at,
		AuthorUserID:      historyModel.AuthorUserID,
		ReceiverID:        historyModel.ReceiverID,
	}
}

func historyModelsToDomain(historyModels []HistoryModel) []domain.Transaction {
	transactionsDomain := make([]domain.Transaction, len(historyModels))
	for i, historyModel := range historyModels {
		transactionsDomain[i] = historyModelToDomain(historyModel)
	}

	return transactionsDomain
}
