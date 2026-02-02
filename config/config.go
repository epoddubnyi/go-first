package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBConfig struct {
	Host     string `env:"DB_HOST" yaml:"host"`
	Port     string `env:"DB_PORT" yaml:"port"`
	User     string `env:"DB_USER" yaml:"user"`
	Password string `env:"DB_PASSWORD" yaml:"password"`
	Name     string `env:"DB_NAME" yaml:"name"`
}

var Dsn string

func LoadConfig() (*DBConfig, error) {
	cfg := &DBConfig{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	Dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	return cfg, nil
}
