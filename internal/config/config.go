package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Version string `env:"VERSION"`
	Mode    string `env:"MODE" envDefault:"DEV"`

	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"ADDR" envDefault:"8080"`
}

func Load() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
