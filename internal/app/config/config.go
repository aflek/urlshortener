package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Version string `env:"VERSION"`
	Mode    string `env:"MODE" envDefault:"DEV"`

	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"ADDR" envDefault:"8080"`

	// Адрес запуска HTTP-сервера (зачем? лучше отдельно делать как выше Host
	// Port - не нужно разделять, поределять есть порт или нет и т.д.)
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	// Базовый адрес результирующего сокращённого URL
	BaseURL string `env:"BASE_URL" envDefault:"localhost:8080"`
}

func Load() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ABC() (s string) {
	return "aaabbb"
}
