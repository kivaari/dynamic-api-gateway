package main

import (
	"net/http"
	"time"

	"github.com/kivaari/dynamic-api-gateway/internal/config"
	"github.com/kivaari/dynamic-api-gateway/internal/gateway"
	"github.com/kivaari/dynamic-api-gateway/internal/logger"
	"github.com/kivaari/dynamic-api-gateway/internal/server"
)

func main() {
	logger.Init()

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		logger.Log.Fatalf("Failed to load config: %v", err)
	}

	router := gateway.NewRouter(cfg)

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + string(cfg.Server.Port),
		Handler: router,
	}

	timeout, _ := time.ParseDuration(cfg.Server.GracefulTimeout)
	server.RunServer(srv, timeout)
}
