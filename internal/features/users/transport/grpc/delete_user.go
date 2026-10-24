package core_user_transport_grpc

import (
	"context"
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteUser(
	ctx context.Context,
	req *users.DeleteUserRequest,
) (*emptypb.Empty, error) {
	if req.GetId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid user id: %s", req.GetId()))
	}

	if err := s.userService.DeleteUser(ctx, int(req.Id)); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid ID")
	}

	return &emptypb.Empty{}, nil
}
