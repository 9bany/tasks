package util

import (
	"os"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	IframeURL     string `mapstructure:"IFRAMELY_URL"`
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func LoadConfig() Config {
	return Config{
		DBDriver:      getenv("DB_DRIVER", "postgres"),
		DBSource:      getenv("DB_SOURCE", "postgresql://root:secret@localhost:5432/task?sslmode=disable"),
		ServerAddress: getenv("SERVER_ADDRESS", "0.0.0.0:8080"),
		IframeURL:     getenv("IFRAMELY_URL", "https://iframe.ly"),
	}
}
