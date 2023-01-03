package utils

import (
	"cosmscan-go/config"
	"cosmscan-go/pkg/env"
	"fmt"

	"go.uber.org/zap"
)

func InitAndReplaceLogger(cfg config.LogConfig) {
	var logger *zap.Logger
	var err error

	if env.FromString(cfg.Environment) == env.Unknown {
		panic(fmt.Errorf("unknown log environment: %s", cfg.Environment))
	}

	if env.FromString(cfg.Environment) == env.Production {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}

	zap.ReplaceGlobals(logger)
}
