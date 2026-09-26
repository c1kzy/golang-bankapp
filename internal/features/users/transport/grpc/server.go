package core_user_transport_grpc

import (
	"github.com/c1kzy/golang-bankapp/generated/users"
	core_user_transport "github.com/c1kzy/golang-bankapp/internal/features/users/transport"
)

type Server struct {
	users.UnimplementedUserServiceServer

	userService core_user_transport.UserService
}

func NewServer(userService core_user_transport.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
