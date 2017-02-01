package webcontents

import (
	"github.com/gopherjs/gopherjs/js"
)

// Methods

// These methods can be accessed from the webContents module:

// const {webContents} = require('electron')
// console.log(webContents)
// webContents.getAllWebContents()
// Returns WebContents[] - An array of all WebContents instances. This will contain web contents for all windows, webviews, opened devtools, and devtools extension background pages.

// webContents.getFocusedWebContents()
// Returns WebContents - The web contents that is focused in this application, otherwise returns null.

// webContents.fromId(id)
// id Integer
// Returns WebContents - A WebContents instance with the given ID.

// ====================================================================

// Class: WebContents

// Render and control the contents of a BrowserWindow instance.
// Process: Main

const (
	EvtDidFinishLoad           = "did-fail-load"
	EvtDidFailLoad             = "did-fail-load"
	EvtDidFrameFinishLoad      = "did-frame-finish-load"
	EvtDidStartLoading         = "did-start-loading"
	EvtDidStopLoading          = "did-stop-loading"
	EvtDidGetResponseDetails   = "did-get-response-details"
	EvtDidGetRedirectRequest   = "did-get-redirect-request"
	EvtDomReady                = "dom-ready"
	EvtPageFaviconUpdated      = "page-favicon-updated"
	EvtNewWindow               = "new-window"
	EvtWillNavigate            = "will-navigate"
	EvtDidNavigate             = "did-navigate"
	EvtDidNavigateInPage       = "did-navigate-in-page"
	EvtCrashed                 = "crashed"
	EvtPluginCrashed           = "plugin-crashed"
	EvtDestroyed               = "destroyed"
	EvtBeforeInputEvent        = "before-input-event"
	EvtDevtoolsOpened          = "devtools-opened"
	EvtDevtoolsClosed          = "devtools-closed"
	EvtDevtoolsFocused         = "devtools-focused"
	EvtCertificateError        = "certificate-error"
	EvtSelectClientCertificate = "select-client-certificate"
	EvtLogin                   = "login"
	EvtFoundInPage             = "found-in-page"
	EvtMediaStartedPlaying     = "media-started-playing"
	EvtMediaPaused             = "media-paused"
	EvtDidChangeThemeColor     = "did-change-theme-color"
	EvtUpdateTargetURL         = "update-target-url"
	EvtCursorChanged           = "cursor-changed"
	EvtContextMenu             = "context-menu"
	EvtSelectBluetoothDevice   = "select-bluetooth-device"
	EvtPaint                   = "paint"
	EvtDevtoolsReloadPage      = "devtools-reload-page"
)

// ============================================================================
// ============================================================================

// WebContents wrapper of webContents
type WebContents struct {
	//    Instance Methods
	// contents.loadURL(url[, options])
	// url String
	// options Object (optional)
	// httpReferrer String (optional) - A HTTP Referrer url.
	// userAgent String (optional) - A user agent originating the request.
	// extraHeaders String (optional) - Extra headers separated by “\n”
	// postData (UploadRawData | UploadFile | UploadFileSystem | UploadBlob)[] - (optional)
	// Loads the url in the window. The url must contain the protocol prefix, e.g. the http:// or file://. If the load should bypass http cache then use the pragma header to achieve it.
	LoadURL func(url string) `js:"loadURL"`

	// const {webContents} = require('electron')
	// const options = {extraHeaders: 'pragma: no-cache\n'}
	// webContents.loadURL('https://github.com', options)
	// contents.downloadURL(url)
	// url String
	// Initiates a download of the resource at url without navigating. The will-download event of session will be triggered.

	// contents.getURL()
	// Returns String - The URL of the current web page.
	GetURL func() string `js:"getURL"`

	// const {BrowserWindow} = require('electron')
	// let win = new BrowserWindow({width: 800, height: 600})
	// win.loadURL('http://github.com')

	// let currentURL = win.webContents.getURL()
	// console.log(currentURL)
	// contents.getTitle()
	// Returns String - The title of the current web page.

	// contents.isDestroyed()
	// Returns Boolean - Whether the web page is destroyed.
	IsDestroyed func() bool `js:"isDestroyed"`

	// contents.isFocused()
	// Returns Boolean - Whether the web page is focused.
	IsFocused func() bool `js:"isFocused"`

	// contents.isLoading()
	// Returns Boolean - Whether web page is still loading resources.
	IsLoading func() bool `js:"isLoading"`

	// contents.isLoadingMainFrame()
	// Returns Boolean - Whether the main frame (and not just iframes or frames within it) is still loading.
	IsLoadingMainFrame func() bool `js:"isLoadingMainFrame"`

	// contents.isWaitingForResponse()
	// Returns Boolean - Whether the web page is waiting for a first-response from the main resource of the page.
	IsWaitingForResponse func() bool `js:"isWaitingForResponse"`

	// contents.stop()
	// Stops any pending navigation.
	Stop func() `js:"stop"`

	// contents.reload()
	// Reloads the current web page.
	Reload func() `js:"reload"`

	// contents.reloadIgnoringCache()
	// Reloads current page and ignores cache.
	ReloadIgnoringCache func() `js:"reloadIgnoringCache"`

	// contents.canGoBack()
	// Returns Boolean - Whether the browser can go back to previous web page.
	CanGoBack func() bool `js:"canGoBack"`

	// contents.canGoForward()
	// Returns Boolean - Whether the browser can go forward to next web page.
	CanGoForward func() bool `js:"canGoForward"`

	// contents.canGoToOffset(offset)
	// offset Integer
	// Returns Boolean - Whether the web page can go to offset.
	CanGoToOffset func(offset int) bool `js:"canGoToOffset"`

	// contents.clearHistory()
	// Clears the navigation history.
	ClearHistory func() `js:"clearHistory"`

	// contents.goBack()
	// Makes the browser go back a web page.
	GoBack func() `js:"goBack"`

	// contents.goForward()
	// Makes the browser go forward a web page.

	// contents.goToIndex(index)
	// index Integer
	// Navigates browser to the specified absolute web page index.

	// contents.goToOffset(offset)
	// offset Integer
	// Navigates to the specified offset from the “current entry”.

	// contents.isCrashed()
	// Returns Boolean - Whether the renderer process has crashed.

	// contents.setUserAgent(userAgent)
	// userAgent String
	// Overrides the user agent for this web page.

	// contents.getUserAgent()
	// Returns String - The user agent for this web page.

	// contents.insertCSS(css)
	// css String
	// Injects CSS into the current web page.

	// contents.executeJavaScript(code[, userGesture, callback])
	// code String
	// userGesture Boolean (optional)
	// callback Function (optional) - Called after script has been executed.
	// result Any
	// Returns Promise - A promise that resolves with the result of the executed code or is rejected if the result of the code is a rejected promise.
	ExecuteJavaScript func(code string, userGesture bool) *js.Object `js:"executeJavaScript"`

	// Evaluates code in page.

	// In the browser window some HTML APIs like requestFullScreen can only be invoked by a gesture from the user. Setting userGesture to true will remove this limitation.

	// If the result of the executed code is a promise the callback result will be the resolved value of the promise. We recommend that you use the returned Promise to handle code that results in a Promise.

	// contents.executeJavaScript('fetch("https://jsonplaceholder.typicode.com/users/1").then(resp => resp.json())', true)
	//   .then((result) => {
	//     console.log(result) // Will be the JSON object from the fetch call
	//   })
	// contents.setAudioMuted(muted)
	// muted Boolean
	// Mute the audio on the current web page.

	// contents.isAudioMuted()
	// Returns Boolean - Whether this page has been muted.

	// contents.setZoomFactor(factor)
	// factor Number - Zoom factor.
	// Changes the zoom factor to the specified factor. Zoom factor is zoom percent divided by 100, so 300% = 3.0.

	// contents.getZoomFactor(callback)
	// callback Function
	// zoomFactor Number
	// Sends a request to get current zoom factor, the callback will be called with callback(zoomFactor).

	// contents.setZoomLevel(level)
	// level Number - Zoom level
	// Changes the zoom level to the specified level. The original size is 0 and each increment above or below represents zooming 20% larger or smaller to default limits of 300% and 50% of original size, respectively.

	// contents.getZoomLevel(callback)
	// callback Function
	// zoomLevel Number
	// Sends a request to get current zoom level, the callback will be called with callback(zoomLevel).

	// contents.setZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Deprecated: Call setVisualZoomLevelLimits instead to set the visual zoom level limits. This method will be removed in Electron 2.0.

	// contents.setVisualZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Sets the maximum and minimum pinch-to-zoom level.

	// contents.setLayoutZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Sets the maximum and minimum layout-based (i.e. non-visual) zoom level.

	// contents.undo()
	// Executes the editing command undo in web page.

	// contents.redo()
	// Executes the editing command redo in web page.

	// contents.cut()
	// Executes the editing command cut in web page.

	// contents.copy()
	// Executes the editing command copy in web page.

	// contents.copyImageAt(x, y)
	// x Integer
	// y Integer
	// Copy the image at the given position to the clipboard.

	// contents.paste()
	// Executes the editing command paste in web page.

	// contents.pasteAndMatchStyle()
	// Executes the editing command pasteAndMatchStyle in web page.

	// contents.delete()
	// Executes the editing command delete in web page.

	// contents.selectAll()
	// Executes the editing command selectAll in web page.

	// contents.unselect()
	// Executes the editing command unselect in web page.

	// contents.replace(text)
	// text String
	// Executes the editing command replace in web page.

	// contents.replaceMisspelling(text)
	// text String
	// Executes the editing command replaceMisspelling in web page.

	// contents.insertText(text)
	// text String
	// Inserts text to the focused element.
	InsertText func(text string) `js:"insertText"`

	// contents.findInPage(text[, options])
	// text String - Content to be searched, must not be empty.
	// options Object (optional)
	// forward Boolean - (optional) Whether to search forward or backward, defaults to true.
	// findNext Boolean - (optional) Whether the operation is first request or a follow up, defaults to false.
	// matchCase Boolean - (optional) Whether search should be case-sensitive, defaults to false.
	// wordStart Boolean - (optional) Whether to look only at the start of words. defaults to false.
	// medialCapitalAsWordStart Boolean - (optional) When combined with wordStart, accepts a match in the middle of a word if the match begins with an uppercase letter followed by a lowercase or non-letter. Accepts several other intra-word matches, defaults to false.
	// Starts a request to find all matches for the text in the web page and returns an Integer representing the request id used for the request. The result of the request can be obtained by subscribing to found-in-page event.

	// contents.stopFindInPage(action)
	// action String - Specifies the action to take place when ending [webContents.findInPage] request.
	// clearSelection - Clear the selection.
	// keepSelection - Translate the selection into a normal selection.
	// activateSelection - Focus and click the selection node.
	// Stops any findInPage request for the webContents with the provided action.

	// const {webContents} = require('electron')
	// webContents.on('found-in-page', (event, result) => {
	//   if (result.finalUpdate) webContents.stopFindInPage('clearSelection')
	// })

	// const requestId = webContents.findInPage('api')
	// console.log(requestId)
	// contents.capturePage([rect, ]callback)
	// rect Rectangle (optional) - The area of the page to be captured
	// callback Function
	// image NativeImage
	// Captures a snapshot of the page within rect. Upon completion callback will be called with callback(image). The image is an instance of NativeImage that stores data of the snapshot. Omitting rect will capture the whole visible page.

	// contents.hasServiceWorker(callback)
	// callback Function
	// hasWorker Boolean
	// Checks if any ServiceWorker is registered and returns a boolean as response to callback.

	// contents.unregisterServiceWorker(callback)
	// callback Function
	// success Boolean
	// Unregisters any ServiceWorker if present and returns a boolean as response to callback when the JS promise is fulfilled or false when the JS promise is rejected.

	// contents.print([options])
	// options Object (optional)
	// silent Boolean - Don’t ask user for print settings. Default is false.
	// printBackground Boolean - Also prints the background color and image of the web page. Default is false.
	// Prints window’s web page. When silent is set to true, Electron will pick up system’s default printer and default settings for printing.
	Print   func()                   `js:"print"`
	PrintEx func(options *js.Object) `js:"print"`

	// Calling window.print() in web page is equivalent to calling webContents.print({silent: false, printBackground: false}).

	// Use page-break-before: always; CSS style to force to print to a new page.

	// contents.printToPDF(options, callback)
	// options Object
	// marginsType Integer - (optional) Specifies the type of margins to use. Uses 0 for default margin, 1 for no margin, and 2 for minimum margin.
	// pageSize String - (optional) Specify page size of the generated PDF. Can be A3, A4, A5, Legal, Letter, Tabloid or an Object containing height and width in microns.
	// printBackground Boolean - (optional) Whether to print CSS backgrounds.
	// printSelectionOnly Boolean - (optional) Whether to print selection only.
	// landscape Boolean - (optional) true for landscape, false for portrait.
	// callback Function
	// error Error
	// data Buffer
	// Prints window’s web page as PDF with Chromium’s preview printing custom settings.

	// The callback will be called with callback(error, data) on completion. The data is a Buffer that contains the generated PDF data.

	// The landscape will be ignored if @page CSS at-rule is used in the web page.

	// By default, an empty options will be regarded as:

	// {
	//   marginsType: 0,
	//   printBackground: false,
	//   printSelectionOnly: false,
	//   landscape: false
	// }
	// Use page-break-before: always; CSS style to force to print to a new page.

	// An example of webContents.printToPDF:

	// const {BrowserWindow} = require('electron')
	// const fs = require('fs')

	// let win = new BrowserWindow({width: 800, height: 600})
	// win.loadURL('http://github.com')

	// win.webContents.on('did-finish-load', () => {
	//   // Use default printing options
	//   win.webContents.printToPDF({}, (error, data) => {
	//     if (error) throw error
	//     fs.writeFile('/tmp/print.pdf', data, (error) => {
	//       if (error) throw error
	//       console.log('Write PDF successfully.')
	//     })
	//   })
	// })
	// contents.addWorkSpace(path)
	// path String
	// Adds the specified path to DevTools workspace. Must be used after DevTools creation:

	// const {BrowserWindow} = require('electron')
	// let win = new BrowserWindow()
	// win.webContents.on('devtools-opened', () => {
	//   win.webContents.addWorkSpace(__dirname)
	// })
	// contents.removeWorkSpace(path)
	// path String
	// Removes the specified path from DevTools workspace.

	// contents.openDevTools([options])
	// options Object (optional)
	// mode String - Opens the devtools with specified dock state, can be right, bottom, undocked, detach. Defaults to last used dock state. In undocked mode it’s possible to dock back. In detach mode it’s not.
	// Opens the devtools.
	OpenDevTools func() `js:"openDevTools"`

	// contents.closeDevTools()
	// Closes the devtools.
	CloseDevTools func() `js:"closeDevTools"`

	// contents.isDevToolsOpened()
	// Returns Boolean - Whether the devtools is opened.

	// contents.isDevToolsFocused()
	// Returns Boolean - Whether the devtools view is focused .

	// contents.toggleDevTools()
	// Toggles the developer tools.

	// contents.inspectElement(x, y)
	// x Integer
	// y Integer
	// Starts inspecting element at position (x, y).

	// contents.inspectServiceWorker()
	// Opens the developer tools for the service worker context.

	// contents.send(channel[, arg1][, arg2][, ...])
	// channel String
	// ...args any[]
	// Send an asynchronous message to renderer process via channel, you can also send arbitrary arguments. Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included.
	Send func(channel string, args ...interface{}) `js:"send"`

	// The renderer process can handle the message by listening to channel with the ipcRenderer module.

	// An example of sending messages from the main process to the renderer process:

	// // In the main process.
	// const {app, BrowserWindow} = require('electron')
	// let win = null

	// app.on('ready', () => {
	//   win = new BrowserWindow({width: 800, height: 600})
	//   win.loadURL(`file://${__dirname}/index.html`)
	//   win.webContents.on('did-finish-load', () => {
	//     win.webContents.send('ping', 'whoooooooh!')
	//   })
	// })
	// <!-- index.html -->
	// <html>
	// <body>
	//   <script>
	//     require('electron').ipcRenderer.on('ping', (event, message) => {
	//       console.log(message)  // Prints 'whoooooooh!'
	//     })
	//   </script>
	// </body>
	// </html>
	// contents.enableDeviceEmulation(parameters)
	// parameters Object
	// screenPosition String - Specify the screen type to emulate (default: desktop)
	// desktop - Desktop screen type
	// mobile - Mobile screen type
	// screenSize Object - Set the emulated screen size (screenPosition == mobile)
	// width Integer - Set the emulated screen width
	// height Integer - Set the emulated screen height
	// viewPosition Object - Position the view on the screen (screenPosition == mobile) (default: {x: 0, y: 0})
	// x Integer - Set the x axis offset from top left corner
	// y Integer - Set the y axis offset from top left corner
	// deviceScaleFactor Integer - Set the device scale factor (if zero defaults to original device scale factor) (default: 0)
	// viewSize Object - Set the emulated view size (empty means no override)
	// width Integer - Set the emulated view width
	// height Integer - Set the emulated view height
	// fitToView Boolean - Whether emulated view should be scaled down if necessary to fit into available space (default: false)
	// offset Object - Offset of the emulated view inside available space (not in fit to view mode) (default: {x: 0, y: 0})
	// x Float - Set the x axis offset from top left corner
	// y Float - Set the y axis offset from top left corner
	// scale Float - Scale of emulated view inside available space (not in fit to view mode) (default: 1)
	// Enable device emulation with the given parameters.

	// contents.disableDeviceEmulation()
	// Disable device emulation enabled by webContents.enableDeviceEmulation.

	// contents.sendInputEvent(event)
	// event Object
	// type String (required) - The type of the event, can be mouseDown, mouseUp, mouseEnter, mouseLeave, contextMenu, mouseWheel, mouseMove, keyDown, keyUp, char.
	// modifiers String[] - An array of modifiers of the event, can include shift, control, alt, meta, isKeypad, isAutoRepeat, leftButtonDown, middleButtonDown, rightButtonDown, capsLock, numLock, left, right.
	// Sends an input event to the page.

	// For keyboard events, the event object also have following properties:

	// keyCode String (required) - The character that will be sent as the keyboard event. Should only use the valid key codes in Accelerator.
	// For mouse events, the event object also have following properties:

	// x Integer (required)
	// y Integer (required)
	// button String - The button pressed, can be left, middle, right
	// globalX Integer
	// globalY Integer
	// movementX Integer
	// movementY Integer
	// clickCount Integer
	// For the mouseWheel event, the event object also have following properties:

	// deltaX Integer
	// deltaY Integer
	// wheelTicksX Integer
	// wheelTicksY Integer
	// accelerationRatioX Integer
	// accelerationRatioY Integer
	// hasPreciseScrollingDeltas Boolean
	// canScroll Boolean
	// contents.beginFrameSubscription([onlyDirty ,]callback)
	// onlyDirty Boolean (optional) - Defaults to false
	// callback Function
	// frameBuffer Buffer
	// dirtyRect Rectangle
	// Begin subscribing for presentation events and captured frames, the callback will be called with callback(frameBuffer, dirtyRect) when there is a presentation event.

	// The frameBuffer is a Buffer that contains raw pixel data. On most machines, the pixel data is effectively stored in 32bit BGRA format, but the actual representation depends on the endianness of the processor (most modern processors are little-endian, on machines with big-endian processors the data is in 32bit ARGB format).

	// The dirtyRect is an object with x, y, width, height properties that describes which part of the page was repainted. If onlyDirty is set to true, frameBuffer will only contain the repainted area. onlyDirty defaults to false.

	// contents.endFrameSubscription()
	// End subscribing for frame presentation events.

	// contents.startDrag(item)
	// item Object
	// file String
	// icon NativeImage
	// Sets the item as dragging item for current drag-drop operation, file is the absolute path of the file to be dragged, and icon is the image showing under the cursor when dragging.

	// contents.savePage(fullPath, saveType, callback)
	// fullPath String - The full file path.
	// saveType String - Specify the save type.
	// HTMLOnly - Save only the HTML of the page.
	// HTMLComplete - Save complete-html page.
	// MHTML - Save complete-html page as MHTML.
	// callback Function - (error) => {}.
	// error Error
	// Returns true if the process of saving page has been initiated successfully.

	// const {BrowserWindow} = require('electron')
	// let win = new BrowserWindow()

	// win.loadURL('https://github.com')

	// win.webContents.on('did-finish-load', () => {
	//   win.webContents.savePage('/tmp/test.html', 'HTMLComplete', (error) => {
	//     if (!error) console.log('Save page successfully')
	//   })
	// })
	// contents.showDefinitionForSelection() macOS
	// Shows pop-up dictionary that searches the selected word on the page.

	// contents.setSize(options)
	// Set the size of the page. This is only supported for <webview> guest contents.

	// options Object
	// normal Object (optional) - Normal size of the page. This can be used in combination with the disableguestresize attribute to manually resize the webview guest contents.
	// width Integer
	// height Integer
	// contents.isOffscreen()
	// Returns Boolean - Indicates whether offscreen rendering is enabled.

	// contents.startPainting()
	// If offscreen rendering is enabled and not painting, start painting.

	// contents.stopPainting()
	// If offscreen rendering is enabled and painting, stop painting.

	// contents.isPainting()
	// Returns Boolean - If offscreen rendering is enabled returns whether it is currently painting.

	// contents.setFrameRate(fps)
	// fps Integer
	// If offscreen rendering is enabled sets the frame rate to the specified number. Only values between 1 and 60 are accepted.

	// contents.getFrameRate()
	// Returns Integer - If offscreen rendering is enabled returns the current frame rate.

	// contents.invalidate()
	// If offscreen rendering is enabled invalidates the frame and generates a new one through the 'paint' event.

	// ======================================================================
	// ======================================================================

	// Instance Properties
	// contents.id
	// A Integer representing the unique ID of this WebContents.
	Id int `js:"id"`

	// contents.session
	// A Session object (session) used by this webContents.
	Session *js.Object `js:"session"`

	// contents.hostWebContents
	// A WebContents instance that might own this WebContents.
	HostWebContents *WebContents `js:"hostWebContents"`

	// contents.devToolsWebContents
	// A WebContents of DevTools for this WebContents.
	DevToolsWebContents *WebContents `js:"devToolsWebContents"`

	// Note: Users should never store this object because it may become null when the DevTools has been closed.

	// contents.debugger
	// A Debugger instance for this webContents.
	Debugger *js.Object `js:"debugger"`
}
