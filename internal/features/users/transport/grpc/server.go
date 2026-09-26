package core_user_transport_grpc

import (
	"github.com/c1kzy/golang-bankapp/generated/users"
	core_user_transport_http "github.com/c1kzy/golang-bankapp/internal/features/users/transport/http"
)

type Server struct {
	users.UnimplementedUserServiceServer

	userService core_user_transport_http.UserService
}

func NewServer(userService core_user_transport_http.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
