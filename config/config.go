package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	AppConfig      *AppConfig
	HttpConfig     *HttpConfig
	DatabaseConfig *DatabaseConfig
}

func loadEnv() error {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func New() (*Config, error) {
	err := loadEnv()
	if err != nil {
		slog.Error("Error loading .env file", err)
		return nil, err
	}

	appConfig := LoadAppConfig()
	httpConfig := LoadHttpConfig()
	databaseConfig := LoadConfig()

	return &Config{
		appConfig,
		httpConfig,
		databaseConfig,
	}, nil
}

func NewMigration() (*DatabaseConfig, error) {
	err := loadEnv()
	if err != nil {
		slog.Error("Error loading .env file", err)
		return nil, err
	}

	postgresConfig := LoadConfig()

	return postgresConfig, nil
}
