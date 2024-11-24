package config

type AppConfig struct {
	Env         string
	SecurityKey string
}

func LoadAppConfig(envFile map[string]string) *AppConfig {

	return &AppConfig{
		Env:         envFile["APP_ENV"],
		SecurityKey: envFile["APP_SECRET_KEY"],
	}
}
