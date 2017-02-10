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
