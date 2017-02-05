package browserwindowproxy

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-dom"
)

func Wrap(o *js.Object) *Proxy {
	return &Proxy{
		Win: dom.WrapWindow(o),
	}
}

// Proxy  Class: BrowserWindowProxy
// Manipulate the child browser window
type Proxy struct {
	*dom.Win
	// Process: Renderer

	// The BrowserWindowProxy object is returned from window.open and provides limited functionality with the child window.

	// Instance Methods
	// The BrowserWindowProxy object has the following instance methods:

	// win.blur()
	// Removes focus from the child window.

	// win.close()
	// Forcefully closes the child window without calling its unload event.

	// win.eval(code)
	// code String
	// Evaluates the code in the child window.
	Eval func(code string) `js:"eval"`

	// win.focus()
	// Focuses the child window (brings the window to front).

	// win.print()
	// Invokes the print dialog on the child window.
	Print func() `js:"print"`

	// win.postMessage(message, targetOrigin)
	// message String
	// targetOrigin String
	// Sends a message to the child window with the specified origin or * for no origin preference.
	PostMessage func(message, targetOrigin string) `js:"postMessage"`

	// In addition to these methods, the child window implements window.opener object with no properties and a single method.

	// Instance Properties
	// The BrowserWindowProxy object has the following instance properties:

	// win.closed
	// A Boolean that is set to true after the child window gets closed.
	Closed bool `js:"closed"`
}
