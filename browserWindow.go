package electron

import (
	"github.com/gopherjs/gopherjs/js"
)

func NewBrowserWindowOption() *BrowserWindowBrowserWindowOptions {
	opt := &BrowserWindowBrowserWindowOptions{
		Object: js.Global.Get("Object").New(),
	}
	opt.Width = 800
	opt.Height = 600
	return opt
}
