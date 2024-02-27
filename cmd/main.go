package main

import (
	"flag"
	"ronnyfriedland/timetracker/v2/internal/cli"
)

func main() {
	mode := flag.String("mode", "cli", "the application mode, available: cli")
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	flag.Parse()

	if *mode == "cli" {
		cli.Run(configPath)
	}
}
