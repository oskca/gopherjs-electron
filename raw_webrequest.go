package electron

import "github.com/gopherjs/gopherjs/js"

// WebRequest version@1.4.15
//
// Intercept and modify the contents of a request at various stages of its lifetime.
type WebRequest struct {
	*js.Object
	// The listener will be called with listener(details, callback) when a request is about to occur. The uploadData is an array of UploadData objects. The callback has to be called with an response object.
	OnBeforeRequest func(Filter *WebRequestOnBeforeRequestFilter, Listener WebRequestOnBeforeRequestListener) `js:"onBeforeRequest"`
	// The listener will be called with listener(details, callback) before sending an HTTP request, once the request headers are available. This may occur after a TCP connection is made to the server, but before any http data is sent. The callback has to be called with an response object.
	OnBeforeSendHeaders func(Filter *WebRequestOnBeforeSendHeadersFilter, Listener WebRequestOnBeforeSendHeadersListener) `js:"onBeforeSendHeaders"`
	// The listener will be called with listener(details) just before a request is going to be sent to the server, modifications of previous onBeforeSendHeaders response are visible by the time this listener is fired.
	OnSendHeaders func(Filter *WebRequestOnSendHeadersFilter, Listener WebRequestOnSendHeadersListener) `js:"onSendHeaders"`
	// The listener will be called with listener(details, callback) when HTTP response headers of a request have been received. The callback has to be called with an response object.
	OnHeadersReceived func(Filter *WebRequestOnHeadersReceivedFilter, Listener WebRequestOnHeadersReceivedListener) `js:"onHeadersReceived"`
	// The listener will be called with listener(details) when first byte of the response body is received. For HTTP requests, this means that the status line and response headers are available.
	OnResponseStarted func(Filter *WebRequestOnResponseStartedFilter, Listener WebRequestOnResponseStartedListener) `js:"onResponseStarted"`
	// The listener will be called with listener(details) when a server initiated redirect is about to occur.
	OnBeforeRedirect func(Filter *WebRequestOnBeforeRedirectFilter, Listener WebRequestOnBeforeRedirectListener) `js:"onBeforeRedirect"`
	// The listener will be called with listener(details) when a request is completed.
	OnCompleted func(Filter *WebRequestOnCompletedFilter, Listener WebRequestOnCompletedListener) `js:"onCompleted"`
	// The listener will be called with listener(details) when an error occurs.
	OnErrorOccurred func(Filter *WebRequestOnErrorOccurredFilter, Listener WebRequestOnErrorOccurredListener) `js:"onErrorOccurred"`
}

func WrapWebRequest(o *js.Object) *WebRequest {
	return &WebRequest{
		Object: o,
	}
}

type WebRequestOnBeforeRequestFilter struct {
	*js.Object
}

type WebRequestOnBeforeRequestListener func(Details *WebRequestListenerDetails, Callback WebRequestListenerCallback)
type WebRequestListenerDetails struct {
	*js.Object
	Id           int64      `js:"id"`
	URL          string     `js:"url"`
	Method       string     `js:"method"`
	ResourceType string     `js:"resourceType"`
	Timestamp    float64    `js:"timestamp"`
	UploadData   *js.Object `js:"uploadData"`
}

type WebRequestListenerCallback func(Response *WebRequestCallbackResponse)
type WebRequestCallbackResponse struct {
	*js.Object
	Cancel bool `js:"cancel"`
	// The original request is prevented from being sent or completed and is instead redirected to the given URL.
	RedirectURL string `js:"redirectURL"`
}

type WebRequestOnSendHeadersListener func(Details *WebRequestListenerDetails2)
type WebRequestListenerDetails2 struct {
	*js.Object
	Id             int64                            `js:"id"`
	URL            string                           `js:"url"`
	Method         string                           `js:"method"`
	ResourceType   string                           `js:"resourceType"`
	Timestamp      float64                          `js:"timestamp"`
	RequestHeaders *WebRequestDetailsRequestHeaders `js:"requestHeaders"`
}

type WebRequestDetailsRequestHeaders struct {
	*js.Object
}

type WebRequestOnHeadersReceivedFilter struct {
	*js.Object
}

type WebRequestOnBeforeRedirectFilter struct {
	*js.Object
}

type WebRequestOnErrorOccurredListener func(Details *WebRequestListenerDetails3)
type WebRequestListenerDetails3 struct {
	*js.Object
	Id           int64   `js:"id"`
	URL          string  `js:"url"`
	Method       string  `js:"method"`
	ResourceType string  `js:"resourceType"`
	Timestamp    float64 `js:"timestamp"`
	FromCache    bool    `js:"fromCache"`
	// The error description.
	Error string `js:"error"`
}

type WebRequestOnResponseStartedFilter struct {
	*js.Object
}

type WebRequestOnBeforeRedirectListener func(Details *WebRequestListenerDetails4)
type WebRequestListenerDetails4 struct {
	*js.Object
	Id           string  `js:"id"`
	URL          string  `js:"url"`
	Method       string  `js:"method"`
	ResourceType string  `js:"resourceType"`
	Timestamp    float64 `js:"timestamp"`
	RedirectURL  string  `js:"redirectURL"`
	StatusCode   int64   `js:"statusCode"`
	// The server IP address that the request was actually sent to.
	Ip              string                            `js:"ip"`
	FromCache       bool                              `js:"fromCache"`
	ResponseHeaders *WebRequestDetailsResponseHeaders `js:"responseHeaders"`
}

type WebRequestDetailsResponseHeaders struct {
	*js.Object
}

type WebRequestOnCompletedFilter struct {
	*js.Object
}

type WebRequestOnResponseStartedListener func(Details *WebRequestListenerDetails5)
type WebRequestListenerDetails5 struct {
	*js.Object
	Id              int64                              `js:"id"`
	URL             string                             `js:"url"`
	Method          string                             `js:"method"`
	ResourceType    string                             `js:"resourceType"`
	Timestamp       float64                            `js:"timestamp"`
	ResponseHeaders *WebRequestDetailsResponseHeaders2 `js:"responseHeaders"`
	// Indicates whether the response was fetched from disk cache.
	FromCache  bool   `js:"fromCache"`
	StatusCode int64  `js:"statusCode"`
	StatusLine string `js:"statusLine"`
}

type WebRequestDetailsResponseHeaders2 struct {
	*js.Object
}

type WebRequestOnCompletedListener func(Details *WebRequestListenerDetails6)
type WebRequestListenerDetails6 struct {
	*js.Object
	Id              int64                              `js:"id"`
	URL             string                             `js:"url"`
	Method          string                             `js:"method"`
	ResourceType    string                             `js:"resourceType"`
	Timestamp       float64                            `js:"timestamp"`
	ResponseHeaders *WebRequestDetailsResponseHeaders3 `js:"responseHeaders"`
	FromCache       bool                               `js:"fromCache"`
	StatusCode      int64                              `js:"statusCode"`
	StatusLine      string                             `js:"statusLine"`
}

type WebRequestDetailsResponseHeaders3 struct {
	*js.Object
}

type WebRequestOnErrorOccurredFilter struct {
	*js.Object
}

type WebRequestOnBeforeSendHeadersFilter struct {
	*js.Object
}

type WebRequestOnBeforeSendHeadersListener func()
type WebRequestOnSendHeadersFilter struct {
	*js.Object
}

type WebRequestOnHeadersReceivedListener func()
