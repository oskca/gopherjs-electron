package electron

import "github.com/gopherjs/gopherjs/js"

var (
	require   = js.Global.Get("require")
	electron  = require.Invoke("electron")
	remote    = electron.Get("remote")
	useRemote = false
)

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
