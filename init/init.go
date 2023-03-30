package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
)

const r = "images"

func main() {
	var files []string
	p := r + "\\image.go"

	createResourceFile(p)

	err := filepath.Walk(r, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".png" {
			files = append(files, path)
		}
		return nil
	})
	check(err)

	rf, err := os.OpenFile(p, os.O_RDWR|os.O_APPEND, 0660)
	check(err)

	for _, f := range files {
		b1 := filepath.Base(f)
		r2 := strings.ReplaceAll(b1, ".png", "")
		c1 := strings.Title(r2)
		r1 := strings.ReplaceAll(c1, "-", "")

		res, err := fyne.LoadResourceFromPath(f)
		check(err)
		rf.WriteString(fmt.Sprintf("%s = %#v\n", r1, res))
	}
	rf.WriteString(")")
	defer rf.Close()
}

// check logs if error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// createResourceFile creates resource file containing images byte array
func createResourceFile(p string) {
	if fileExist(p) {
		os.Remove(p)
	}

	f, err := os.Create(p)
	check(err)
	f.WriteString("// **** THIS FILE IS AUTO-GENERATED **** //\n\npackage images\n\nimport \"fyne.io/fyne/v2\"\n\nvar (\n")
	defer f.Close()
}

// fileExist check if a file exist
func fileExist(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}
