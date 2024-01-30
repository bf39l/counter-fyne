package counter

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func initialCounterUI(app *fyne.App) *counterUI {
	ui := &counterUI{
		label:       generateTextLabel(),
		btnPlus:     widget.NewButtonWithIcon("", theme.ContentAddIcon(), nil),
		btnMinus:    widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), nil),
		btnSettings: widget.NewButtonWithIcon("", theme.SettingsIcon(), nil),
		app:         app,
		settings: &counterSettings{
			items: generateSettingsForm(),
		},
	}

	return ui
}

func generateSettingsForm() map[string]*widget.FormItem {
	items := map[string]*widget.FormItem{}
	stepItem := widget.NewFormItem("Step", widget.NewEntry())
	stepItem.Widget.(*widget.Entry).Validator = validation.NewRegexp(`[-]{0,1}[1-9]+[0-9]{0,}`, "Step has to be an integer")
	items["step"] = stepItem

	return items
}

func generateTextLabel() *canvas.Text {
	label := canvas.NewText("Label", nil)
	label.TextSize = 50
	label.Alignment = fyne.TextAlignCenter
	return label
}
