package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/bf39l/counter-fyne/cmd/counter"
)

func main() {
	app := app.New()

	counterSvc := counter.NewCounterService(&app)
	counterSvc.CreateUIWindow().Show()

	app.Run()
}
