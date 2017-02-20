package electron

import "github.com/gopherjs/gopherjs/js"

// NetModule version@1.4.15
//
// Issue HTTP/HTTPS requests using Chromium's native networking library
type NetModule struct {
	*js.Object
	// Creates a ClientRequest instance using the provided options which are directly forwarded to the ClientRequest constructor. The net.request method would be used to issue both secure and insecure HTTP requests according to the specified protocol scheme in the options object.
	Request func(Options *NetModuleRequestOptions) (Obj *ClientRequest) `js:"request"`
}

func GetNetModule() *NetModule {
	o := Get("net")
	return &NetModule{
		Object: o,
	}
}

type NetModuleRequestOptions struct {
	*js.Object
}
