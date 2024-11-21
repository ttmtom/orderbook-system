package config

type HttpConfig struct {
	Host string
	Port string
}

func LoadHttpConfig(envFile map[string]string) *HttpConfig {
	return &HttpConfig{
		Host: envFile["HTTP_HOST"],
		Port: envFile["HTTP_PORT"],
	}
}
