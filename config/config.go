package config

import (
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	AppConfig      *AppConfig
	HttpConfig     *HttpConfig
	DatabaseConfig *DatabaseConfig
}

func New() (*Config, error) {
	envFile, envFileLoadingErr := godotenv.Read(".env")
	if envFileLoadingErr != nil {
		return nil, envFileLoadingErr
	}

	slog.Info("Env File load successfully", "file", envFile)

	appConfig := LoadAppConfig(envFile)
	httpConfig := LoadHttpConfig(envFile)
	databaseConfig := LoadConfig(envFile)

	return &Config{
		appConfig,
		httpConfig,
		databaseConfig,
	}, nil
}

func NewMigration() (*DatabaseConfig, error) {
	envFile, envFileLoadingErr := godotenv.Read(".env")
	if envFileLoadingErr != nil {
		return nil, envFileLoadingErr
	}

	postgresConfig := LoadConfig(envFile)

	return postgresConfig, nil
}
