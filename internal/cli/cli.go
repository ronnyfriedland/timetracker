package cli

import (
	"log"
	"ronnyfriedland/timetracker/v2/internal/logic"
)

const dateLayout = "02.01.2006"
const timeLayout = "15:04:05"

func Run(configPath *string) {
	duration := logic.Execute(configPath)
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
