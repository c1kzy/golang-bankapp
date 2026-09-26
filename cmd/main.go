package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/c1kzy/golang-bankapp/generated/users"
	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_middleware "github.com/c1kzy/golang-bankapp/internal/core/middleware"
	core_pgx_pool "github.com/c1kzy/golang-bankapp/internal/core/postgres"
	core_grpc_server "github.com/c1kzy/golang-bankapp/internal/core/transport/grpc"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_history_repository "github.com/c1kzy/golang-bankapp/internal/features/history/repository"
	core_history_service "github.com/c1kzy/golang-bankapp/internal/features/history/service"
	core_history_transport "github.com/c1kzy/golang-bankapp/internal/features/history/transport"
	core_transactions_repository "github.com/c1kzy/golang-bankapp/internal/features/transactions/repository"
	core_transactions_service "github.com/c1kzy/golang-bankapp/internal/features/transactions/service"
	core_transactions_transport "github.com/c1kzy/golang-bankapp/internal/features/transactions/transport"
	core_http_respository "github.com/c1kzy/golang-bankapp/internal/features/users/respository"
	core_http_service "github.com/c1kzy/golang-bankapp/internal/features/users/service"
	core_user_transport_grpc "github.com/c1kzy/golang-bankapp/internal/features/users/transport/grpc"
	core_user_transport "github.com/c1kzy/golang-bankapp/internal/features/users/transport/http"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logConfig, err := core_logger.NewConfig()
	if err != nil {
		fmt.Println("failed to get logger config", err)

		panic(err)
	}
	logger, err := core_logger.NewLogger(logConfig)
	if err != nil {
		fmt.Println("failed to initialize logger")

		panic(err)
	}

	logger.Info("Initializing postgres pool")
	poolConfig, err := core_pgx_pool.NewConfig()
	if err != nil {
		panic(err)
	}
	pool, err := core_pgx_pool.NewPostgresPool(ctx, poolConfig)

	serverConfig, err := core_http_server.NewConfig()
	if err != nil {
		panic(err)
	}

	grpcConfig, err := core_grpc_server.NewConfig()
	if err != nil {
		panic(err)
	}

	logger.Info("Initializing user feature")
	userRepository := core_http_respository.NewUserRepository(pool)
	userService := core_http_service.NewUserService(userRepository)
	userHandler := core_user_transport.NewUserHTTPHandler(userService)

	logger.Info("Initializing transactions feature")
	transactionsRepository := core_transactions_repository.NewTransactionsRepository(pool)
	transactionsService := core_transactions_service.NewTransactionsService(transactionsRepository, userService)
	transactionsHandler := core_transactions_transport.NewTransactionsHandler(transactionsService)

	logger.Info("Initializing history feature")
	historyRepository := core_history_repository.NewHistoryRepository(pool)
	historyService := core_history_service.NewHistoryService(historyRepository)
	historyTransport := core_history_transport.NewHistoryHandler(historyService)

	logger.Info("Initializing GRPC server")
	grpcServer := core_grpc_server.NewServer(grpcConfig, logger)

	userGRPCServer := core_user_transport_grpc.NewServer(userService)

	users.RegisterUserServiceServer(
		grpcServer.GRPCServer(),
		userGRPCServer,
	)

	go func() {
		if err := grpcServer.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	logger.Info("Initializing HTTP server")
	server := core_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_middleware.Logger(logger),
	)

	server.RegisterRoutes(userHandler.Routes()...)
	server.RegisterRoutes(transactionsHandler.Routes()...)
	server.RegisterRoutes(historyTransport.Routes()...)

	if err := server.Run(ctx); err != nil {
		logger.Error("Server run error", err)
	}

}
