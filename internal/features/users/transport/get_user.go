package core_http_transport

import (
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

func (h *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	id, err := core_http_server.GetIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "get path value error")

		return
	}

	userDomain, err := h.userService.GetUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "get user error")

		return
	}

	user := userDomainToDTO(userDomain)

	responseHandler.JSONResponse(user, http.StatusOK)
}
