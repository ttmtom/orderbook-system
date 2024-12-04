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
)

func main() {
	logger.Init()

	c := config.New()
	db := postgres.New(*c.DatabaseConfig)
	km := kafka.NewKafkaManager(c)
	v := validator.New()

	security.InitJwtSecurity(c.AppConfig.SecurityKey)
	moduleContainer := module.InitModuleContainer(db, v, c, km)
	middlewareContainer := router.InitMiddlewareContainer(c.AppConfig, moduleContainer)

	r := router.NewRouter(
		c.HttpConfig,
		moduleContainer,
		middlewareContainer,
	)

	err := r.Serve()
	if err != nil {
		slog.Error("Error on Echo Start", err)
		panic(err)
	}
}
