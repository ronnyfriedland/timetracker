package main

import (
	"flag"
	"os"
	"ronnyfriedland/timetracker/v2/internal/cli"
)

func main() {
	mode, configPath := parseArguments()
	if mode == "" {
		os.Exit(0)
	} else if mode == "cli" {
		cli.Run(&configPath)
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func parseArguments() (string, string) {
	mode := flag.String("mode", "cli", "the application mode, available: cli")
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	help := flag.Bool("help", false, "print this help message")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return "", ""
	} else {
		return *mode, *configPath
	}
}
