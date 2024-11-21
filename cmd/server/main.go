package main

import (
	"github.com/go-playground/validator"
	"log/slog"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/handler"
	"orderbook/internal/adapter/router"
	"orderbook/internal/core/service"
	"orderbook/pkg/logger"
	"os"
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

	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(v, userService)

	r := router.NewRouter(
		c.HttpConfig,
		userHandler,
	)

	err = r.Serve()
	if err != nil {
		slog.Error("Error on Echo Start", err)
		os.Exit(1)
	}
}
