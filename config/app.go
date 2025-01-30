package config

import "os"

type AppConfig struct {
	Env              string
	SecurityKey      string
	AdminBuild       bool
	AdminSecurityKey string
}

func LoadAppConfig() *AppConfig {

	return &AppConfig{
		Env:              os.Getenv("APP_ENV"),
		SecurityKey:      os.Getenv("APP_SECRET_KEY"),
		AdminBuild:       os.Getenv("APP_ADMIN_BUILD") == "true",
		AdminSecurityKey: os.Getenv("APP_ADMIN_SECRET_KEY"),
	}
}
