package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/kivaari/dynamic-api-gateway/internal/config"
	"github.com/kivaari/dynamic-api-gateway/internal/gateway"
	"github.com/kivaari/dynamic-api-gateway/internal/logger"
	"github.com/kivaari/dynamic-api-gateway/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("Не удалось загрузить .env файл, используем значения по умолчанию")
	}

	logger.Init()

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		logger.Log.Fatalf("Failed to load config: %v", err)
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.Security.JWT.Secret = jwtSecret
	}

	router := gateway.NewRouter(cfg)

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port),
		Handler: router,
	}

	timeout, _ := time.ParseDuration(cfg.Server.GracefulTimeout)
	server.RunServer(srv, timeout)
}
