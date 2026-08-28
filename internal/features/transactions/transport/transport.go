package core_transactions_transport

import (
	"context"
	"net/http"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
)

type TransactionsService interface {
	Deposit(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Withdrawal(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	Transfer(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

type TransactionsHTTPHandler struct {
	transactionsService TransactionsService
}

func NewTransactionsHandler(
	transactionService TransactionsService,
) *TransactionsHTTPHandler {
	return &TransactionsHTTPHandler{
		transactionsService: transactionService,
	}
}

func (h *TransactionsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/transactions/deposit",
			Handler: h.Deposit,
		},
		{
			Method:  http.MethodPost,
			Path:    "/transactions/withdrawal",
			Handler: h.Withdrawal,
		},
		{
			Method:  http.MethodPost,
			Path:    "/transactions/transfer",
			Handler: h.Transfer,
		},
	}
}
