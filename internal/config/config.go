package config

import (
	"fmt"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"INFO"`
	DbHost     string `env:"POSTGRES_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"user"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DbName     string `env:"DB_NAME" envDefault:"db"`
	SSLMode    string `env:"SSLMode" envDefault:"disable"`
}

func LoadConfig() (cfg *Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}
	return cfg, nil
}
