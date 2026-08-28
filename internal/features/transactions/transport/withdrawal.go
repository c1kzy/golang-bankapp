package core_transactions_transport

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_response "github.com/c1kzy/golang-bankapp/internal/core/transport/http/response"
)

type TransactionWithdrawalRequest TransactionRequest


func (h *TransactionsHTTPHandler) Withdrawal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(w, log)

	var request TransactionWithdrawalRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "transaction decode and validate error")

		return
	}

	transactionDomain := TransactionDTOtoDomain(TransactionRequest(request))

	transaction, err := h.transactionsService.Withdrawal(ctx, transactionDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "withdrawal error")

		return
	}

	responseHandler.JSONResponse(transaction, http.StatusOK)
}
