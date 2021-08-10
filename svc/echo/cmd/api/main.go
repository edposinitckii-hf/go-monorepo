package main

import (
	"fmt"
	"net/http"

	"github.com/edposinitckii-hf/monorepo/svc/echo/internal/config"
	"github.com/edposinitckii-hf/monorepo/svc/echo/internal/service"
	"github.com/edposinitckii-hf/monorepo/svc/echo/internal/transport"
	"github.com/edposinitckii-hf/monorepo/toolkit/db/postgres"
	"github.com/edposinitckii-hf/monorepo/toolkit/log"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	logger, _ := log.NewLogger()
	db, _ := postgres.NewPostgresConnection(cfg.DSN)
	svc := service.NewEcho(db, logger)
	router := transport.NewRouter(svc)

	logger.Info("starting echo server", zap.Int("port", cfg.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		logger.Error("Shutdown with error", zap.Error(err))
		return
	}
}
