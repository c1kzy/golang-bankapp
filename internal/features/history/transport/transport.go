package core_history_transport

import (
	"context"
	"net/http"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
)

type HistoryService interface {
	GetAllOperations(ctx context.Context, userID int) ([]domain.Transaction, error)
	GetTransferOperations(ctx context.Context, userID int) ([]domain.Transaction, error)
}

type HistoryHTTPHandler struct {
	historyService HistoryService
}

func NewHistoryHandler(historyService HistoryService) *HistoryHTTPHandler {
	return &HistoryHTTPHandler{
		historyService: historyService,
	}
}

func (h *HistoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/transactions/all/{user_id}",
			Handler: h.GetAllOperations,
		},
		{
			Method:  http.MethodGet,
			Path:    "/transactions/transfer/{user_id}",
			Handler: h.GetTransferOperations,
		},
	}
}
