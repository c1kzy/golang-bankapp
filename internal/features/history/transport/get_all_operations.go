package core_history_transport

import (
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

func (h *HistoryHTTPHandler) GetAllOperations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	userID, err := core_http_server.GetIDPathValue(r, "user_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to get user_id")

		return
	}

	transactions, err := h.historyService.GetAllOperations(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to get transactions")

		return
	}

	responseHandler.JSONResponse(transactions, http.StatusOK)
}
