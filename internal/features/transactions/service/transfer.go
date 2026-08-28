package core_transactions_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

func (s *TransactionsService) Transfer(
	ctx context.Context,
	transaction domain.Transaction,
) (domain.Transaction, error) {
	if transaction.ReceiverID == nil {
		return domain.Transaction{}, fmt.Errorf(
			"receiver_id cannot be nil: %w",
			core_errors.ErrInvalidArgument,
	)
	}

	//Sender info
	user, err := s.userService.GetUser(ctx, transaction.AuthorUserID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("unable to get user: %w", err)
	}

	newTransaction := domain.NewTransaction(
		transaction.Amount,
		domain.Transfer,
		domain.Completed,
		transaction.AuthorUserID,
		transaction.ReceiverID,
	)

	newUserBalance, err := domain.NewUserBalance(
		domain.Transfer,
		transaction.Amount,
		*user.Balance,
	)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("new user balance error: %w", err)
	}

	transferUser, err := s.transactionsRepository.Transfer(ctx, newTransaction)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("pay error: %w", err)
	}

	patchUser := domain.UserPatch{
		FullName: user.FullName,
		Balance:  &newUserBalance,
	}

	_, err = s.userService.PatchUser(ctx, patchUser, user.ID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("patch user error: %w", err)
	}

	//Update receiver info
	receiverUser, err := s.userService.GetUser(ctx, *transaction.ReceiverID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("unable to get receiver: %w", err)
	}

	newReceiverBalance, err := domain.NewUserBalance(
		domain.Deposit,
		transaction.Amount,
		*receiverUser.Balance,
	)

	patchReceiverUser := domain.UserPatch{
		FullName: receiverUser.FullName,
		Balance:  &newReceiverBalance,
	}

	_, err = s.userService.PatchUser(ctx, patchReceiverUser, receiverUser.ID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("patch user error: %w", err)
	}

	return transferUser, nil
}
