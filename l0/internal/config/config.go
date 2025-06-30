package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"dbname"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return &config, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &config, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &config, nil
}

func (config *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName)
}

