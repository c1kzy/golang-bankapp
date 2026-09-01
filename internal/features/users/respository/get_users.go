package core_http_respository

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *UserRepository) GetUsers(
	ctx context.Context,
) ([]domain.User, error) {
	query := `
	SELECT id, version, full_name, balance
	FROM bankapp.users
	ORDER BY id DESC;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to get users from db: %w", err)
	}
	defer rows.Close()

	var usersModel []UserModel
	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.Balance,
		)

		if err != nil {
			return nil, fmt.Errorf("get users scan error: %w", err)
		}

		usersModel = append(usersModel, userModel)
	}

	usersDomain, err := userModelsToDomain(usersModel)
	if err != nil {
		return nil, fmt.Errorf("unable to convert user models to domain: %w", err)
	}
	return usersDomain, nil
}
