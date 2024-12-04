package main

import (
	"log/slog"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/adapter/router"
	"orderbook/internal/core/module"
	"orderbook/internal/pkg/security"
	"orderbook/internal/pkg/validator"
	"orderbook/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	c := config.New()
	db := postgres.New(*c.DatabaseConfig)
	km := kafka.NewKafkaManager(c)
	v := validator.New()

	security.InitJwtSecurity(c.AppConfig.SecurityKey)
	moduleContainer := module.InitModuleContainer(db, v, c, km)

	r := router.NewRouter(
		c.HttpConfig,
		moduleContainer,
	)

	go func() {
		err := r.Serve()
		if err != nil {
			slog.Error("Error on Echo Start", err)
			panic(err)
		}
	}()
	km.StartPolling()

	<-quit
	slog.Info("Shutting down server...")
	km.CloseAll()
}
