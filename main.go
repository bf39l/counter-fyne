package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/bf39l/counter-fyne/counter"
)

func main() {
	app := app.New()

	counterSvc := counter.NewCounterService(&app)
	counterSvc.Gui.Window.Show()

	app.Run()
}
