package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBConfig struct {
	Host     string `env:"DB_HOST" yaml:"host"`
	Port     int    `env:"DB_PORT" yaml:"port"`
	User     string `env:"DB_USER" yaml:"user"`
	Password string `env:"DB_PASSWORD" yaml:"password"`
}

var Dsn string

func LoadConfig() (*DBConfig, error) {
	fmt.Println(os.Getenv("DB_USER"))
	cfg := &DBConfig{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	Dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return cfg, nil
}
