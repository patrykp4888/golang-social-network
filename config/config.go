package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Environment string
	Address     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		Address:     os.Getenv("API_PORT"),
	}, nil
}
