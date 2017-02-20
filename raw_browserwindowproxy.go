package electron

import "github.com/gopherjs/gopherjs/js"

// BrowserWindowProxy version@1.4.15
//
// Manipulate the child browser window
type BrowserWindowProxy struct {
	*js.Object
	// A Boolean that is set to true after the child window gets closed.
	Closed bool `js:"closed"`
	// Removes focus from the child window.
	Blur func() `js:"blur"`
	// Forcefully closes the child window without calling its unload event.
	Close func() `js:"close"`
	// Evaluates the code in the child window.
	Eval func(Code string) `js:"eval"`
	// Focuses the child window (brings the window to front).
	Focus func() `js:"focus"`
	// Invokes the print dialog on the child window.
	Print func() `js:"print"`
	// Sends a message to the child window with the specified origin or * for no origin preference. In addition to these methods, the child window implements window.opener object with no properties and a single method.
	PostMessage func(Message string, TargetOrigin string) `js:"postMessage"`
}

func WrapBrowserWindowProxy(o *js.Object) *BrowserWindowProxy {
	return &BrowserWindowProxy{
		Object: o,
	}
}
