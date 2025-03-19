package excel

import (
	"log"
	"os"
	"ronnyfriedland/timetracker/v2/internal/logic"
	"testing"
	"time"

	"github.com/xuri/excelize/v2"
)

func TestExport(t *testing.T) {

	testData := logic.Duration{
		Date: time.Now(),
		StartTime: time.Now(),
		EndTime: time.Now(),
		Duration: time.Duration(1000),
	}

	excelDirectory := os.TempDir()
	excelFile := Export(&excelDirectory, testData)

	file, err := excelize.OpenFile(excelFile)
	if err != nil {
		log.Fatalf("Expected excel file")
	}

	rows, err := file.GetRows(SheetName)
	if err != nil {
		log.Fatalf("Expected no error getting rows")
	}

	if len(rows) != 2 {
		log.Fatalf("Unexpected line count - expected 2, got %d", len(rows))
	}

	headers := rows[0]
	for i := 0; i < len(headers); i++ {
		if headers[i] != Headers[i] {
			log.Fatalf("Unexpected header value - expected %s, got %s", Headers[i], headers[i])
		}
	}

	defer os.Remove(excelFile)
}
