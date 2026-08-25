package core_http_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("user validate error: %w", err)
	}
	newUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}
