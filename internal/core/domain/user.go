package domain

import (
	"fmt"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName string
	Balance  *int
}

func NewUnitializedUser(fullName string, balance int) User {
	return User{
		ID:       UnitializedID,
		Version:  UnitializedVersion,
		FullName: fullName,
		Balance:  &balance,
	}
}

func NewUser(
	id int,
	version int,
	fullName string,
	balance int,
) User {
	return User{
		ID:       id,
		Version:  version,
		FullName: fullName,
		Balance:  &balance,
	}
}

func (u *User) Validate() error {
	if u.FullName == "" {
		return fmt.Errorf("Full name cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	fullNameLength := len([]rune(u.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid FullName length: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.Balance == nil {
		return fmt.Errorf("Balance is required: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if *u.Balance < 0 {
		return fmt.Errorf("Balance cannot be negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
