package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

func TestParameter(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	params := []string{"-mode", "testmode", "-configpath", "/foo", "-archivedata", "true"}

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = append([]string{"params"}, params...)

	mode, configPath, archiveData := parseArguments()
	if mode != "testmode" {
		log.Fatalf("Got unexpected mode result, got %s, expected %s", "testmode", mode)
	}
	if configPath != "/foo" {
		log.Fatalf("Got unexpected config path result, got %s, expected %s", "/foo", configPath)
	}
	if !archiveData {
		log.Fatalf("Got unexpected archive data flag result, got %t, expected %t", true, archiveData)
	}
}

func TestParameterDefaults(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	params := []string{}

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = append([]string{"defaults"}, params...)

	mode, configPath, archiveData := parseArguments()
	if mode != "cli" {
		log.Fatalf("Got unexpected mode result, got %s, expected %s", "cli", mode)
	}
	if configPath != "/var/lib/timetracker" {
		log.Fatalf("Got unexpected config path result, got %s, expected %s", "/var/lib/timetracker", configPath)
	}
	if archiveData {
		log.Fatalf("Got unexpected archive data flag result, got %t, expected %t", false, archiveData)
	}
}

func TestParameterHelp(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	params := []string{"-help"}

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = append([]string{"help"}, params...)

	mode, configPath, archiveData := parseArguments()
	if mode != "" {
		log.Fatalf("Got unexpected mode result, got %s, expected %s", "", mode)
	}
	if configPath != "" {
		log.Fatalf("Got unexpected config path result, got %s, expected %s", "", configPath)
	}
	if archiveData {
		log.Fatalf("Got unexpected archive data flag result, got %t, expected %t", false, archiveData)
	}
}
