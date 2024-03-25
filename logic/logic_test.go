package logic

import (
	"os"
	"testing"
	"time"
)

func TestAddCurrentToStatus(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Unable to create temporary file %v", err)
	}
	defer os.Remove(file.Name())

	// Get the current time
	currentTime := time.Now()

	// Call the method with the temporary file
	addCurrentToStatus(file.Name(), currentTime)

	// Read the content of the file
	content, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Unable to read file %v", err)
	}

	// Check if the content is the current time in RFC3339 format
	if currentTime.Format(time.RFC3339)+"\n" != string(content) {
		t.Fatalf("Status content does not match expected content")
	}

}

func TestGetDurationFromStatus(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Unable to create temporary file %v", err)
	}
	defer os.Remove(file.Name())

	// Write some content to the file
	content := "2022-03-01T10:00:00Z\n2022-03-01T12:00:00Z\n"
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Unable to write to file %v", err)
	}

	// Call the method with the temporary file
	duration, startTime, endTime := getDurationFromStatus(file.Name())

	if 2*time.Hour != duration {
		t.Fatalf("Duration does not match expected duration")
	}

	if time.Date(2022, 3, 1, 10, 0, 0, 0, time.UTC) != startTime {
		t.Fatalf("Start time does not match expected start time")
	}

	if time.Date(2022, 3, 1, 12, 0, 0, 0, time.UTC) != endTime {
		t.Fatalf("End time does not match expected end time")
	}
}

func TestCleanupStatus(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Unable to create temporary file %v", err)
	}
	defer os.Remove(file.Name())

	// Write some content to the file
	content := "2022-03-01T10:00:00Z\n2022-03-01T12:00:00Z\n"
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Unable to write to file %v", err)
	}

	// Call the method with the temporary file
	cleanupStatus(file.Name())

	var truncated_content []byte

	// Read the content of the file
	truncated_content, err = os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Unable to read file %v", err)
	}

	if string(truncated_content) != "" {
		t.Fatalf("Content of truncated file does not match expected content")
	}
}
