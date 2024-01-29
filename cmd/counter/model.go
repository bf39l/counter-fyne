package counter

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type counterData struct {
	count int
	step  int
}

type counterUI struct {
	label       *canvas.Text
	btnPlus     *widget.Button
	btnMinus    *widget.Button
	settingsBtn *widget.Button
	app         *fyne.App
}

type counterSettings struct {
}

type CounterSvc struct {
	Data *counterData
	Gui  *counterUI
}
