package main

import (
	"flag"
)

// Version is set via build flag -ldflags -X main.Version
var (
	Version  string
	Branch   string
	Revision string
)

var flagConfigFile = flag.String("config-file", "", "path to config file")

func main() {
	flag.Parse()
	if *flagConfigFile == "" {
		panic("config-file flag is required")
	}
}
