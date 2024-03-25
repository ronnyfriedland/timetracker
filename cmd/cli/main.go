package main

import (
	"flag"
	"log"
	"ronnyfriedland/timetracker/v2/logic"
)

const dateLayout = "02.01.2006"
const timeLayout = "15:04:05"

func main() {
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	flag.Parse()

	execute(configPath)
}

func execute(directory *string) {
	duration := logic.Execute(directory)
	if duration.Complete {
		log.Printf("[%s] - Work duration: %2.2fh",
			duration.Date.Format(dateLayout),
			duration.Duration.Hours())

		log.Printf("[%s] - Start: %s, End: %s",
			duration.Date.Format(dateLayout),
			duration.StartTime.Format(timeLayout),
			duration.EndTime.Format(timeLayout))
	}

}
