package main

import (
	"log/slog"
	"orderbook/internal/core/module"
	"orderbook/internal/pkg/security"
	"os"

	"orderbook/config"
	"orderbook/internal/adapter/database/postgres"
	"orderbook/internal/adapter/router"
	"orderbook/internal/pkg/validator"
	"orderbook/pkg/logger"
)

func main() {
	logger.Init()

	c, err := config.New()
	if err != nil {
		slog.Error("Error loading configuration:", err)
		os.Exit(1)
	}

	db, err := postgres.New(*c.DatabaseConfig)
	if err != nil {
		slog.Error("Error on database connection", err)
		os.Exit(1)
	}

	v := validator.New()

	security.InitJwtSecurity(c.AppConfig.SecurityKey)

	moduleContainer := module.InitModuleContainer(db.DB, v)
	middlewareContainer := router.InitMiddlewareContainer(moduleContainer)

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
