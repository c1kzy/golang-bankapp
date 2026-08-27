package core_http_respository

import (
	"context"
	"errors"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.db.OpTimeOut)
	defer cancel()
	query := `
		SELECT id, version, full_name, balance
		FROM bankapp.users
		WHERE id=$1;
		`

	var userModel UserModel

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Balance,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user %w",
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("get user row scan error: %w", err)

	}

	user := userModelToDomain(userModel)

	return user, nil

}
