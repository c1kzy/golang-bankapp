package core_user_transport_http

import (
	"net/http"

	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_user_transport "github.com/c1kzy/golang-bankapp/internal/features/users/transport"
)

type UserHTTPHandler struct {
	userService core_user_transport.UserService
}

func NewUserHTTPHandler(userService core_user_transport.UserService) *UserHTTPHandler {
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
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
	}
}
