package core_http_respository

import (
	"context"
	"fmt"
	"time"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO bankapp.users (full_name, balance)
		VALUES ($1, $2)
		RETURNING id, version, full_name, balance
		`

	var userModel UserModel

	row := r.db.QueryRow(ctx, query, user.FullName, user.Balance)
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Balance,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("row scan error: %w", err)
	}

	newUser := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.Balance,
	)

	return newUser, nil

}
