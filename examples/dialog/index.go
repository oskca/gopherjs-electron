package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
)

func openFileDialog() {
	log.Println("openDilog called")
	d := electron.GetDialogModule()
	files := d.ShowOpenDialogEx(electron.DialogOptionOpen{
		Title:       "Renderer Opener",
		DefaultPath: "~/tmp",
		Properties: []string{
			electron.DialogPropOpenFile,
			electron.DialogPropShowHiddenFiles,
			electron.DialogPropMultiSelections,
		},
		Filters: []electron.FileFilterEx{
			{
				Name:       "Go Files",
				Extensions: []string{"go", "js"},
			},
		},
	})
	log.Println("selected files:", files)
}

func main() {
	electron.UseRemote()
	js.Global.Set("openFileDialog", openFileDialog)
	js.Global.Set("showMessage", func() {
		d := electron.GetDialogModule()
		d.ShowMessageEx("hello", "world")
	})
}
