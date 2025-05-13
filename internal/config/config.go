package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func Load() *Config {
	_ = godotenv.Load()

	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	return &Config{
		Host: apiHost,
		Port: apiPort,
	}

}
