package core_http_transport

import (
	"context"
	"net/http"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
}

type UserHTTPHandler struct {
	userService UserService
}

func NewUserHTTPHandler(userService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
	}
}
