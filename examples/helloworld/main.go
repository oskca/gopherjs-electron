package main

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/browserwindow"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

func main() {
	app := electron.GetApp()
	app.On(electron.EvtReady, func(args ...*js.Object) {
		opt := browserwindow.NewOption()
		bw := browserwindow.New(opt)
		bw.LoadURL("file://" + nodejs.DirName() + "/index.html")
	})
}
