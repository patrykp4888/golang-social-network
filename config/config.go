package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Environment string
	Address     string
	ApiVersion  string
	DB          dbConfig
}

type dbConfig struct {
	Address      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		Address:     os.Getenv("API_PORT"),
		ApiVersion:  os.Getenv("API_VERSION"),
		DB: dbConfig{
			Address:      os.Getenv("DB_ADDRESS"),
			MaxOpenConns: GetInt(os.Getenv("DB_MAX_OPEN_CONNECTIONS")),
			MaxIdleConns: GetInt(os.Getenv("DB_MAX_IDLE_CONNECTIONS")),
			MaxIdleTime:  os.Getenv("DB_MAX_IDLE_TIME"),
		},
	}, nil
}

func GetInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return value
}
