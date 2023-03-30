//go:generate goversioninfo -platform-specific=true resources/versioninfo.json
package main

import (
	"clickss/internal/gui"
)

func main() {
	gui.CreateApp()
}
