package core_http_response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ResponseHandler struct {
	w   http.ResponseWriter
	log *logrus.Entry
}

func NewResponseHandler(w http.ResponseWriter, log *logrus.Entry) *ResponseHandler {
	return &ResponseHandler{
		w:   w,
		log: log,
	}
}

func (h *ResponseHandler) JSONResponse(responseBody any, statusCode int) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("failed to send JSON response")
	}

	h.log.Info("Status code: ", statusCode)
}

func (h *ResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(args ...any)
	)

	switch {
	case errors.Is(err, ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn

	case errors.Is(err, ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug

	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg)
	h.errorResponse(statusCode, err, msg)

}

func (h *ResponseHandler) errorResponse(statusCode int, err error, msg string) {
	response := map[string]any{
		"error":   err.Error(),
		"message": msg,
	}

	h.JSONResponse(response, statusCode)

}
