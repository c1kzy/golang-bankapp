package core_grpc_server

import (
	"context"
	"fmt"
	"net"
	"time"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	"google.golang.org/grpc"
)

type Server struct {
	config     Config
	log        *core_logger.Logger
	grpcServer *grpc.Server
}

func NewServer(config Config, logger *core_logger.Logger) *Server {
	return &Server{
		config:     config,
		log:        logger,
		grpcServer: grpc.NewServer(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.config.Addr)
	if err != nil {
		return err
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Infof("grpc server running on %s", s.config.Addr)
		err := s.grpcServer.Serve(listener)

		if err != nil {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and serve GRPC: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("Stopping GRPC server...")

		done := make(chan struct{})

		go func() {
			s.grpcServer.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			s.log.Warn("gRPC server stopped")

		case <-time.After(s.config.ShutdownTimeOut):
			s.log.Warn("gRPC graceful shutdown timeout")

			s.grpcServer.Stop()

			s.log.Warn("gRPC server forcefully stopped")
		}
	}

	return nil
}

func (s *Server) GRPCServer() *grpc.Server {
	return s.grpcServer
}

func (s *Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}
