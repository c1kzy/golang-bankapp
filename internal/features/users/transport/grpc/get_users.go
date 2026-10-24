package core_user_transport_grpc

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
)

func (s *Server) GetUsers(
	ctx context.Context,
	req *users.GetUsersRequest,
) (*users.GetUsersResponse, error) {

	usersDomain, err := s.userService.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("Get users error: %w", err)
	}

	var usersResponse []*users.User

	for _, user := range usersDomain {
		usersResponse = append(usersResponse, &users.User{
			Id:       int64(user.ID),
			Version:  int64(user.Version),
			FullName: user.FullName,
			Balance:  int64(*user.Balance),
		})
	}

	return &users.GetUsersResponse{
		Users: usersResponse,
	}, nil
}
