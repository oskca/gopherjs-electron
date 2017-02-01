package electron

import "github.com/gopherjs/gopherjs/js"

var (
	require  = js.Global.Get("require")
	electron = require.Invoke("electron")
)

func Get(name string) *js.Object {
	return electron.Get(name)
}
