package core_transactions_transport

import "github.com/c1kzy/golang-bankapp/internal/core/domain"

type TransactionRequest struct {
	Amount       int  `json:"amount"`
	AuthorUserID int  `json:"author_user_id"`
	ReceiverID   *int `json:"receiver_id"`
}

func TransactionDTOtoDomain(request TransactionRequest) domain.Transaction {
	return domain.Transaction{
		ID:           domain.UnitializedID,
		Version:      domain.UnitializedVersion,
		Amount:       request.Amount,
		AuthorUserID: request.AuthorUserID,
		ReceiverID:   request.ReceiverID,
	}
}
