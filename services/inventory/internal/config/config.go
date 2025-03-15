package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	EnvLocal      = "local"
	EnvStaging    = "staging"
	EnvProduction = "production"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Port string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")

	if env == EnvLocal {
		if err := godotenv.Load(fmt.Sprintf("./configs/%s.env", EnvLocal)); err != nil {
			return nil, err
		}
	}

	return &Config{
		App: AppConfig{
			Port: os.Getenv("APP_PORT"),
		},
		Postgres: PostgresConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			DB:       os.Getenv("POSTGRES_DB"),
		},
	}, nil
}
