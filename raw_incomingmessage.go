package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// The data event is the usual method of transferring response data into applicative code.
	EvtIncomingMessageData = "data"
	// Indicates that response body has ended.
	EvtIncomingMessageEnd = "end"
	// Emitted when a request has been canceled during an ongoing HTTP transaction.
	EvtIncomingMessageAborted = "aborted"
	// error Error - Typically holds an error string identifying failure root cause. Emitted when an error was encountered while streaming response data events. For instance, if the server closes the underlying while the response is still streaming, an error event will be emitted on the response object and a close event will subsequently follow on the request object.
	EvtIncomingMessageError = "error"
)

// IncomingMessage version@1.4.15
//
// Handle responses to HTTP/HTTPS requests.
type IncomingMessage struct {
	*events.Emitter
	// An Integer indicating the HTTP response status code.
	StatusCode int64 `js:"statusCode"`
	// A String representing the HTTP status message.
	StatusMessage string `js:"statusMessage"`
	// An Object representing the response HTTP headers. The headers object is formatted as follows:
	Headers *IncomingMessageIncomingMessageHeaders `js:"headers"`
	// A String indicating the HTTP protocol version number. Typical values are '1.0' or '1.1'. Additionally httpVersionMajor and httpVersionMinor are two Integer-valued readable properties that return respectively the HTTP major and minor version numbers.
	HttpVersion string `js:"httpVersion"`
	// An Integer indicating the HTTP protocol major version number.
	HttpVersionMajor int64 `js:"httpVersionMajor"`
	// An Integer indicating the HTTP protocol minor version number.
	HttpVersionMinor int64 `js:"httpVersionMinor"`
}

func WrapIncomingMessage(o *js.Object) *IncomingMessage {
	return &IncomingMessage{
		Emitter: events.New(o),
	}
}

type IncomingMessageIncomingMessageHeaders struct {
	*js.Object
}
