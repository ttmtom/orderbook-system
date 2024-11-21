package config

type AppConfig struct {
	Env string
}

func LoadAppConfig(envFile map[string]string) *AppConfig {

	return &AppConfig{
		Env: envFile["APP_ENV"],
	}
}
