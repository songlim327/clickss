package gui

import "fyne.io/fyne/v2"

// fullScreenHandler Enable/Disable Fullscreen
func fullScreenHandler(w fyne.Window) {
	w.SetFullScreen(!w.FullScreen())
}

// quitHandler quit Application
func quitHandler(app fyne.App) {
	app.Quit()
}
