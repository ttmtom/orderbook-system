package config

import "os"

type HttpConfig struct {
	Host string
	Port string
}

func LoadHttpConfig() *HttpConfig {
	return &HttpConfig{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
