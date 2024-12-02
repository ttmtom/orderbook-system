package main

import (
	"log/slog"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres"
	"orderbook/internal/adapter/router"
	"orderbook/internal/core/module"
	"orderbook/internal/pkg/security"
	"orderbook/internal/pkg/validator"
	"orderbook/pkg/logger"
	"os"
)

func main() {
	logger.Init()

	c, err := config.New()
	if err != nil {
		slog.Error("Error loading configuration:", err)
		panic(err)
	}

	db, err := postgres.New(*c.DatabaseConfig)
	if err != nil {
		slog.Error("Error on database connection", err)
		panic(err)
	}

	v := validator.New()

	security.InitJwtSecurity(c.AppConfig.SecurityKey)

	moduleContainer := module.InitModuleContainer(db.DB, v, c)
	middlewareContainer := router.InitMiddlewareContainer(c.AppConfig, moduleContainer)

	r := router.NewRouter(
		c.HttpConfig,
		moduleContainer,
		middlewareContainer,
	)

	err = r.Serve()
	if err != nil {
		slog.Error("Error on Echo Start", err)
		os.Exit(1)
	}
}
