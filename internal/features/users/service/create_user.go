package core_http_service

import (
	"context"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	user, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
