package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"time"
)

const dateLayout = "02.01.2006"
const timeLayout = "15:04:05"

func main() {
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	flag.Parse()

	var currentTime = time.Now()
	var modificationTime time.Time

	timeTrackerStatusFile := *configPath + "/status"
	fileInfo, err := os.Stat(timeTrackerStatusFile)
	if err != nil {
		log.Printf("Error getting file status: %v", err)
		modificationTime = currentTime
	} else {
		modificationTime = fileInfo.ModTime()
	}

	if math.Abs(float64(modificationTime.Day()-currentTime.Day())) == 0 {
		// same day, add to status
		addCurrentToStatus(timeTrackerStatusFile, currentTime)
	} else {
		duration, from, to := getDurationFromStatus(timeTrackerStatusFile)
		log.Printf("[%s] - Work duration: %2.2fh",
			modificationTime.Format(dateLayout),
			duration.Hours())
		log.Printf("[%s] - Start: %s, End: %s",
			modificationTime.Format(dateLayout),
			from.Format(timeLayout),
			to.Format(timeLayout))

		// remove data from status
		cleanupStatus(timeTrackerStatusFile)

		// new day, add first status
		addCurrentToStatus(timeTrackerStatusFile, currentTime)
	}
}

func cleanupStatus(timeTrackerStatusFile string) {
	if err := os.Truncate(timeTrackerStatusFile, 0); err != nil {
		log.Fatalf("Failed to cleanup status: %v", err)
	}
}

func getDurationFromStatus(timeTrackerStatusFile string) (time.Duration, time.Time, time.Time) {
	file, err := os.Open(timeTrackerStatusFile)
	if err != nil {
		log.Fatalf("Error opening status: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Unable to close file %v", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	first, err := time.Parse(time.RFC3339, lines[0])
	if err != nil {
		log.Fatalf("Error getting first date of status: %v", err)
	}
	last, err := time.Parse(time.RFC3339, lines[len(lines)-1])
	if err != nil {
		log.Fatalf("Error getting last date of status: %v", err)
	}
	duration := last.Sub(first)
	return duration, first, last
}

func addCurrentToStatus(timeTrackerStatusFile string, currentTime time.Time) {
	statusDateString := currentTime.Format(time.RFC3339) + "\n"

	file, err := os.OpenFile(timeTrackerStatusFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Unable to open file %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to close file %v", err)
		}
	}(file)

	_, err = file.WriteString(statusDateString)
	if err != nil {
		log.Fatalf("Unable to write to status file %v", err)
	}
}
