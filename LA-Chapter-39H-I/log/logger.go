package log

import (
	"golang-chapter-39/LA-Chapter-39H-I/config"

	"go.uber.org/zap"
)

func InitZapLogger(cfg config.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	if cfg.AppDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}
