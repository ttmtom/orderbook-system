package config

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig(envFile map[string]string) *DatabaseConfig {
	return &DatabaseConfig{
		User:     envFile["POSTGRES_USER"],
		Password: envFile["POSTGRES_PASSWORD"],
		Host:     envFile["POSTGRES_HOST"],
		Port:     envFile["POSTGRES_PORT"],
		DBName:   envFile["POSTGRES_DB"],
	}
}
