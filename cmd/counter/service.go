package counter

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewCounterService(app *fyne.App) *CounterSvc {
	data := &counterData{
		count: 0,
		step:  1,
	}
	ui, _ := initialCounterUI(app)

	return &CounterSvc{
		Data: data,
		Gui:  ui,
	}
}

func (c *CounterSvc) CreateUIWindow() fyne.Window {
	window := (*c.Gui.app).NewWindow("Counter")
	window.Resize(fyne.Size{Height: 120, Width: 255})

	// Reset Label
	c.Gui.label.Text = fmt.Sprintf("%d", c.Data.count)
	// Update btn tap func
	c.Gui.btnPlus.OnTapped = c.btnClick(1)
	c.Gui.btnMinus.OnTapped = c.btnClick(-1)
	// TODO update settings btn tap func

	// Group UI elements
	// Top, Btm, Left, Right, Others...
	content := container.NewBorder(
		container.NewGridWithRows(2,
			c.Gui.label,
			container.NewGridWithColumns(2,
				c.Gui.btnPlus, c.Gui.btnMinus,
			),
		),
		c.Gui.settingsBtn,
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

func (c *CounterSvc) updateCount(mul int) {
	c.Data.count = c.Data.count + mul*c.Data.step
}

func (u *CounterSvc) updateUI() {
	// TODO
}
