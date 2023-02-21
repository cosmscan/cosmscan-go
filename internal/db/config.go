package db

import "flag"

type Config struct {
	UseInMemory bool   `yaml:"use_in_memory"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.UseInMemory, "db.use_memory", false, "determines whether use sqlite driver or real database")
	fs.StringVar(&cfg.Host, "db.host", "localhost", "database host")
	fs.IntVar(&cfg.Port, "db.port", 5432, "database port")
	fs.StringVar(&cfg.User, "db.user", "", "database user")
	fs.StringVar(&cfg.Password, "db.password", "", "database password")
	fs.StringVar(&cfg.Database, "db.database", "", "database name")
}
