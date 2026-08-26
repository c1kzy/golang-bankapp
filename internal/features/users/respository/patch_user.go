package core_http_respository

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *UserRepository) PatchUser(
	ctx context.Context,
	user domain.User,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.db.OpTimeOut)
	defer cancel()

	query := `
	UPDATE bankapp.users
	SET full_name=$1,
		balance=$2,
		version=version+1
	WHERE id=$3
	RETURNING id,version, full_name, balance;
	`

	var userModel UserModel
	row := r.db.QueryRow(ctx, query, user.FullName, user.Balance, user.ID)
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Balance,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("postgres patch user error: %w", err)
	}

	patchedUser := userModelToDomain(userModel)

	return patchedUser, nil

}
