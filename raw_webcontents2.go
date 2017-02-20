package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when the navigation is done, i.e. the spinner of the tab has stopped spinning, and the onload event was dispatched.
	EvtWebContentsDidFinishLoad = "did-finish-load"
	// This event is like did-finish-load but emitted when the load failed or was cancelled, e.g. window.stop() is invoked. The full list of error codes and their meaning is available here. Note that redirect responses will emit errorCode -3; you may want to ignore that error explicitly.
	EvtWebContentsDidFailLoad = "did-fail-load"
	// Emitted when a frame has done navigation.
	EvtWebContentsDidFrameFinishLoad = "did-frame-finish-load"
	// Corresponds to the points in time when the spinner of the tab started spinning.
	EvtWebContentsDidStartLoading = "did-start-loading"
	// Corresponds to the points in time when the spinner of the tab stopped spinning.
	EvtWebContentsDidStopLoading = "did-stop-loading"
	// Emitted when details regarding a requested resource are available. status indicates the socket connection to download the resource.
	EvtWebContentsDidGetResponseDetails = "did-get-response-details"
	// Emitted when a redirect is received while requesting a resource.
	EvtWebContentsDidGetRedirectRequest = "did-get-redirect-request"
	// Emitted when the document in the given frame is loaded.
	EvtWebContentsDomReady = "dom-ready"
	// Emitted when page receives favicon urls.
	EvtWebContentsPageFaviconUpdated = "page-favicon-updated"
	// Emitted when the page requests to open a new window for a url. It could be requested by window.open or an external link like . By default a new BrowserWindow will be created for the url. Calling event.preventDefault() will prevent creating new windows. In such case, the event.newGuest may be set with a reference to a BrowserWindow instance to make it used by the Electron's runtime.
	EvtWebContentsNewWindow = "new-window"
	// Emitted when a user or the page wants to start navigation. It can happen when the window.location object is changed or a user clicks a link in the page. This event will not emit when the navigation is started programmatically with APIs like webContents.loadURL and webContents.back. It is also not emitted for in-page navigations, such as clicking anchor links or updating the window.location.hash. Use did-navigate-in-page event for this purpose. Calling event.preventDefault() will prevent the navigation.
	EvtWebContentsWillNavigate = "will-navigate"
	// Emitted when a navigation is done. This event is not emitted for in-page navigations, such as clicking anchor links or updating the window.location.hash. Use did-navigate-in-page event for this purpose.
	EvtWebContentsDidNavigate = "did-navigate"
	// Emitted when an in-page navigation happened. When in-page navigation happens, the page URL changes but does not cause navigation outside of the page. Examples of this occurring are when anchor links are clicked or when the DOM hashchange event is triggered.
	EvtWebContentsDidNavigateInPage = "did-navigate-in-page"
	// Emitted when the renderer process crashes or is killed.
	EvtWebContentsCrashed = "crashed"
	// Emitted when a plugin process has crashed.
	EvtWebContentsPluginCrashed = "plugin-crashed"
	// Emitted when webContents is destroyed.
	EvtWebContentsDestroyed = "destroyed"
	// Emitted before dispatching the keydown and keyup events in the page. Calling event.preventDefault will prevent the page keydown/keyup events from being dispatched.
	EvtWebContentsBeforeInputEvent = "before-input-event"
	// Emitted when DevTools is opened.
	EvtWebContentsDevtoolsOpened = "devtools-opened"
	// Emitted when DevTools is closed.
	EvtWebContentsDevtoolsClosed = "devtools-closed"
	// Emitted when DevTools is focused / opened.
	EvtWebContentsDevtoolsFocused = "devtools-focused"
	// Emitted when failed to verify the certificate for url. The usage is the same with the certificate-error event of app.
	EvtWebContentsCertificateError = "certificate-error"
	// Emitted when a client certificate is requested. The usage is the same with the select-client-certificate event of app.
	EvtWebContentsSelectClientCertificate = "select-client-certificate"
	// Emitted when webContents wants to do basic auth. The usage is the same with the login event of app.
	EvtWebContentsLogin = "login"
	// Emitted when a result is available for [webContents.findInPage] request.
	EvtWebContentsFoundInPage = "found-in-page"
	// Emitted when media starts playing.
	EvtWebContentsMediaStartedPlaying = "media-started-playing"
	// Emitted when media is paused or done playing.
	EvtWebContentsMediaPaused = "media-paused"
	// Emitted when a page's theme color changes. This is usually due to encountering a meta tag:
	EvtWebContentsDidChangeThemeColor = "did-change-theme-color"
	// Emitted when mouse moves over a link or the keyboard moves the focus to a link.
	EvtWebContentsUpdateTargetURL = "update-target-url"
	// Emitted when the cursor's type changes. The type parameter can be default, crosshair, pointer, text, wait, help, e-resize, n-resize, ne-resize, nw-resize, s-resize, se-resize, sw-resize, w-resize, ns-resize, ew-resize, nesw-resize, nwse-resize, col-resize, row-resize, m-panning, e-panning, n-panning, ne-panning, nw-panning, s-panning, se-panning, sw-panning, w-panning, move, vertical-text, cell, context-menu, alias, progress, nodrop, copy, none, not-allowed, zoom-in, zoom-out, grab, grabbing, custom. If the type parameter is custom, the image parameter will hold the custom cursor image in a NativeImage, and scale, size and hotspot will hold additional information about the custom cursor.
	EvtWebContentsCursorChanged = "cursor-changed"
	// Emitted when there is a new context menu that needs to be handled.
	EvtWebContentsContextMenu = "context-menu"
	// Emitted when bluetooth device needs to be selected on call to navigator.bluetooth.requestDevice. To use navigator.bluetooth api webBluetooth should be enabled.  If event.preventDefault is not called, first available device will be selected. callback should be called with deviceId to be selected, passing empty string to callback will cancel the request.
	EvtWebContentsSelectBluetoothDevice = "select-bluetooth-device"
	// Emitted when a new frame is generated. Only the dirty area is passed in the buffer.
	EvtWebContentsPaint = "paint"
	// Emitted when the devtools window instructs the webContents to reload
	EvtWebContentsDevtoolsReloadPage = "devtools-reload-page"
)

// WebContents version@1.4.15
//
// Render and control the contents of a BrowserWindow instance.
type WebContents struct {
	*events.Emitter
	// A Integer representing the unique ID of this WebContents.
	Id int64 `js:"id"`
	// A Session object (session) used by this webContents.
	Session *Session `js:"session"`
	// A WebContents instance that might own this WebContents.
	HostWebContents *WebContents `js:"hostWebContents"`
	// A WebContents of DevTools for this WebContents. Note: Users should never store this object because it may become null when the DevTools has been closed.
	DevToolsWebContents *WebContents `js:"devToolsWebContents"`
	// A Debugger instance for this webContents.
	Debugger *js.Object `js:"debugger"`
	// Loads the url in the window. The url must contain the protocol prefix, e.g. the http:// or file://. If the load should bypass http cache then use the pragma header to achieve it.
	LoadURL func(URL string, Options *WebContentsLoadURLOptions) `js:"loadURL"`
	// Initiates a download of the resource at url without navigating. The will-download event of session will be triggered.
	DownloadURL          func(URL string)    `js:"downloadURL"`
	GetURL               func() (Obj string) `js:"getURL"`
	GetTitle             func() (Obj string) `js:"getTitle"`
	IsDestroyed          func() (Obj bool)   `js:"isDestroyed"`
	IsFocused            func() (Obj bool)   `js:"isFocused"`
	IsLoading            func() (Obj bool)   `js:"isLoading"`
	IsLoadingMainFrame   func() (Obj bool)   `js:"isLoadingMainFrame"`
	IsWaitingForResponse func() (Obj bool)   `js:"isWaitingForResponse"`
	// Stops any pending navigation.
	Stop func() `js:"stop"`
	// Reloads the current web page.
	Reload func() `js:"reload"`
	// Reloads current page and ignores cache.
	ReloadIgnoringCache func()                        `js:"reloadIgnoringCache"`
	CanGoBack           func() (Obj bool)             `js:"canGoBack"`
	CanGoForward        func() (Obj bool)             `js:"canGoForward"`
	CanGoToOffset       func(Offset int64) (Obj bool) `js:"canGoToOffset"`
	// Clears the navigation history.
	ClearHistory func() `js:"clearHistory"`
	// Makes the browser go back a web page.
	GoBack func() `js:"goBack"`
	// Makes the browser go forward a web page.
	GoForward func() `js:"goForward"`
	// Navigates browser to the specified absolute web page index.
	GoToIndex func(Index int64) `js:"goToIndex"`
	// Navigates to the specified offset from the "current entry".
	GoToOffset func(Offset int64) `js:"goToOffset"`
	IsCrashed  func() (Obj bool)  `js:"isCrashed"`
	// Overrides the user agent for this web page.
	SetUserAgent func(UserAgent string) `js:"setUserAgent"`
	GetUserAgent func() (Obj string)    `js:"getUserAgent"`
	// Injects CSS into the current web page.
	InsertCSS func(Css string) `js:"insertCSS"`
	// Evaluates code in page. In the browser window some HTML APIs like requestFullScreen can only be invoked by a gesture from the user. Setting userGesture to true will remove this limitation. If the result of the executed code is a promise the callback result will be the resolved value of the promise.  We recommend that you use the returned Promise to handle code that results in a Promise.
	ExecuteJavaScript func(Code string, UserGesture bool, Callback WebContentsExecuteJavaScriptCallback) (Obj *js.Object) `js:"executeJavaScript"`
	// Mute the audio on the current web page.
	SetAudioMuted func(Muted bool)  `js:"setAudioMuted"`
	IsAudioMuted  func() (Obj bool) `js:"isAudioMuted"`
	// Changes the zoom factor to the specified factor. Zoom factor is zoom percent divided by 100, so 300% = 3.0.
	SetZoomFactor func(Factor float64) `js:"setZoomFactor"`
	// Sends a request to get current zoom factor, the callback will be called with callback(zoomFactor).
	GetZoomFactor func(Callback WebContentsGetZoomFactorCallback) `js:"getZoomFactor"`
	// Changes the zoom level to the specified level. The original size is 0 and each increment above or below represents zooming 20% larger or smaller to default limits of 300% and 50% of original size, respectively.
	SetZoomLevel func(Level float64) `js:"setZoomLevel"`
	// Sends a request to get current zoom level, the callback will be called with callback(zoomLevel).
	GetZoomLevel func(Callback WebContentsGetZoomLevelCallback) `js:"getZoomLevel"`
	// Deprecated: Call setVisualZoomLevelLimits instead to set the visual zoom level limits. This method will be removed in Electron 2.0.
	SetZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setZoomLevelLimits"`
	// Sets the maximum and minimum pinch-to-zoom level.
	SetVisualZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setVisualZoomLevelLimits"`
	// Sets the maximum and minimum layout-based (i.e. non-visual) zoom level.
	SetLayoutZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setLayoutZoomLevelLimits"`
	// Executes the editing command undo in web page.
	Undo func() `js:"undo"`
	// Executes the editing command redo in web page.
	Redo func() `js:"redo"`
	// Executes the editing command cut in web page.
	Cut func() `js:"cut"`
	// Executes the editing command copy in web page.
	Copy func() `js:"copy"`
	// Copy the image at the given position to the clipboard.
	CopyImageAt func(X int64, Y int64) `js:"copyImageAt"`
	// Executes the editing command paste in web page.
	Paste func() `js:"paste"`
	// Executes the editing command pasteAndMatchStyle in web page.
	PasteAndMatchStyle func() `js:"pasteAndMatchStyle"`
	// Executes the editing command delete in web page.
	Delete func() `js:"delete"`
	// Executes the editing command selectAll in web page.
	SelectAll func() `js:"selectAll"`
	// Executes the editing command unselect in web page.
	Unselect func() `js:"unselect"`
	// Executes the editing command replace in web page.
	Replace func(Text string) `js:"replace"`
	// Executes the editing command replaceMisspelling in web page.
	ReplaceMisspelling func(Text string) `js:"replaceMisspelling"`
	// Inserts text to the focused element.
	InsertText func(Text string) `js:"insertText"`
	// Starts a request to find all matches for the text in the web page and returns an Integer representing the request id used for the request. The result of the request can be obtained by subscribing to found-in-page event.
	FindInPage func(Text string, Options *WebContentsFindInPageOptions) `js:"findInPage"`
	// Stops any findInPage request for the webContents with the provided action.
	StopFindInPage func(Action WebContentsStopFindInPageAction) `js:"stopFindInPage"`
	// Captures a snapshot of the page within rect. Upon completion callback will be called with callback(image). The image is an instance of NativeImage that stores data of the snapshot. Omitting rect will capture the whole visible page.
	CapturePage func(Rect *js.Object, Callback WebContentsCapturePageCallback) `js:"capturePage"`
	// Checks if any ServiceWorker is registered and returns a boolean as response to callback.
	HasServiceWorker func(Callback WebContentsHasServiceWorkerCallback) `js:"hasServiceWorker"`
	// Unregisters any ServiceWorker if present and returns a boolean as response to callback when the JS promise is fulfilled or false when the JS promise is rejected.
	UnregisterServiceWorker func(Callback WebContentsUnregisterServiceWorkerCallback) `js:"unregisterServiceWorker"`
	// Prints window's web page. When silent is set to true, Electron will pick up system's default printer and default settings for printing. Calling window.print() in web page is equivalent to calling webContents.print({silent: false, printBackground: false}). Use page-break-before: always; CSS style to force to print to a new page.
	Print func(Options *WebContentsPrintOptions) `js:"print"`
	// Prints window's web page as PDF with Chromium's preview printing custom settings. The callback will be called with callback(error, data) on completion. The data is a Buffer that contains the generated PDF data. The landscape will be ignored if @page CSS at-rule is used in the web page. By default, an empty options will be regarded as: Use page-break-before: always; CSS style to force to print to a new page. An example of webContents.printToPDF:
	PrintToPDF func(Options *WebContentsPrintToPDFOptions, Callback WebContentsPrintToPDFCallback) `js:"printToPDF"`
	// Adds the specified path to DevTools workspace. Must be used after DevTools creation:
	AddWorkSpace func(Path string) `js:"addWorkSpace"`
	// Removes the specified path from DevTools workspace.
	RemoveWorkSpace func(Path string) `js:"removeWorkSpace"`
	// Opens the devtools.
	OpenDevTools func(Options *WebContentsOpenDevToolsOptions) `js:"openDevTools"`
	// Closes the devtools.
	CloseDevTools     func()            `js:"closeDevTools"`
	IsDevToolsOpened  func() (Obj bool) `js:"isDevToolsOpened"`
	IsDevToolsFocused func() (Obj bool) `js:"isDevToolsFocused"`
	// Toggles the developer tools.
	ToggleDevTools func() `js:"toggleDevTools"`
	// Starts inspecting element at position (x, y).
	InspectElement func(X int64, Y int64) `js:"inspectElement"`
	// Opens the developer tools for the service worker context.
	InspectServiceWorker func() `js:"inspectServiceWorker"`
	// Send an asynchronous message to renderer process via channel, you can also send arbitrary arguments. Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included. The renderer process can handle the message by listening to channel with the ipcRenderer module. An example of sending messages from the main process to the renderer process:
	Send func(Channel string, Args *js.Object) `js:"send"`
	// Enable device emulation with the given parameters.
	EnableDeviceEmulation func(Parameters *WebContentsEnableDeviceEmulationParameters) `js:"enableDeviceEmulation"`
	// Disable device emulation enabled by webContents.enableDeviceEmulation.
	DisableDeviceEmulation func() `js:"disableDeviceEmulation"`
	// Sends an input event to the page. For keyboard events, the event object also have following properties: For mouse events, the event object also have following properties: For the mouseWheel event, the event object also have following properties:
	SendInputEvent func(Event *WebContentsSendInputEventEvent) `js:"sendInputEvent"`
	// Begin subscribing for presentation events and captured frames, the callback will be called with callback(frameBuffer, dirtyRect) when there is a presentation event. The frameBuffer is a Buffer that contains raw pixel data. On most machines, the pixel data is effectively stored in 32bit BGRA format, but the actual representation depends on the endianness of the processor (most modern processors are little-endian, on machines with big-endian processors the data is in 32bit ARGB format). The dirtyRect is an object with x, y, width, height properties that describes which part of the page was repainted. If onlyDirty is set to true, frameBuffer will only contain the repainted area. onlyDirty defaults to false.
	BeginFrameSubscription func(OnlyDirty bool, Callback WebContentsBeginFrameSubscriptionCallback) `js:"beginFrameSubscription"`
	// End subscribing for frame presentation events.
	EndFrameSubscription func() `js:"endFrameSubscription"`
	// Sets the item as dragging item for current drag-drop operation, file is the absolute path of the file to be dragged, and icon is the image showing under the cursor when dragging.
	StartDrag func(Item *WebContentsStartDragItem) `js:"startDrag"`
	// Returns true if the process of saving page has been initiated successfully.
	SavePage func(FullPath string, SaveType WebContentsSavePageSaveType, Callback WebContentsSavePageCallback) `js:"savePage"`
	// Shows pop-up dictionary that searches the selected word on the page.
	ShowDefinitionForSelection func() `js:"showDefinitionForSelection"`
	// Set the size of the page. This is only supported for  guest contents.
	SetSize     func(Options *WebContentsSetSizeOptions) `js:"setSize"`
	IsOffscreen func() (Obj bool)                        `js:"isOffscreen"`
	// If offscreen rendering is enabled and not painting, start painting.
	StartPainting func() `js:"startPainting"`
	// If offscreen rendering is enabled and painting, stop painting.
	StopPainting func()            `js:"stopPainting"`
	IsPainting   func() (Obj bool) `js:"isPainting"`
	// If offscreen rendering is enabled sets the frame rate to the specified number. Only values between 1 and 60 are accepted.
	SetFrameRate func(Fps int64)    `js:"setFrameRate"`
	GetFrameRate func() (Obj int64) `js:"getFrameRate"`
	// If offscreen rendering is enabled invalidates the frame and generates a new one through the 'paint' event.
	Invalidate func() `js:"invalidate"`
}

func WrapWebContents(o *js.Object) *WebContents {
	return &WebContents{
		Emitter: events.New(o),
	}
}

type WebContentsExecuteJavaScriptCallback func(Result *js.Object)
type WebContentsGetZoomFactorCallback func(ZoomFactor float64)
type WebContentsPrintToPDFCallback func(Error *js.Object, Data *js.Object)
type WebContentsSavePageCallback func(Error *js.Object)
type WebContentsLoadURLOptions struct {
	*js.Object
	// A HTTP Referrer url.
	HttpReferrer string `js:"httpReferrer"`
	// A user agent originating the request.
	UserAgent string `js:"userAgent"`
	// Extra headers separated by "\n"
	ExtraHeaders string `js:"extraHeaders"`
	// [] (optional)
	PostData *js.Object `js:"postData"`
}

type WebContentsFindInPageOptions struct {
	*js.Object
	// (optional) Whether to search forward or backward, defaults to .
	Forward bool `js:"forward"`
	// (optional) Whether the operation is first request or a follow up, defaults to .
	FindNext bool `js:"findNext"`
	// (optional) Whether search should be case-sensitive, defaults to .
	MatchCase bool `js:"matchCase"`
	// (optional) Whether to look only at the start of words. defaults to .
	WordStart bool `js:"wordStart"`
	// (optional) When combined with , accepts a match in the middle of a word if the match begins with an uppercase letter followed by a lowercase or non-letter. Accepts several other intra-word matches, defaults to .
	MedialCapitalAsWordStart bool `js:"medialCapitalAsWordStart"`
}

type WebContentsOpenDevToolsOptions struct {
	*js.Object
	// Opens the devtools with specified dock state, can be , , , . Defaults to last used dock state. In mode it's possible to dock back. In mode it's not.
	Mode WebContentsOptionsMode `js:"mode"`
}

type WebContentsOptionsMode string

// consts
const (
	WebContentsOptionsModeRight    WebContentsOptionsMode = "right"
	WebContentsOptionsModeBottom   WebContentsOptionsMode = "bottom"
	WebContentsOptionsModeUndocked WebContentsOptionsMode = "undocked"
	WebContentsOptionsModeDetach   WebContentsOptionsMode = "detach"
)

type WebContentsEnableDeviceEmulationParameters struct {
	*js.Object
	// Specify the screen type to emulate (default: )
	ScreenPosition WebContentsParametersScreenPosition `js:"screenPosition"`
	// Set the emulated screen size (screenPosition == mobile)
	ScreenSize *WebContentsParametersScreenSize `js:"screenSize"`
	// Position the view on the screen (screenPosition == mobile) (default: )
	ViewPosition *WebContentsParametersViewPosition `js:"viewPosition"`
	// Set the device scale factor (if zero defaults to original device scale factor) (default: )
	DeviceScaleFactor int64 `js:"deviceScaleFactor"`
	// Set the emulated view size (empty means no override)
	ViewSize *WebContentsParametersViewSize `js:"viewSize"`
	// Whether emulated view should be scaled down if necessary to fit into available space (default: )
	FitToView bool `js:"fitToView"`
	// Offset of the emulated view inside available space (not in fit to view mode) (default: )
	Offset *WebContentsParametersOffset `js:"offset"`
	// Scale of emulated view inside available space (not in fit to view mode) (default: )
	Scale float64 `js:"scale"`
}

type WebContentsParametersViewPosition struct {
	*js.Object
	// Set the x axis offset from top left corner
	X int64 `js:"x"`
	// Set the y axis offset from top left corner
	Y int64 `js:"y"`
}

type WebContentsParametersViewSize struct {
	*js.Object
	// Set the emulated view width
	Width int64 `js:"width"`
	// Set the emulated view height
	Height int64 `js:"height"`
}

type WebContentsParametersOffset struct {
	*js.Object
	// Set the x axis offset from top left corner
	X float64 `js:"x"`
	// Set the y axis offset from top left corner
	Y float64 `js:"y"`
}

type WebContentsParametersScreenSize struct {
	*js.Object
	// Set the emulated screen width
	Width int64 `js:"width"`
	// Set the emulated screen height
	Height int64 `js:"height"`
}

type WebContentsParametersScreenPosition string

// consts
const (
	WebContentsParametersScreenPositionDesktop WebContentsParametersScreenPosition = "desktop"
	WebContentsParametersScreenPositionMobile  WebContentsParametersScreenPosition = "mobile"
)

type WebContentsBeginFrameSubscriptionCallback func(FrameBuffer *js.Object, DirtyRect *js.Object)
type WebContentsStartDragItem struct {
	*js.Object
	File string       `js:"file"`
	Icon *NativeImage `js:"icon"`
}

type WebContentsSetSizeOptions struct {
	*js.Object
	// Normal size of the page. This can be used in combination with the attribute to manually resize the webview guest contents.
	Normal *WebContentsOptionsNormal `js:"normal"`
}

type WebContentsOptionsNormal struct {
	*js.Object
	Width  int64 `js:"width"`
	Height int64 `js:"height"`
}

type WebContentsGetZoomLevelCallback func(ZoomLevel float64)
type WebContentsHasServiceWorkerCallback func(HasWorker bool)
type WebContentsUnregisterServiceWorkerCallback func(Success bool)
type WebContentsPrintToPDFOptions struct {
	*js.Object
	// (optional) Specifies the type of margins to use. Uses 0 for default margin, 1 for no margin, and 2 for minimum margin.
	MarginsType int64 `js:"marginsType"`
	// (optional) Specify page size of the generated PDF. Can be , , , , , or an Object containing and in microns.
	PageSize string `js:"pageSize"`
	// (optional) Whether to print CSS backgrounds.
	PrintBackground bool `js:"printBackground"`
	// (optional) Whether to print selection only.
	PrintSelectionOnly bool `js:"printSelectionOnly"`
	// (optional) for landscape, for portrait.
	Landscape bool `js:"landscape"`
}

type WebContentsCapturePageCallback func(Image *NativeImage)
type WebContentsPrintOptions struct {
	*js.Object
	// Don't ask user for print settings. Default is .
	Silent bool `js:"silent"`
	// Also prints the background color and image of the web page. Default is .
	PrintBackground bool `js:"printBackground"`
}

type WebContentsSendInputEventEvent struct {
	*js.Object
	// () The type of the event, can be , , , , , , , , , .
	Type WebContentsEventType `js:"type"`
	// An array of modifiers of the event, can include , , , , , , , , , , , , .
	Modifiers *js.Object `js:"modifiers"`
}

type WebContentsEventType string

// consts
const (
	WebContentsEventTypeMouseDown   WebContentsEventType = "mouseDown"
	WebContentsEventTypeMouseUp     WebContentsEventType = "mouseUp"
	WebContentsEventTypeMouseEnter  WebContentsEventType = "mouseEnter"
	WebContentsEventTypeMouseLeave  WebContentsEventType = "mouseLeave"
	WebContentsEventTypeContextMenu WebContentsEventType = "contextMenu"
	WebContentsEventTypeMouseWheel  WebContentsEventType = "mouseWheel"
	WebContentsEventTypeMouseMove   WebContentsEventType = "mouseMove"
	WebContentsEventTypeKeyDown     WebContentsEventType = "keyDown"
	WebContentsEventTypeKeyUp       WebContentsEventType = "keyUp"
	WebContentsEventTypeChar        WebContentsEventType = "char"
)

type WebContentsStopFindInPageAction string

// consts
const (
	WebContentsStopFindInPageActionClearSelection    WebContentsStopFindInPageAction = "clearSelection"
	WebContentsStopFindInPageActionKeepSelection     WebContentsStopFindInPageAction = "keepSelection"
	WebContentsStopFindInPageActionActivateSelection WebContentsStopFindInPageAction = "activateSelection"
)

type WebContentsSavePageSaveType string

// consts
const (
	WebContentsSavePageSaveTypeHTMLOnly     WebContentsSavePageSaveType = "HTMLOnly"
	WebContentsSavePageSaveTypeHTMLComplete WebContentsSavePageSaveType = "HTMLComplete"
	WebContentsSavePageSaveTypeMHTML        WebContentsSavePageSaveType = "MHTML"
)
