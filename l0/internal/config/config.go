package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type DatabaseConfig struct {
	Host         string `env:"DB_HOST,required"`
	Port         int    `env:"DB_PORT" envDefault:"5432"`
	Username     string `env:"DB_USER,required"`
	Password     string `env:"DB_PASSWORD,required"`
	DatabaseName string `env:"DB_NAME,required"`
}

type RedisConfig struct {
	Addr string        `env:"REDIS_ADDR,required"`
	TTL  time.Duration `env:"REDIS_TTL" envDefault:"3600s"`
}

type KafkaConfig struct {
	Brokers []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
	Topic   string   `env:"KAFKA_TOPIC" envDefault:"orders"`
	GroupID string   `env:"KAFKA_GROUP_ID" envDefault:"orders-consumer-group"`
}

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (config *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName)
}

func (config *Config) GetRedisTTL() time.Duration {
	return time.Duration(config.Redis.TTL) * time.Second
}
