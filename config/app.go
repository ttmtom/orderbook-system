package config

import "os"

type AppConfig struct {
	Env         string
	SecurityKey string
}

func LoadAppConfig() *AppConfig {

	return &AppConfig{
		Env:         os.Getenv("APP_ENV"),
		SecurityKey: os.Getenv("APP_SECRET_KEY"),
	}
}
