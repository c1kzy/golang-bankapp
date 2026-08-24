package core_http_middleware

import (
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	"github.com/sirupsen/logrus"
)

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-REQUEST-ID")

			l := log.WithFields(logrus.Fields{
				"request_id": requestID,
				"url":        r.URL.String(),
			},
			)

			ctx := core_logger.ToContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
