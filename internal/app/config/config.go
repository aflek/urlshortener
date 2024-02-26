package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	// Адрес запуска HTTP-сервера
	ServerAddress string `env:"SERVER_ADDRESS"`

	// Базовый адрес результирующего сокращённого URL
	BaseURL string `env:"BASE_URL"`

	// Куда пишем лог
	LoggerPath string `env:"LOGGER_PATH" envDefault:""`
	// Уровень логирования
	LoggerLevel string `env:"LOGGER_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {
	// Приоритет параметров сервера:
	// 1. Если указаны переменнае окружения, то используется они.
	// 2. Если нет переменных окружения, но есть аргументы командной строки (флаг), то используется они.
	// 3. Если нет ни переменной окружения, ни флага, то используется значение по умолчанию.

	// 1. parse env with defaut value
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	// 2. parese command line and set params if empty
	// format:  -a=localhost:8888 -b=http://localhost:8080
	if len(os.Args) > 1 {
		// define flags
		a := flag.String("a", "", "server url (format: localhost:8888)")
		b := flag.String("b", "", "response url (format: localhost:8080)")

		flag.Parse()

		if a != nil && cfg.ServerAddress == "" {
			cfg.ServerAddress = *a
		}

		if b != nil && cfg.BaseURL == "" {
			cfg.BaseURL = *b
		}
	}

	// 3. default values
	if cfg.ServerAddress == "" {
		cfg.ServerAddress = "localhost:8080"
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "http://localhost:8080"
	}

	return &cfg, nil
}
