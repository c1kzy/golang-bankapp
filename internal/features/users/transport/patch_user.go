package core_http_transport

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type UpdateUserRequest UserRequest

func (h *UserHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	id, err := core_http_server.GetIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to get id")

		return
	}

	var request UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode update user request")

		return
	}

	user := userDTOtoUserPatch(UserRequest(request))

	patchedUser, err := h.userService.PatchUser(ctx, user, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to patch user")

		return
	}

	userResponse := userDomainToDTO(patchedUser)

	responseHandler.JSONResponse(userResponse, http.StatusOK)
}
