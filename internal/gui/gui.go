package gui

import (
	"clickss/images"
	"clickss/internal/constants"
	"clickss/internal/core"
	"clickss/internal/gui/components"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Variables
var (
	fApp fyne.App
	fWin fyne.Window
)

// CreateApp create main application window
func CreateApp() {
	fApp = app.NewWithID(constants.AppName)

	// load custom theme
	fApp.Settings().SetTheme(&CustomTheme{})

	fWin = fApp.NewWindow(constants.AppName)

	fWin.CenterOnScreen()
	fWin.SetMaster()
	fWin.SetIcon(images.Logo)

	fWin.Resize(fyne.NewSize(400, 600))

	// set main menu for main application window
	fWin.SetMainMenu(createMainMenu())

	// create application content
	createMain()

	// add shortcut keys
	captureKeys()

	fWin.ShowAndRun()
}

// captureKeys capture key combination throughout the whole application
func captureKeys() {
	fWin.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyF11 {
			fWin.SetFullScreen(!fWin.FullScreen())
		}
	})
}

// createMainMenu render menu options
func createMainMenu() *fyne.MainMenu {
	l := []*fyne.Menu{
		fyne.NewMenu("View", fyne.NewMenuItem("Fullscreen	F11", func() { fullScreenHandler(fWin) })),
		fyne.NewMenu("Help", fyne.NewMenuItem("About", func() { createAbout() })),
	}
	return fyne.NewMainMenu(l...)
}

// createAbout render about UI
func createAbout() {
	w := fApp.NewWindow("About")
	w.CenterOnScreen()
	w.SetIcon(images.Logo)
	w.Resize(fyne.NewSize(360, 270))
	w.SetFixedSize(true)

	logo := canvas.NewImageFromResource(images.Logo)
	logo.SetMinSize(fyne.NewSize(64, 64))
	t1 := createLabel(constants.AppName+" v"+constants.AppVer, true)
	t2 := createText(constants.AppDesc, 14, false)
	t3 := createText("Author: ", 12, false)
	u, _ := url.Parse("mailto:" + constants.Author)
	h := widget.NewHyperlink(constants.Author, u)
	h.Alignment = fyne.TextAlignCenter

	l := container.New(layout.NewVBoxLayout(), centralize(logo), t1, t2, t3, h)
	c := widget.NewCard("", "", l)
	content := container.New(layout.NewCenterLayout(), c)

	w.SetContent(content)
	w.Show()
}

// createMain render main window
func createMain() {
	// render main screen logo and application name
	i1 := canvas.NewImageFromResource(images.Logo)
	i1.SetMinSize(fyne.NewSize(128, 128))
	i1.FillMode = canvas.ImageFillContain
	t1 := createText(constants.AppName, 28, true)
	l1 := container.New(layout.NewVBoxLayout(), i1, t1)

	// frequency numerical input
	freqEntry := components.NewNumericalEntry()
	freqEntry.SetPlaceHolder("click frequency")
	freqEntry.SetText("100")

	// left or right click radio group
	radioOpts := []string{"Left", "Right"}
	buttonRadio := widget.NewRadioGroup(radioOpts, func(value string) {})
	buttonRadio.SetSelected(radioOpts[0])

	// shortcut key to start/stop
	scKeyOpts := []string{"F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F12"}
	scKeyCombo := widget.NewSelect(scKeyOpts, func(value string) {})
	scKeyCombo.SetSelectedIndex(7)

	// render app options
	action := ""
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Clicks per second", Widget: freqEntry},
			{Text: "button", Widget: buttonRadio},
			{Text: "Start/Stop shortcut", Widget: scKeyCombo},
		},
		OnSubmit: func() {
			duration, err := strconv.Atoi(freqEntry.Text)
			if err != nil {
				log.Fatal("[Error] Duration string parse: ", err)
			}

			if action != "start" {
				core.StartHandler(strings.ToLower(buttonRadio.Selected), duration, strings.ToLower(scKeyCombo.Selected))
				action = "start"
			}
			fyne.CurrentApp().SendNotification(fyne.NewNotification(constants.AppName, fmt.Sprintf("Click %s to start", scKeyCombo.Selected)))
		},
		OnCancel: func() {
			if action != "stop" {
				core.StopHandler()
				action = "stop"
			}
		},
		SubmitText: "Start",
		CancelText: "Stop",
	}

	l2 := container.New(layout.NewBorderLayout(nil, pad(), nil, pad()), form)
	c2 := container.New(layout.NewGridLayout(1), centralize(l1), l2)
	fWin.SetContent(c2)
}
