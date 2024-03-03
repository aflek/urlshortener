package logger

import (
	"urlshortener/internal/app/config"

	"go.uber.org/zap"
)

type Logger struct{}

func New(srvConf *config.Config) (*zap.Logger, error) {
	// текстовый уровень логирования из конфига сервера
	level := srvConf.LoggerLevel
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl

	// если в кофиге приложения задан путь для записи логов
	if srvConf.LoggerPath != "" {
		cfg.OutputPaths = append(cfg.OutputPaths, srvConf.LoggerPath)
		cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, srvConf.LoggerPath)
	}

	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return zl, nil
}
