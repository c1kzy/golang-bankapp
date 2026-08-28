package core_transactions_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *TransactionsService) Deposit(
	ctx context.Context,
	transaction domain.Transaction,
) (domain.Transaction, error) {
	user, err := s.userService.GetUser(ctx, transaction.AuthorUserID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("get user for transaction error: %w", err)
	}

	newTransaction := domain.NewTransaction(
		transaction.Amount,
		domain.Deposit,
		domain.Completed,
		transaction.AuthorUserID,
		transaction.ReceiverID,
	)

	newBalance, err := domain.NewUserBalance(domain.Deposit, transaction.Amount, *user.Balance)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("new balance error: %w", err)
	}

	deposit, err := s.transactionsRepository.Deposit(ctx, newTransaction)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("deposit error: %w", err)
	}

	patch := domain.UserPatch{
		FullName: user.FullName,
		Balance:  &newBalance,
	}

	_, err = s.userService.PatchUser(ctx, patch, user.ID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("patch user for transaction error:%w", err)
	}

	return deposit, nil
}
