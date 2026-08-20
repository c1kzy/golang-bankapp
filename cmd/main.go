package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/c1kzy/golang-bankapp/internal/core/logger"
	core_http_server "github.com/c1kzy/golang-bankapp/internal/core/transport/http"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logConfig, err := core_logger.NewConfig()
	if err != nil {
		fmt.Println("failed to init logger", err)
		os.Exit(1)
	}
	logger, err := core_logger.NewLogger(logConfig)
	if err != nil {
		fmt.Println("failed to initialize logger")

		os.Exit(1)
	}

	serverConfig, err := core_http_server.NewConfig()
	if err != nil {
		panic(err)
	}

	server := core_http_server.NewHTTPServer(serverConfig, logger)

	if err := server.Run(ctx); err != nil {
		logger.Error("Server run error", err)
	}
}
