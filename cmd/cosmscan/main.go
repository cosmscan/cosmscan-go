package main

import (
	"flag"
)

var flagConfigFile = flag.String("config-file", "", "path to config file")

func main() {
	flag.Parse()
	if *flagConfigFile == "" {
		panic("config-file flag is required")
	}
}
