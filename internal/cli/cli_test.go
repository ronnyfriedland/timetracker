package cli

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRunNotComplete(t *testing.T) {
	// Test if there is no log message if modification date is current
	var logContent bytes.Buffer

	log.SetOutput(&logContent)

	directory, _ := createStatusFile()
	archiveData := false

	Run(&directory, &archiveData)

	if logContent.String() != "" {
		log.Fatalf("Expected empty logmessage")
	}
}

func TestRunComplete(t *testing.T) {
	// Test if there are log messages if modification date is in the past
	var logContent bytes.Buffer

	log.SetOutput(&logContent)

	directory, fileName := createStatusFile()
	archiveData := false

	setModificationDate(fileName, "28.02.2022")

	Run(&directory, &archiveData)

	if logContent.String() == "" {
		log.Fatalf("Expected logmessage")
	}
	if !strings.Contains(logContent.String(), "Work duration: 0.25h") {
		log.Fatalf("Expected duration in logmessage")
	}
	if !strings.Contains(logContent.String(), "Start: 07:45:00, End: 08:00:00") {
		log.Fatalf("Expected start and end date in logmessage")
	}

}

func setModificationDate(fileName string, dateString string) {
	// Set modification time to past
	timestamp, _ := time.Parse(dateLayout, dateString)

	err3 := os.Chtimes(fileName, timestamp, timestamp)
	if err3 != nil {
		log.Fatal(err3)
	}
}

func createStatusFile() (string, string) {
	// Create a temporary directory
	directory, err := os.MkdirTemp("", "unittest")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(directory)

	fileName := directory + "/status"

	// Create a test status file
	file, err2 := os.Create(fileName)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file.Close()

	// Add test content to status file
	file.WriteString("2022-03-01T07:45:00+01:00\n")
	file.WriteString("2022-03-01T08:00:00+01:00\n")

	return directory, fileName
}
