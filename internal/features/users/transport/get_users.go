package core_user_transport

import (
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type GetUsersResponse []UserResponse

func (h *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	usersDomain, err := h.userService.GetUsers(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to get users")

		return
	}

	usersDTO := GetUsersResponse(usersDomainToDTO(usersDomain))

	responseHandler.JSONResponse(usersDTO, http.StatusOK)

}
