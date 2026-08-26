package core_http_transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type CreateUserRequest UserRequest

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

	if err := request.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "validate request error")

		return
	}

	newUser, err := newUserFromDTO(request)
	if err != nil {
		responseHandler.ErrorResponse(err, "new user from DTO")

		return
	}

	user, err := h.userService.CreateUser(ctx, newUser)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to create user")

		return
	}

	userResponse := CreateUserResponse(userDomainToDTO(user))

	responseHandler.JSONResponse(userResponse, http.StatusOK)

}

func (r *CreateUserRequest) Validate() error {
	if r.FullName == "" {
		return fmt.Errorf("full_name cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if r.Balance == nil {
		return fmt.Errorf("balance cannot be nil: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
