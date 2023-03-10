package app

import (
	"cosmscan-go/internal/db"
	"cosmscan-go/modules/indexer"
	"cosmscan-go/pkg/log"
	"flag"
	"fmt"

	"github.com/weaveworks/common/server"
)

type Config struct {
	// Target is the target to run.
	Target string `yaml:"target"`

	Logging  log.Config     `yaml:"logging,omitempty"`
	Server   server.Config  `yaml:"server,omitempty"`
	Database db.Config      `yaml:"database,omitempty"`
	Indexer  indexer.Config `yaml:"indexer,omitempty"`
}

type Registerer interface {
	RegisterFlags(*flag.FlagSet)
}

func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	c.RegisterDefaults(fs, c.Server, c.Database, c.Indexer)
}

func (c *Config) RegisterDefaults(fs *flag.FlagSet, settings ...interface{}) {
	for _, v := range settings {
		switch r := v.(type) {
		case Registerer:
			r.RegisterFlags(fs)
		default:
			panic(fmt.Sprintf("%v are not Registerer", v))
		}
	}
}
