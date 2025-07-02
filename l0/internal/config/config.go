package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DatabaseConfig struct {
	Host         string `env:"DB_HOST"`
	Port         int    `env:"DB_PORT"`
	Username     string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
}

type Config struct {
	Database DatabaseConfig
}

func LoadConfig() (*Config, error) {
	var dbcfg DatabaseConfig

	if err := env.Parse(&dbcfg); err != nil {
		return nil, err
	}
	
	return &Config{Database: dbcfg}, nil
}

func (config *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName)
}
