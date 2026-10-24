package core_user_transport_grpc

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(
	ctx context.Context,
	req *users.GetUserRequest,
) (*users.User, error) {
	if req.GetId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid user id: %s", req.GetId()))
	}

	user, err := s.userService.GetUser(ctx, int(req.GetId()))
	if err != nil {
		return nil, fmt.Errorf("Get user error: %w", err)
	}

	return &users.User{
		Id:       int64(user.ID),
		Version:  int64(user.Version),
		FullName: user.FullName,
		Balance:  int64(*user.Balance),
	}, nil
}
