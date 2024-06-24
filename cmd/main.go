package main

import (
	"context"

	"github.com/armanzhankin/echo-htmx/config"
	"github.com/armanzhankin/echo-htmx/internal/database"
	"github.com/armanzhankin/echo-htmx/internal/handler"
	"github.com/armanzhankin/echo-htmx/internal/server"
	"github.com/armanzhankin/echo-htmx/internal/service"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, _ := zap.NewProduction()
	sugaredLogger := logger.Sugar()

	cfg, err := config.LoadConfig("string")
	if err != nil {
		sugaredLogger.Fatalf("couldn't load config: %s", err.Error())
	}

	db, err := database.New(ctx, cfg.DB)

	service, err := service.NewService(ctx, db)
	if err != nil {
		sugaredLogger.Fatalf("couldn't create service layers: %s", err.Error())
	}

	handler := handler.NewHandler(service, sugaredLogger, cfg.SecretKey)

	if err = server.RunServer(cfg, handler.Routes()); err != nil {
		sugaredLogger.Fatalf("couldn't run server: %s", err.Error())
	}
}
