package db

import "flag"

type Config struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.Driver, "db.driver", "", "database driver, currently only sqlite and postresql drivers are available.")
	fs.StringVar(&cfg.Host, "db.host", "localhost", "database host")
	fs.IntVar(&cfg.Port, "db.port", 5432, "database port")
	fs.StringVar(&cfg.User, "db.user", "", "database user")
	fs.StringVar(&cfg.Password, "db.password", "", "database password")
	fs.StringVar(&cfg.Database, "db.database", "", "database name")
}
