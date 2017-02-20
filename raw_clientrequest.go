package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	EvtClientRequestResponse = "response"
	// Emitted when an authenticating proxy is asking for user credentials. The callback function is expected to be called back with user credentials: Providing empty credentials will cancel the request and report an authentication error on the response object:
	EvtClientRequestLogin = "login"
	// Emitted just after the last chunk of the request's data has been written into the request object.
	EvtClientRequestFinish = "finish"
	// Emitted when the request is aborted. The abort event will not be fired if the request is already closed.
	EvtClientRequestAbort = "abort"
	// Emitted when the net module fails to issue a network request. Typically when the request object emits an error event, a close event will subsequently follow and no response object will be provided.
	EvtClientRequestError = "error"
	// Emitted as the last event in the HTTP request-response transaction. The close event indicates that no more events will be emitted on either the request or response objects.
	EvtClientRequestClose = "close"
)

// ClientRequest version@1.4.15
//
// Make HTTP/HTTPS requests.
type ClientRequest struct {
	*events.Emitter
	// A Boolean specifying whether the request will use HTTP chunked transfer encoding or not. Defaults to false. The property is readable and writable, however it can be set only before the first write operation as the HTTP headers are not yet put on the wire. Trying to set the chunkedEncoding property after the first write will throw an error. Using chunked encoding is strongly recommended if you need to send a large request body as data will be streamed in small chunks instead of being internally buffered inside Electron process memory.
	ChunkedEncoding bool `js:"chunkedEncoding"`
	// Adds an extra HTTP header. The header name will issued as it is without lowercasing. It can be called only before first write. Calling this method after the first write will throw an error.
	SetHeader func(Name string, Value string) `js:"setHeader"`
	// Returns String - The value of a previously set extra header name.
	GetHeader func(Name string) `js:"getHeader"`
	// Removes a previously set extra header name. This method can be called only before first write. Trying to call it after the first write will throw an error.
	RemoveHeader func(Name string) `js:"removeHeader"`
	// callback is essentially a dummy function introduced in the purpose of keeping similarity with the Node.js API. It is called asynchronously in the next tick after chunk content have been delivered to the Chromium networking layer. Contrary to the Node.js implementation, it is not guaranteed that chunk content have been flushed on the wire before callback is called. Adds a chunk of data to the request body. The first write operation may cause the request headers to be issued on the wire. After the first write operation, it is not allowed to add or remove a custom header.
	Write func(Chunk string, Encoding string, Callback ClientRequestWriteCallback) `js:"write"`
	// Sends the last chunk of the request data. Subsequent write or end operations will not be allowed. The finish event is emitted just after the end operation.
	End func(Chunk string, Encoding string, Callback ClientRequestEndCallback) `js:"end"`
	// Cancels an ongoing HTTP transaction. If the request has already emitted the close event, the abort operation will have no effect. Otherwise an ongoing event will emit abort and close events. Additionally, if there is an ongoing response object,it will emit the aborted event.
	Abort func() `js:"abort"`
}

func WrapClientRequest(o *js.Object) *ClientRequest {
	return &ClientRequest{
		Emitter: events.New(o),
	}
}

func NewClientRequest(Options *ClientRequestClientRequestOptions) *ClientRequest {
	o := electron.Get("ClientRequest")
	ret := o.New(Options)
	return WrapClientRequest(ret)
}

type ClientRequestWriteCallback func()
type ClientRequestEndCallback func()
type ClientRequestClientRequestOptions struct {
	*js.Object
}
