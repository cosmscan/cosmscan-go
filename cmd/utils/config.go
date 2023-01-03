package utils

import "cosmscan-go/config"

func MustLoadIndexerConfig(filename string) *config.IndexerConfig {
	cfg, err := config.LoadIndexerConfig(filename)
	if err != nil {
		panic(err)
	}
	return cfg
}
