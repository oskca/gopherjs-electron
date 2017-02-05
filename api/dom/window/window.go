package window

import "github.com/oskca/gopherjs-dom"
import "github.com/oskca/gopherjs-electron/api/renderer/browserwindowproxy"

// Open opens a url in renderer process
func Open(url string) *browserwindowproxy.Proxy {
	o := dom.Window().Call("open", url)
	return browserwindowproxy.Wrap(o)
}
