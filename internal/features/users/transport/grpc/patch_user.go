package core_user_transport_grpc

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PatchUser(
	ctx context.Context,
	req *users.PatchUserRequest,
) (*users.User, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid user ID: %v", req.GetId()))
	}

	user := DTOtoUserPatch(req)
	patchedUser, err := s.userService.PatchUser(ctx, user, int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("patch user error: %w", err)
	}

	return &users.User{
		Id:       int64(patchedUser.ID),
		Version:  int64(patchedUser.Version),
		FullName: patchedUser.FullName,
		Balance:  int64(*patchedUser.Balance),
	}, nil
}
