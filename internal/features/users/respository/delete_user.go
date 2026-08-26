package core_http_respository

import (
	"context"
	"fmt"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.db.OpTimeOut)
	defer cancel()
	query := `
	DELETE FROM bankapp.users
	WHERE id=$1;
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user from db error: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id=%v: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
