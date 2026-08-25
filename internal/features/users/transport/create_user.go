package core_http_transport

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Balance  int    `json:"balance"`
}

type CreateUserResponse UserResponse

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	var request CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode create user request")

		return
	}

	newUser := newUserFromDTO(request)

	user, err := h.userService.CreateUser(ctx, newUser)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to create user")

		return
	}

	userResponse := CreateUserResponse(userDomainToDTO(user))

	responseHandler.JSONResponse(userResponse, http.StatusOK)

}
