package main

import (
	"cosmscan-go/cmd/utils"
	"cosmscan-go/indexer"
	"flag"
	"fmt"

	"go.uber.org/zap"
)

var flagConfigFile = flag.String("config-file", "", "path to config file")

func main() {
	flag.Parse()
	if *flagConfigFile == "" {
		panic("config-file flag is required")
	}

	// load config
	cfg := utils.MustLoadIndexerConfig(*flagConfigFile)

	// setting up logger
	// it replaces the global logger according to the configuration
	utils.InitAndReplaceLogger(cfg.Log)
	defer func() {
		err := zap.L().Sync()
		if err != nil {
			panic(fmt.Errorf("failed to sync logger during shutdown: %w", err))
		}
	}()

	// set up indexer
	app, err := indexer.NewIndexer(cfg)
	if err != nil {
		zap.L().Error("failed to create indexer app", zap.Error(err))
	}

	if err := app.Run(); err != nil {
		zap.L().Error("failed to run indexer app", zap.Error(err))
	}
}
