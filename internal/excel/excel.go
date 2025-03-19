package excel

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"

	"ronnyfriedland/timetracker/v2/internal/logic"
)

var SheetName = time.Now().Format("2006")
var Headers = []string{"Date", "From", "To", "Duration"}

const dateLayout = "02.01.2006"
const timeLayout = "15:04:05"

func Export(configPath *string, duration logic.Duration) string {
	archiveDataFile := *configPath + "/timetracker.xlsx"

	file, err := excelize.OpenFile(archiveDataFile)
	if err != nil {
		file = excelize.NewFile()
	}
	defer func() {
        err := file.Close()
        if err != nil {
            log.Fatal(err)
        }
    }()

	_, err = file.NewSheet(SheetName)
	if err != nil {
		log.Fatalf("Failed to ensure sheet created: %v", err)
	}
	file.DeleteSheet("Sheet1")

	for i, header := range Headers {
		   file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	rows, err := file.GetRows(SheetName)
	if err != nil {
		log.Fatalf("Failed to get rows from sheet: %v", err)
	}

	next := len(rows) + 1

	file.SetCellValue(SheetName, fmt.Sprintf("A%d", next), duration.Date.Format(dateLayout))
	file.SetCellValue(SheetName, fmt.Sprintf("B%d", next), duration.StartTime.Format(timeLayout))
	file.SetCellValue(SheetName, fmt.Sprintf("C%d", next), duration.EndTime.Format(timeLayout))
	file.SetCellValue(SheetName, fmt.Sprintf("D%d", next), fmt.Sprintf("%2.2f", duration.Duration.Hours()))

	if err := file.SaveAs(archiveDataFile);
	err != nil {
		log.Fatal(err)
	}

	return archiveDataFile
}
