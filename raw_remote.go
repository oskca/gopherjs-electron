package electron

import "github.com/gopherjs/gopherjs/js"

// RemoteModule version@1.4.15
//
// Use main process modules from the renderer process.
type RemoteModule struct {
	*js.Object
	// The process object in the main process. This is the same as remote.getGlobal('process') but is cached.
	Process               string                               `js:"process"`
	Require               func(Module string) (Obj *js.Object) `js:"require"`
	GetCurrentWindow      func() (Obj *BrowserWindow)          `js:"getCurrentWindow"`
	GetCurrentWebContents func() (Obj *WebContents)            `js:"getCurrentWebContents"`
	GetGlobal             func(Name string) (Obj *js.Object)   `js:"getGlobal"`
}

func GetRemoteModule() *RemoteModule {
	o := Get("remote")
	return &RemoteModule{
		Object: o,
	}
}
