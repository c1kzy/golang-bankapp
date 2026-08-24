package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_middleware "github.com/c1kzy/golang-bankapp/internal/core/middleware"
	core_middleware "github.com/c1kzy/golang-bankapp/internal/core/middleware"
)

type Server struct {
	mux    *http.ServeMux
	config Config

	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	logger *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *Server {
	return &Server{
		mux:    http.NewServeMux(),
		config: config,
		log:    logger,

		middleware: middleware,
	}
}

func (s *Server) Run(ctx context.Context) error {
	mux := core_middleware.ChainMiddleware(s.mux, s.middleware...)
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Infof("server running on %s", s.config.Addr)
		err := server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("Stopping server...")

		shutdownCtx, cancel := context.WithTimeout(ctx, s.config.ShutdownTimeOut)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server, %w", err)
		}

		s.log.Warn("Server stopped")
	}

	return nil
}

func (s *Server) RegisterRoutes(routes ...Route) {

	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		s.mux.Handle(pattern, route.WithMiddleware())
	}
}
