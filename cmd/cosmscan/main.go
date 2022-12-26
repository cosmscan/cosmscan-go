package main

import (
	"context"
	"cosmscan-go/indexer"
	"cosmscan-go/pkg/env"
	"errors"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"syscall"
)

var flagConfigFile = flag.String("config-file", "", "path to config file")

func main() {
	flag.Parse()
	if *flagConfigFile == "" {
		panic("config-file flag is required")
	}

	// load config
	cfg, err := indexer.LoadConfig(*flagConfigFile)
	if err != nil {
		panic(fmt.Errorf("failed load the config file: %v", err))
	}

	// setting up logger
	var logger *zap.Logger
	if env.FromString(cfg.Log.Environment) == env.Unknown {
		panic(fmt.Errorf("unknown log environment: %s", cfg.Log.Environment))
	}
	if env.FromString(cfg.Log.Environment) == env.Production {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}
	zap.ReplaceGlobals(logger)
	defer func() {
		if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			panic(fmt.Errorf("failed to sync logger: %v", err))
		}
	}()

	// set up indexer
	app, err := indexer.NewIndexer(cfg)
	if err != nil {
		logger.Error("failed to create indexer app", zap.Error(err))
	}

	if err := app.Run(context.Background()); err != nil {
		logger.Error("failed to run indexer app", zap.Error(err))
	}
}
