package electron

import (
	"github.com/gopherjs/gopherjs/js"
	raw "github.com/oskca/gopherjs-electron/rawapi"
)

func NewBrowserWindowOption() *raw.BrowserWindowBrowserWindowOptions {
	opt := &raw.BrowserWindowBrowserWindowOptions{
		Object: js.Global.Get("Object").New(),
	}
	opt.Width = 800
	opt.Height = 600
	return opt
}
