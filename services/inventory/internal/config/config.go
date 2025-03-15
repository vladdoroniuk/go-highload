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
	App       AppConfig
	Cassandra CassandraConfig
}

type AppConfig struct {
	Port string
}

type CassandraConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Keyspace string
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
		Cassandra: CassandraConfig{
			User:     os.Getenv("CASSANDRA_USER"),
			Password: os.Getenv("CASSANDRA_PASSWORD"),
			Host:     os.Getenv("CASSANDRA_HOST"),
			Port:     os.Getenv("CASSANDRA_PORT"),
			Keyspace: os.Getenv("CASSANDRA_KEYSPACE"),
		},
	}, nil
}
