package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"orderbook/config"
	"orderbook/pkg/logger"
	"orderbook/pkg/utils"

	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger.Init()

	actionList := []string{"up", "down"}
	argType := "up"
	if len(os.Args) > 1 {
		if !utils.Contains(actionList, os.Args[1]) {
			slog.Error("Invalid action type", os.Args)
			os.Exit(1)
		}

		argType = os.Args[1]
	}
	slog.Info(argType)

	pgConfig, err := config.NewMigration()
	if err != nil {
		slog.Error("Error on loading db config", err)
		os.Exit(1)
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.DBName,
	)

	sourceURL := "file://internal/adapter/database/postgres/migration"

	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)
	if err != nil {
		slog.Error("Error on db connection", err)
		os.Exit(1)
	}

	if argType == "up" {
		err = m.Up()
	} else if argType == "down" {
		err = m.Steps(-1)
	}

	if err != nil {
		slog.Error(fmt.Sprintf("Error on migration %s", argType), err)
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("Migration %s completed successfully", argType))
}
