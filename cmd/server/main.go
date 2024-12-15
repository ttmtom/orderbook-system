package main

import (
	"log/slog"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres"
	"orderbook/internal/adapter/router"
	"orderbook/internal/core/module"
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
	v := validator.New()

	moduleContainer := module.InitModuleContainer(db, v, c)

	r := router.NewRouter(
		c.HttpConfig,
		c.AppConfig,
		moduleContainer,
	)

	go func() {
		err := r.Serve()
		if err != nil {
			slog.Error("Error on Echo Start", err)
			panic(err)
		}
	}()
	moduleContainer.EventModule.StartPolling()

	<-quit
	slog.Info("Shutting down server...")
	moduleContainer.EventModule.CloseAll()
}
