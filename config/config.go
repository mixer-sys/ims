package config

import (
	"fmt"
	"strconv"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port     string `env:"SERVER_PORT" envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"db"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"user"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	Name     string `env:"DB_NAME" envDefault:"db"`
	SSLMode  string `env:"SSLMode" envDefault:"disable"`
	Driver   string `env:"GOOSE_DRIVER" envDefault:"postgres"`
}

type Config struct {
	Server                  ServerConfig `env:"SERVER"`
	DB                      DBConfig     `env:"DATABASE"`
	ReadHeaderTimeoutSecond int          `env:"READ_HEADER_TIMEOUT_SECOND" envDefault:"5"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if _, err := strconv.Atoi(cfg.Server.Port); err != nil {
		return fmt.Errorf("invalid server port: %s", cfg.Server.Port)
	}
	if _, err := strconv.Atoi(cfg.DB.Port); err != nil {
		return fmt.Errorf("invalid database port: %s", cfg.DB.Port)
	}

	return nil
}
