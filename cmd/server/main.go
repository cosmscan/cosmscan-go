package main

import (
	"cosmscan-go/api"
	"cosmscan-go/cmd/utils"
	"flag"
	"fmt"

	"go.uber.org/zap"
)

var flagConfigFile = flag.String("config-file", "", "path to config file")

func main() {
	zap.L().Info("lets run the server !")

	flag.Parse()
	if *flagConfigFile == "" {
		panic("config-file flag is required")
	}

	// load config
	cfg := utils.MustLoadServerConfig(*flagConfigFile)

	// set up logger
	utils.InitAndReplaceLogger(cfg.Log)
	defer func() {
		err := zap.L().Sync()
		if err != nil {
			panic(fmt.Errorf("failed to sync logger during shutdown: %w", err))
		}
	}()

	sv := api.NewServer(cfg)
	if err := sv.Serve(); err != nil {
		zap.L().Error("failed to run server", zap.Error(err))
	}
}
