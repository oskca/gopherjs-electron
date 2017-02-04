package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/browserwindow"
	"github.com/oskca/gopherjs-electron/api/main/dialog"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

func main() {
	app := electron.GetApp()
	app.On(electron.EvtReady, func(args ...*js.Object) {
		opt := browserwindow.NewOption()
		bw := browserwindow.New(opt)
		bw.LoadURL("file://" + nodejs.DirName() + "/index.html")
		d := dialog.Get()
		d.ShowErrorBox("dialog from main process", "can you see this?")
		files := d.ShowOpenDialog(dialog.OptionOpen{
			Title:       "opener dialog",
			DefaultPath: "/tmp",
			Properties:  []string{dialog.PropOpenFile, dialog.PropShowHiddenFiles, dialog.PropMultiSelections},
			Filters: []dialog.FileFilter{
				dialog.FileFilter{
					Name:       "Go Files",
					Extensions: []string{"go", "js"},
				},
			},
		})
		log.Println("selected files:", files)
	})
}
