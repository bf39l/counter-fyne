package counter

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func NewCounterService(app *fyne.App) *CounterSvc {
	data := &counterData{
		count: 0,
		step:  1,
	}
	ui, _ := initialCounterUI(app)
	svc := &CounterSvc{
		Data: data,
		Gui:  ui,
	}
	svc.Gui.Window = svc.createUIWindow()
	svc.bindBtnEvts()

	return svc
}

func (c *CounterSvc) bindBtnEvts() {
	// Update btn tap func
	c.Gui.btnPlus.OnTapped = c.btnClick(1)
	c.Gui.btnMinus.OnTapped = c.btnClick(-1)

	c.Gui.btnSettings.OnTapped = c.btnClickSettings()
	// TODO:
}

func (c *CounterSvc) createUIWindow() fyne.Window {
	window := (*c.Gui.app).NewWindow("Counter")
	window.Resize(fyne.Size{Height: 120, Width: 255})

	// Reset Label
	c.Gui.label.Text = fmt.Sprintf("%d", c.Data.count)

	// Group UI elements
	// Top, Btm, Left, Right, Others...
	content := container.NewBorder(
		container.NewGridWithRows(2,
			c.Gui.label,
			container.NewGridWithColumns(2,
				c.Gui.btnPlus, c.Gui.btnMinus,
			),
		),
		c.Gui.btnSettings,
		nil,
		nil,
	)
	window.SetContent(content)
	// Focus on plus button by default
	window.Canvas().Focus(c.Gui.btnPlus)
	return window
}

func (c *CounterSvc) btnClick(mul int) func() {
	return func() {
		c.updateCount(mul)
		c.Gui.label.Text = fmt.Sprintf("%d", c.Data.count)
		c.Gui.label.Refresh()
	}
}

func (c *CounterSvc) btnClickSettings() func() {
	return func() {
		notImplementedDialog := dialog.NewInformation("Settings", "Not Implemented yet", c.Gui.Window)
		notImplementedDialog.Resize(fyne.NewSize(50, 50))
		notImplementedDialog.Show()
	}
}

func (c *CounterSvc) updateCount(mul int) {
	c.Data.count = c.Data.count + mul*c.Data.step
}

func (u *CounterSvc) updateUI() {
	// TODO
}
