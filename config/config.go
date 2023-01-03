package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ServerConfig is the configuration for the server.
type ServerConfig struct {
	Log  LogConfig  `json:"log" yaml:"log"`
	DB   DBConfig   `json:"db" yaml:"db"`
	Http HTTPConfig `json:"http" yaml:"http"`
}

// IndexerConfig is the configuration for the indexer.
type IndexerConfig struct {
	RPCEndpoint string      `json:"rpc_endpoint" yaml:"rpc_endpoint"`
	StartBlock  uint64      `json:"start_block" yaml:"start_block"`
	Log         LogConfig   `json:"log" yaml:"log"`
	Chain       ChainConfig `json:"chain" yaml:"chain"`
	DB          DBConfig    `json:"db" yaml:"db"`
}

type LogConfig struct {
	Environment string `json:"environment" yaml:"environment"`
}

type ChainConfig struct {
	ID   string `json:"chain_id" yaml:"chain_id"`
	Name string `json:"chain_name" yaml:"chain_name"`
}

type DBConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

type HTTPConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	CORS bool   `json:"cors" yaml:"cors"`
}

func LoadIndexerConfig(filename string) (*IndexerConfig, error) {
	var cfg IndexerConfig
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(contents, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadServerConfig(filename string) (*ServerConfig, error) {
	var cfg ServerConfig
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(contents, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
