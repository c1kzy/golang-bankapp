package core_user_transport_grpc

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(
	ctx context.Context,
	req *users.CreateUserRequest,
) (*users.User, error) {
	if err := Validate(req.FullName, req.Balance); err != nil {
		return nil, status.Error(codes.InvalidArgument, core_errors.InvalidArgumentMessage)
	}

	newUser, err := newUserFromMessage(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, core_errors.InvalidArgumentMessage)
	}

	user, err := s.userService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			core_errors.InvalidArgumentMessage)
	}

	return &users.User{
		Id:       int64(user.ID),
		Version:  int64(user.Version),
		FullName: user.FullName,
		Balance:  int64(*user.Balance),
	}, nil
}

func Validate(fullName string, balance int64) error {
	if fullName == "" {
		return fmt.Errorf("fullName cannot be empty: %w", core_errors.InvalidArgumentMessage)
	}

	if balance <= 0 {
		return fmt.Errorf("balance cannot be nil: %w", core_errors.InvalidArgumentMessage)
	}

	return nil
}
