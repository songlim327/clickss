package core

import (
	"clickss/internal/constants"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	quitChan  chan bool         // click event go routine control channel
	shortChan chan bool         // shortcut keyboard event go routine control channel
	isActive  bool      = false // auto clicking state
)

// StartHandler start the clicking process
func StartHandler(button string, duration int, key string) {
	shortChan = make(chan bool)
	quitChan = make(chan bool)

	shortcut(button, duration, key)
}

// StopHandler stop the clicking process
func StopHandler() {
	log.Println("Disabled")
	if quitChan != nil && isActive {
		quitChan <- true
	}
	if shortChan != nil {
		shortChan <- true
	}
}

// register create a go routine which do the clicking
func register(button string, duration int) {
	go func(button string, duration int) {
		for {
			select {
			case <-quitChan:
				return
			default:
				// run the clicks
				robotgo.Click(button, false)
				time.Sleep(time.Millisecond * time.Duration(duration))
			}
		}
	}(button, duration)
}

func shortcut(button string, duration int, key string) {
	go func(button string, duration int, key string) {
		h := hook.Start()
		defer hook.End()
		for {
			select {
			case <-shortChan:
				return
			default:
				i := <-h
				if i.Kind > hook.KeyDown && i.Kind < hook.KeyUp {
					if i.Rawcode == hook.KeychartoRawcode(key) {
						isActive = !isActive
						if isActive {
							fyne.CurrentApp().SendNotification(fyne.NewNotification(constants.AppName, "Service activated"))
							register(button, duration)
						} else {
							quitChan <- true
						}
					}
				}
			}
		}
	}(button, duration, key)
}
