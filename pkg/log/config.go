package log

import (
	"flag"
)

type Config struct {
	Encoding string `yaml:"encoding"`
	Level    string `yaml:"log_level"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.Encoding, "log.encoding", "", "log encoding format (e.g. json | console)")
	fs.StringVar(&cfg.Level, "log.level", "info", "log level to use")
}
