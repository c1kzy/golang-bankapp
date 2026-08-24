package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_middleware "github.com/c1kzy/golang-bankapp/internal/core/middleware"
	core_pgx_pool "github.com/c1kzy/golang-bankapp/internal/core/postgres"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
	core_http_respository "github.com/c1kzy/golang-bankapp/internal/features/users/respository"
	core_http_service "github.com/c1kzy/golang-bankapp/internal/features/users/service"
	core_http_transport "github.com/c1kzy/golang-bankapp/internal/features/users/transport"
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

	logger.Info("Initializing user feature")
	userRepository := core_http_respository.NewUserRepository(pool)
	userService := core_http_service.NewUserService(userRepository)
	userHandler := core_http_transport.NewUserHTTPHandler(userService)

	logger.Info("Initializing HTTP server")
	server := core_http_server.NewHTTPServer(
		serverConfig,
		logger,
		core_middleware.Logger(logger),
	)

	server.RegisterRoutes(userHandler.Routes()...)

	if err := server.Run(ctx); err != nil {
		logger.Error("Server run error", err)
	}
}
