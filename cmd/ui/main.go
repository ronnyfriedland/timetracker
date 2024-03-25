package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"ronnyfriedland/timetracker/v2/logic"
)

func main() {
	configPath := flag.String("configpath", "/var/lib/timetracker", "the config path")
	flag.Parse()

	duration := logic.Execute(configPath)
	if duration.Complete {

		persistStatus(duration)

	}

	a := app.New()
	w := a.NewWindow("Timetracker")

	list := widget.NewList(
		func() int {
			return 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			if i == 0 {
				o.(*widget.Label).SetText(duration.StartTime.GoString())
			}
			if i == 1 {
				o.(*widget.Label).SetText(duration.EndTime.GoString())
			}
			if i == 2 {
				o.(*widget.Label).SetText(duration.Duration.String())
			}
		})

	w.SetContent(list)

	w.ShowAndRun()
}


func persistStatus(duration logic.Duration) {
	file, err := os.Create("records.csv")
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{duration.Date.Format("02.01.2006"), duration.StartTime.Format("02.01.2006"), duration.EndTime.Format("02.01.2006"), duration.Duration.String()}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}
