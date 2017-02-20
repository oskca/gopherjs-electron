package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when Electron is about to download item in webContents. Calling event.preventDefault() will cancel the download and item will not be available from next tick of the process.
	EvtSessionWillDownload = "will-download"
)

// Session version@1.4.15
//
// Get and set properties of a session.
type Session struct {
	*events.Emitter
	// A Cookies object for this session.
	Cookies *Cookies `js:"cookies"`
	// A WebRequest object for this session.
	WebRequest *WebRequest `js:"webRequest"`
	// A Protocol object (an instance of protocol module) for this session.
	Protocol *js.Object `js:"protocol"`
	// Returns the session's current cache size.
	GetCacheSize func(Callback SessionGetCacheSizeCallback) `js:"getCacheSize"`
	// Clears the session’s HTTP cache.
	ClearCache func(Callback SessionClearCacheCallback) `js:"clearCache"`
	// Clears the data of web storages.
	ClearStorageData func(Options *SessionClearStorageDataOptions, Callback SessionClearStorageDataCallback) `js:"clearStorageData"`
	// Writes any unwritten DOMStorage data to disk.
	FlushStorageData func() `js:"flushStorageData"`
	// Sets the proxy settings. When pacScript and proxyRules are provided together, the proxyRules option is ignored and pacScript configuration is applied. The proxyRules has to follow the rules below: For example: The proxyBypassRules is a comma separated list of rules described below:
	SetProxy func(Config *SessionSetProxyConfig, Callback SessionSetProxyCallback) `js:"setProxy"`
	// Resolves the proxy information for url. The callback will be called with callback(proxy) when the request is performed.
	ResolveProxy func(URL *js.Object, Callback SessionResolveProxyCallback) `js:"resolveProxy"`
	// Sets download saving directory. By default, the download directory will be the Downloads under the respective app folder.
	SetDownloadPath func(Path string) `js:"setDownloadPath"`
	// Emulates network with the given configuration for the session.
	EnableNetworkEmulation func(Options *SessionEnableNetworkEmulationOptions) `js:"enableNetworkEmulation"`
	// Disables any network emulation already active for the session. Resets to the original network configuration.
	DisableNetworkEmulation func() `js:"disableNetworkEmulation"`
	// Sets the certificate verify proc for session, the proc will be called with proc(hostname, certificate, callback) whenever a server certificate verification is requested. Calling callback(true) accepts the certificate, calling callback(false) rejects it. Calling setCertificateVerifyProc(null) will revert back to default certificate verify proc.
	SetCertificateVerifyProc func(Proc SessionSetCertificateVerifyProcProc) `js:"setCertificateVerifyProc"`
	// Sets the handler which can be used to respond to permission requests for the session. Calling callback(true) will allow the permission and callback(false) will reject it.
	SetPermissionRequestHandler func(Handler SessionSetPermissionRequestHandlerHandler) `js:"setPermissionRequestHandler"`
	// Clears the host resolver cache.
	ClearHostResolverCache func(Callback SessionClearHostResolverCacheCallback) `js:"clearHostResolverCache"`
	// Dynamically sets whether to always send credentials for HTTP NTLM or Negotiate authentication.
	AllowNTLMCredentialsForDomains func(Domains string) `js:"allowNTLMCredentialsForDomains"`
	// Overrides the userAgent and acceptLanguages for this session. The acceptLanguages must a comma separated ordered list of language codes, for example "en-US,fr,de,ko,zh-CN,ja". This doesn't affect existing WebContents, and each WebContents can use webContents.setUserAgent to override the session-wide user agent.
	SetUserAgent func(UserAgent string, AcceptLanguages string)                                `js:"setUserAgent"`
	GetUserAgent func() (Obj string)                                                           `js:"getUserAgent"`
	GetBlobData  func(Identifier string, Callback SessionGetBlobDataCallback) (Obj *js.Object) `js:"getBlobData"`
	// Allows resuming cancelled or interrupted downloads from previous Session. The API will generate a DownloadItem that can be accessed with the will-download event. The DownloadItem will not have any WebContents associated with it and the initial state will be interrupted. The download will start only when the resume API is called on the DownloadItem.
	CreateInterruptedDownload func(Options *SessionCreateInterruptedDownloadOptions) `js:"createInterruptedDownload"`
	// Clears the session’s HTTP authentication cache.
	ClearAuthCache func(Options *js.Object, Callback SessionClearAuthCacheCallback) `js:"clearAuthCache"`
}

func WrapSession(o *js.Object) *Session {
	return &Session{
		Emitter: events.New(o),
	}
}

type SessionClearHostResolverCacheCallback func()
type SessionResolveProxyCallback func(Proxy *SessionCallbackProxy)
type SessionCallbackProxy struct {
	*js.Object
}

type SessionSetCertificateVerifyProcProc func(Hostname string, Certificate *js.Object, Callback SessionProcCallback)
type SessionProcCallback func( // Determines if the certificate should be trusted
	IsTrusted bool)
type SessionCreateInterruptedDownloadOptions struct {
	*js.Object
	// Absolute path of the download.
	Path string `js:"path"`
	// Complete URL chain for the download.
	URLChain *js.Object `js:"urlChain"`
	MimeType string     `js:"mimeType"`
	// Start range for the download.
	Offset int64 `js:"offset"`
	// Total length of the download.
	Length int64 `js:"length"`
	// Last-Modified header value.
	LastModified string `js:"lastModified"`
	// ETag header value.
	ETag string `js:"eTag"`
	// Time when download was started in number of seconds since UNIX epoch.
	StartTime float64 `js:"startTime"`
}

type SessionClearCacheCallback func()
type SessionClearStorageDataCallback func()
type SessionSetProxyCallback func()
type SessionEnableNetworkEmulationOptions struct {
	*js.Object
	// Whether to emulate network outage. Defaults to false.
	Offline bool `js:"offline"`
	// RTT in ms. Defaults to 0 which will disable latency throttling.
	Latency float64 `js:"latency"`
	// Download rate in Bps. Defaults to 0 which will disable download throttling.
	DownloadThroughput float64 `js:"downloadThroughput"`
	// Upload rate in Bps. Defaults to 0 which will disable upload throttling.
	UploadThroughput float64 `js:"uploadThroughput"`
}

type SessionGetBlobDataCallback func( // Blob data.
					Result *js.Object)
type SessionGetCacheSizeCallback func( // Cache size used in bytes.
	Size int64)
type SessionClearStorageDataOptions struct {
	*js.Object
	// Should follow ’s representation .
	Origin string `js:"origin"`
	// The types of storages to clear, can contain: , , , , , , ,
	Storages *js.Object `js:"storages"`
	// The types of quotas to clear, can contain: , , .
	Quotas *js.Object `js:"quotas"`
}

type SessionSetProxyConfig struct {
	*js.Object
	// The URL associated with the PAC file.
	PacScript string `js:"pacScript"`
	// Rules indicating which proxies to use.
	ProxyRules string `js:"proxyRules"`
	// Rules indicating which URLs should bypass the proxy settings.
	ProxyBypassRules string `js:"proxyBypassRules"`
}

type SessionSetPermissionRequestHandlerHandler func( // requesting the permission.
	WebContents *SessionHandlerWebContents, // Enum of 'media', 'geolocation', 'notifications', 'midiSysex', 'pointerLock', 'fullscreen', 'openExternal'.
	Permission string, Callback SessionHandlerCallback)
type SessionHandlerWebContents struct {
	*js.Object
}

type SessionHandlerCallback func( // Allow or deny the permission
	PermissionGranted bool)
type SessionClearAuthCacheCallback func()
