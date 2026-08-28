package core_transactions_transport

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type TransactionTransferRequest TransactionRequest

func (h *TransactionsHTTPHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	var request TransactionTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode pay request")

		return
	}

	payDomain := TransactionDTOtoDomain(TransactionRequest(request))

	transaction, err := h.transactionsService.Transfer(ctx, payDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "pay transaction error")

		return
	}

	responseHandler.JSONResponse(transaction, http.StatusOK)
}
