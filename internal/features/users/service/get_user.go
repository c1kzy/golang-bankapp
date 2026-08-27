package core_http_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}
