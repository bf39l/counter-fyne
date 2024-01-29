package counter

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func initialCounterUI(app *fyne.App) (*counterUI, *counterSettings) {
	ui := &counterUI{
		label:       generateTextLabel(),
		btnPlus:     widget.NewButtonWithIcon("", theme.ContentAddIcon(), nil),
		btnMinus:    widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), nil),
		settingsBtn: widget.NewButtonWithIcon("", theme.SettingsIcon(), nil),
		app:         app,
	}
	settings := &counterSettings{}

	return ui, settings
}

func generateTextLabel() *canvas.Text {
	label := canvas.NewText("Label", nil)
	label.TextSize = 50
	label.Alignment = fyne.TextAlignCenter
	return label
}
