package core_transactions_repository

import (
	"time"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type TransactionModel struct {
	ID                 int
	Version            int
	Amount             int
	Transaction_type   domain.TransactionType
	Transaction_status domain.TransactionStatus

	Created_at *time.Time

	AuthorUserID int
	ReceiverID   *int
}

func modelToDomain(transactionModel TransactionModel) domain.Transaction {
	return domain.Transaction{
		ID:                transactionModel.ID,
		Version:           transactionModel.Version,
		Amount:            transactionModel.Amount,
		TransactionType:   transactionModel.Transaction_type,
		TransactionStatus: transactionModel.Transaction_status,
		CreatedAt:         transactionModel.Created_at,
		AuthorUserID:      transactionModel.AuthorUserID,
		ReceiverID:        transactionModel.ReceiverID,
	}
}
