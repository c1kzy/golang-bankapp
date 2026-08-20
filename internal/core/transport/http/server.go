package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	mux    *http.ServeMux
	config Config

	log *logrus.Logger
}

func NewHTTPServer(config Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		log:    logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: s.mux,
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
