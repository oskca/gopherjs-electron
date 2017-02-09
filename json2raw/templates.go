package main

var electronTemplate = `package electron

import "github.com/gopherjs/gopherjs/js"

var (
	require   = js.Global.Get("require")
	electron  = require.Invoke("electron")
	remote    = electron.Get("remote")
	useRemote = false
)

// Get returns a electron or electron.remote module
func Get(name string) *js.Object {
	if useRemote {
		return remote.Get(name)
	}
	return electron.Get(name)
}

// UseRemote switch electron.Get to get module through electron.remote
// this functon must be call when use electron.remote in render process
func UseRemote() {
	useRemote = true
}`

var moduleConstructorTemplate = `
func GetModule() *Module {
	o := electron.Get("Module")
	return &Module{
		Object: o,
	}
}
`

var classWithEmitterConstructorTemplate = `
func WrapClassWithEmitter(o *js.Object) ClassWithEmitter {
	return &ClassWithEmitter{
		Emitter: events.New(o),
	}
}
`

var classConstructorTemplate = `
func WrapClass(o *js.Object) Class {
	return &Class{
		Object: o,
	}
}
`
