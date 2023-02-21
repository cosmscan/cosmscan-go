package indexer

import "flag"

type Config struct {
	RPCEndpoint string `json:"rpc_endpoint" yaml:"rpc_endpoint"`
	StartBlock  uint64 `json:"start_block" yaml:"start_block"`
	ChainID     string `json:"chain_id" yaml:"chain_id"`
	ChainName   string `json:"chain_name" yaml:"chain_name"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.RPCEndpoint, "indexer.rpc_endpoint", "http://localhost:26657/", "database host")
	fs.Uint64Var(&cfg.StartBlock, "indexer.start_block", 1, "start block is a what indexer start to fetch first. Indexer sequentially fetch blocks")
	fs.StringVar(&cfg.ChainID, "indexer.chain_id", "", "chain id")
	fs.StringVar(&cfg.ChainName, "indexer.chain_name", "", "chain name for cosmos blockchain")
}
