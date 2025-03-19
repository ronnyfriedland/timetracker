package main

import (
	"flag"
	"os"
	"ronnyfriedland/timetracker/v2/internal/cli"
)

func main() {
	mode, configPath, archiveData := parseArguments()
	if mode == "" {
		os.Exit(0)
	} else if mode == "cli" {
		cli.Run(&configPath, &archiveData)
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func parseArguments() (string, string, bool) {
	mode := flag.String("mode", "cli", "the application mode, available: cli")
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	archiveData := flag.Bool("archivedata", false, "flag to enable data archiving")
	help := flag.Bool("help", false, "print this help message")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return "", "", false
	} else {
		return *mode, *configPath, *archiveData
	}
}
