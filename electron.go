// Package electron implements an API translator and some convenience functions for creating
// Electron APP using Gopherjs.
//
// Usage
//
// go get github.com/oskca/gopherjs-electron
// go generate # calling the API translator
//
// API Translator
//
// The translator reads electron api json file and output the corresponding gopherjs struct and
// functions into files with raw_ prefix.
//
// Easy binding functions
//
// Hand written binding functions/struct lives in ex_*.go files and have Ex postfix.
package electron

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	require   = js.Global.Get("require")
	electron  = require.Invoke("electron")
	remote    = electron.Get("remote")
	useRemote = false
)

//go:generate -command json2rawApi go run json2rawApi/main.go json2rawApi/types.go json2rawApi/templates.go
//go:generate json2rawApi -c -o . json2rawApi/electron-api-1.4.15.json

func GetApp() *AppModule {
	return GetAppModule()
}

// Get returns a electron or `electron.remote` module
func Get(name string) *js.Object {
	if useRemote {
		return remote.Get(name)
	}
	return electron.Get(name)
}

// UseRemote switch `electron.Get` to get module through `electron.remote`
// this functon must be call when use electron.remote in render process
func UseRemote() {
	useRemote = true
}
