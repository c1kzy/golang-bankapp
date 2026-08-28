package domain

import (
	"fmt"
	"time"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

type TransactionType string

type TransactionStatus string

const (
	Deposit    TransactionType = "Deposit"
	Withdrawal TransactionType = "Withdrawal"
	Pay        TransactionType = "Pay"
	Transfer   TransactionType = "Transfer"

	Completed TransactionStatus = "Completed"
)

type Transaction struct {
	ID                int
	Version           int
	Amount            int
	TransactionType   TransactionType
	TransactionStatus TransactionStatus

	CreatedAt *time.Time

	AuthorUserID int
	ReceiverID   *int
}

func NewTransaction(
	amount int,
	transactionType TransactionType,
	transactionStatus TransactionStatus,
	authorUserID int,
	receiverID *int,
) Transaction {
	now := time.Now()
	return Transaction{
		ID:                UnitializedID,
		Version:           UnitializedVersion,
		Amount:            amount,
		TransactionType:   transactionType,
		TransactionStatus: transactionStatus,
		CreatedAt:         &now,
		AuthorUserID:      authorUserID,
		ReceiverID:        receiverID,
	}
}

func NewUserBalance(transactionType TransactionType, amount int, balance int) (int, error) {
	var newBalance int
	switch transactionType {
	case Deposit:
		newBalance = amount + balance
	case Withdrawal:
		newBalance = balance - amount
	case Transfer:
		newBalance = balance - amount

	default:
		return 0, fmt.Errorf("unknown operation: %v", transactionType)
	}

	if newBalance < 0 {
		return 0, fmt.Errorf("new balance cannot be negative: %w", core_errors.ErrBadRequest)
	}

	return newBalance, nil
}
