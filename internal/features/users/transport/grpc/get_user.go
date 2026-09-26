package core_user_transport_grpc

import (
	"context"

	"github.com/c1kzy/golang-bankapp/generated/users"
)

func (s *Server) GetUser(
	ctx context.Context,
	req *users.GetUserRequest,
) (*users.User, error) {

	user, err := s.userService.GetUser(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	var balance int64

	if user.Balance != nil {
		balance = int64(*user.Balance)
	}

	return &users.User{
		Id:       int64(user.ID),
		Version:  int64(user.Version),
		FullName: user.FullName,
		Balance:  balance,
	}, nil
}
