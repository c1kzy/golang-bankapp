package core_http_service

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func (s *UserService) PatchUser(
	ctx context.Context,
	patch domain.UserPatch,
	id int,
) (domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("user validate error:%w", err)
	}

	if err := user.PatchUser(patch); err != nil {
		return domain.User{}, fmt.Errorf("patch user error: %w", err)
	}

	patchedUser, err := s.userRepository.PatchUser(ctx, user, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user error: %w", err)
	}

	return patchedUser, nil
}
