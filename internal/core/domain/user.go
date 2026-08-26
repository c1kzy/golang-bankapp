package domain

import (
	"fmt"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

type UserPatch struct {
	FullName string
	Balance  *int
}

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

func (p *UserPatch) Validate() error {
	if p.FullName == "" {
		return fmt.Errorf("Full name cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	fullNameLength := len([]rune(p.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid FullName length: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Balance == nil {
		return fmt.Errorf("Balance is required: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if *p.Balance < 0 {
		return fmt.Errorf("Balance cannot be negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) PatchUser(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch error: %w", err)
	}

	tmp := *u

	tmp.FullName = patch.FullName
	tmp.Balance = patch.Balance

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil

}
