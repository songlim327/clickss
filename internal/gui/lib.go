package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// centralize centralize a fyne object
func centralize(o fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewCenterLayout(), o)
}

// createLabel create fyne widget.Label with specific fields
func createLabel(text string, bold bool) *widget.Label {
	l := widget.NewLabel(text)
	l.TextStyle.Bold = bold
	l.Alignment = fyne.TextAlignCenter
	return l
}

// createText create canvas.Text with specific fields
func createText(text string, size float32, bold bool) *canvas.Text {
	c := canvas.NewText(text, color.Black)
	c.TextStyle.Bold = bold
	c.TextSize = size
	c.Alignment = fyne.TextAlignCenter
	return c
}

// pad create empty label for padding
func pad() fyne.CanvasObject {
	return container.New(layout.NewHBoxLayout(), widget.NewLabel(""))
}
