package main

import (
	"cosmscan-go/cmd/cosmscan/app"
	"cosmscan-go/pkg/log"
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Version is set via build flag -ldflags -X main.Version
var (
	Version  string
	Branch   string
	Revision string
)

const configFileOption = "config.file"

func loadConfig() (*app.Config, error) {
	var configFile string

	args := os.Args[1:]
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&configFile, configFileOption, "", "Path to the config file")
	for len(args) > 0 {
		_ = fs.Parse(args)
		args = args[1:]
	}

	if configFile == "" {
		fs.Usage()
		return nil, errors.New("unable to find config file")
	}

	c, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s, err: %w", configFile, err)
	}

	cfg := &app.Config{}
	if err := yaml.Unmarshal(c, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s, err: %w", configFile, err)
	}

	return cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := log.InitLogger(cfg.Logging); err != nil {
		fmt.Printf("failed to init logger, err: %v\n", err)
		return
	}

	log.Logger.Debugw("app is about to start", "version", Version, "branch", Branch, "revision", Revision)
}
