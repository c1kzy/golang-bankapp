package core_http_service

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (r *UserService) GetUsers(
	ctx context.Context,
) ([]domain.User, error) {

	users, err := r.userRepository.GetUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}
