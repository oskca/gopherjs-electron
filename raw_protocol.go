package electron

import "github.com/gopherjs/gopherjs/js"

// ProtocolModule version@1.4.15
//
// Register a custom protocol and intercept existing protocol requests.
type ProtocolModule struct {
	*js.Object
	// A standard scheme adheres to what RFC 3986 calls generic URI syntax. For example http and https are standard schemes, while file is not. Registering a scheme as standard, will allow relative and absolute resources to be resolved correctly when served. Otherwise the scheme will behave like the file protocol, but without the ability to resolve relative URLs. For example when you load following page with custom protocol without registering it as standard scheme, the image will not be loaded because non-standard schemes can not recognize relative URLs: Registering a scheme as standard will allow access to files through the FileSystem API. Otherwise the renderer will throw a security error for the scheme. By default web storage apis (localStorage, sessionStorage, webSQL, indexedDB, cookies) are disabled for non standard schemes. So in general if you want to register a custom protocol to replace the http protocol, you have to register it as a standard scheme: Note: This method can only be used before the ready event of the app module gets emitted.
	RegisterStandardSchemes      func(Schemes *js.Object, Options *ProtocolModuleRegisterStandardSchemesOptions) `js:"registerStandardSchemes"`
	RegisterServiceWorkerSchemes func(Schemes *js.Object)                                                        `js:"registerServiceWorkerSchemes"`
	// Registers a protocol of scheme that will send the file as a response. The handler will be called with handler(request, callback) when a request is going to be created with scheme. completion will be called with completion(null) when scheme is successfully registered or completion(error) when failed. To handle the request, the callback should be called with either the file's path or an object that has a path property, e.g. callback(filePath) or callback({path: filePath}). When callback is called with nothing, a number, or an object that has an error property, the request will fail with the error number you specified. For the available error numbers you can use, please see the net error list. By default the scheme is treated like http:, which is parsed differently than protocols that follow the "generic URI syntax" like file:, so you probably want to call protocol.registerStandardSchemes to have your scheme treated as a standard scheme.
	RegisterFileProtocol func(Scheme string, Handler ProtocolModuleRegisterFileProtocolHandler, Completion ProtocolModuleRegisterFileProtocolCompletion) `js:"registerFileProtocol"`
	// Registers a protocol of scheme that will send a Buffer as a response. The usage is the same with registerFileProtocol, except that the callback should be called with either a Buffer object or an object that has the data, mimeType, and charset properties. Example:
	RegisterBufferProtocol func(Scheme string, Handler ProtocolModuleRegisterBufferProtocolHandler, Completion ProtocolModuleRegisterBufferProtocolCompletion) `js:"registerBufferProtocol"`
	// Registers a protocol of scheme that will send a String as a response. The usage is the same with registerFileProtocol, except that the callback should be called with either a String or an object that has the data, mimeType, and charset properties.
	RegisterStringProtocol func(Scheme string, Handler ProtocolModuleRegisterStringProtocolHandler, Completion ProtocolModuleRegisterStringProtocolCompletion) `js:"registerStringProtocol"`
	// Registers a protocol of scheme that will send an HTTP request as a response. The usage is the same with registerFileProtocol, except that the callback should be called with a redirectRequest object that has the url, method, referrer, uploadData and session properties. By default the HTTP request will reuse the current session. If you want the request to have a different session you should set session to null. For POST requests the uploadData object must be provided.
	RegisterHttpProtocol func(Scheme string, Handler ProtocolModuleRegisterHttpProtocolHandler, Completion ProtocolModuleRegisterHttpProtocolCompletion) `js:"registerHttpProtocol"`
	// Unregisters the custom protocol of scheme.
	UnregisterProtocol func(Scheme string, Completion ProtocolModuleUnregisterProtocolCompletion) `js:"unregisterProtocol"`
	// The callback will be called with a boolean that indicates whether there is already a handler for scheme.
	IsProtocolHandled func(Scheme string, Callback ProtocolModuleIsProtocolHandledCallback) `js:"isProtocolHandled"`
	// Intercepts scheme protocol and uses handler as the protocol's new handler which sends a file as a response.
	InterceptFileProtocol func(Scheme string, Handler ProtocolModuleInterceptFileProtocolHandler, Completion ProtocolModuleInterceptFileProtocolCompletion) `js:"interceptFileProtocol"`
	// Intercepts scheme protocol and uses handler as the protocol's new handler which sends a String as a response.
	InterceptStringProtocol func(Scheme string, Handler ProtocolModuleInterceptStringProtocolHandler, Completion ProtocolModuleInterceptStringProtocolCompletion) `js:"interceptStringProtocol"`
	// Intercepts scheme protocol and uses handler as the protocol's new handler which sends a Buffer as a response.
	InterceptBufferProtocol func(Scheme string, Handler ProtocolModuleInterceptBufferProtocolHandler, Completion ProtocolModuleInterceptBufferProtocolCompletion) `js:"interceptBufferProtocol"`
	// Intercepts scheme protocol and uses handler as the protocol's new handler which sends a new HTTP request as a response.
	InterceptHttpProtocol func(Scheme string, Handler ProtocolModuleInterceptHttpProtocolHandler, Completion ProtocolModuleInterceptHttpProtocolCompletion) `js:"interceptHttpProtocol"`
	// Remove the interceptor installed for scheme and restore its original handler.
	UninterceptProtocol func(Scheme string, Completion ProtocolModuleUninterceptProtocolCompletion) `js:"uninterceptProtocol"`
}

func GetProtocolModule() *ProtocolModule {
	o := Get("protocol")
	return &ProtocolModule{
		Object: o,
	}
}

type ProtocolModuleInterceptHttpProtocolCompletion func(Error *js.Object)
type ProtocolModuleRegisterBufferProtocolHandler func(Request *ProtocolModuleHandlerRequest, Callback ProtocolModuleHandlerCallback)
type ProtocolModuleHandlerRequest struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback func(Buffer *js.Object)
type ProtocolModuleRegisterStringProtocolHandler func(Request *ProtocolModuleHandlerRequest2, Callback ProtocolModuleHandlerCallback2)
type ProtocolModuleHandlerRequest2 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback2 func(Data string)
type ProtocolModuleInterceptFileProtocolHandler func(Request *ProtocolModuleHandlerRequest3, Callback ProtocolModuleHandlerCallback3)
type ProtocolModuleHandlerRequest3 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback3 func(FilePath string)
type ProtocolModuleInterceptStringProtocolCompletion func(Error *js.Object)
type ProtocolModuleInterceptBufferProtocolHandler func(Request *ProtocolModuleHandlerRequest4, Callback ProtocolModuleHandlerCallback4)
type ProtocolModuleHandlerRequest4 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback4 func(Buffer *js.Object)
type ProtocolModuleRegisterFileProtocolCompletion func(Error *js.Object)
type ProtocolModuleInterceptFileProtocolCompletion func(Error *js.Object)
type ProtocolModuleInterceptBufferProtocolCompletion func(Error *js.Object)
type ProtocolModuleUninterceptProtocolCompletion func(Error *js.Object)
type ProtocolModuleRegisterFileProtocolHandler func(Request *ProtocolModuleHandlerRequest5, Callback ProtocolModuleHandlerCallback5)
type ProtocolModuleHandlerCallback5 func(FilePath string)
type ProtocolModuleHandlerRequest5 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleRegisterHttpProtocolHandler func(Request *ProtocolModuleHandlerRequest6, Callback ProtocolModuleHandlerCallback6)
type ProtocolModuleHandlerRequest6 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback6 func(RedirectRequest *ProtocolModuleCallbackRedirectRequest)
type ProtocolModuleCallbackRedirectRequest struct {
	*js.Object
	URL        string                                   `js:"url"`
	Method     string                                   `js:"method"`
	Session    *ProtocolModuleRedirectRequestSession    `js:"session"`
	UploadData *ProtocolModuleRedirectRequestUploadData `js:"uploadData"`
}

type ProtocolModuleRedirectRequestSession struct {
	*js.Object
}

type ProtocolModuleRedirectRequestUploadData struct {
	*js.Object
	// MIME type of the content.
	ContentType string `js:"contentType"`
	// Content to be sent.
	Data string `js:"data"`
}

type ProtocolModuleUnregisterProtocolCompletion func(Error *js.Object)
type ProtocolModuleInterceptHttpProtocolHandler func(Request *ProtocolModuleHandlerRequest7, Callback ProtocolModuleHandlerCallback7)
type ProtocolModuleHandlerRequest7 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleHandlerCallback7 func(RedirectRequest *ProtocolModuleCallbackRedirectRequest2)
type ProtocolModuleCallbackRedirectRequest2 struct {
	*js.Object
	URL        string                                    `js:"url"`
	Method     string                                    `js:"method"`
	Session    *ProtocolModuleRedirectRequestSession2    `js:"session"`
	UploadData *ProtocolModuleRedirectRequestUploadData2 `js:"uploadData"`
}

type ProtocolModuleRedirectRequestSession2 struct {
	*js.Object
}

type ProtocolModuleRedirectRequestUploadData2 struct {
	*js.Object
	// MIME type of the content.
	ContentType string `js:"contentType"`
	// Content to be sent.
	Data string `js:"data"`
}

type ProtocolModuleInterceptStringProtocolHandler func(Request *ProtocolModuleHandlerRequest8, Callback ProtocolModuleHandlerCallback8)
type ProtocolModuleHandlerCallback8 func(Data string)
type ProtocolModuleHandlerRequest8 struct {
	*js.Object
	URL        string     `js:"url"`
	Referrer   string     `js:"referrer"`
	Method     string     `js:"method"`
	UploadData *js.Object `js:"uploadData"`
}

type ProtocolModuleRegisterStandardSchemesOptions struct {
	*js.Object
	// to register the scheme as secure. Default .
	Secure bool `js:"secure"`
}

type ProtocolModuleRegisterBufferProtocolCompletion func(Error *js.Object)
type ProtocolModuleRegisterStringProtocolCompletion func(Error *js.Object)
type ProtocolModuleRegisterHttpProtocolCompletion func(Error *js.Object)
type ProtocolModuleIsProtocolHandledCallback func(Error *js.Object)
