package core_transactions_service

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type TransactionsRepository interface {
	Deposit(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Withdrawal(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Transfer(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

type UserService interface {
	GetUser(ctx context.Context, id int) (domain.User, error)
	PatchUser(ctx context.Context, patch domain.UserPatch, id int) (domain.User, error)
}

type TransactionsService struct {
	transactionsRepository TransactionsRepository
	userService            UserService
}

func NewTransactionsService(
	transactionsRepository TransactionsRepository,
	userService UserService,
) *TransactionsService {
	return &TransactionsService{
		transactionsRepository: transactionsRepository,
		userService:            userService,
	}
}
