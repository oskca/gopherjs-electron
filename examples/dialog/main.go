package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

func main() {
	app := electron.GetApp()
	app.On(electron.EvtAppReady, func(args ...*js.Object) {
		opt := electron.NewBrowserWindowOption()
		bw := electron.NewBrowserWindow(opt)
		bw.LoadURL("file://"+nodejs.DirName()+"/index.html", nil)
		d := electron.GetDialogModule()
		d.ShowErrorBox("dialog from main process", "can you see this?")
		files := d.ShowOpenDialogEx(electron.DialogOptionOpen{
			Title:       "opener dialog",
			DefaultPath: "/tmp",
			Properties: []string{electron.DialogPropOpenFile,
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
	})
}
