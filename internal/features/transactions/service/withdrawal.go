package core_transactions_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *TransactionsService) Withdrawal(
	ctx context.Context,
	transaction domain.Transaction,
) (domain.Transaction, error) {
	user, err := s.userService.GetUser(ctx, transaction.AuthorUserID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("unable to get user for withdrawal: %w", err)
	}

	newTransaction := domain.NewTransaction(
		transaction.Amount,
		domain.Withdrawal,
		domain.Completed,
		transaction.AuthorUserID,
		transaction.ReceiverID,
	)

	newUserBalance, err := domain.NewUserBalance(domain.Withdrawal, transaction.Amount, *user.Balance)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("new user balance error: %w", err)
	}

	withDrawalTransaction, err := s.transactionsRepository.Withdrawal(ctx, newTransaction)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("withdrawal error: %w", err)
	}

	patchUser := domain.UserPatch{
		FullName: user.FullName,
		Balance:  &newUserBalance,
	}

	_, err = s.userService.PatchUser(ctx, patchUser, user.ID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("patch user for withdrawal error: %w", err)
	}

	return withDrawalTransaction, nil
}
