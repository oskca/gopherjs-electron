package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when the document changed its title, calling event.preventDefault() will prevent the native window's title from changing.
	EvtBrowserWindowPageTitleUpdated = "page-title-updated"
	// Emitted when the window is going to be closed. It's emitted before the beforeunload and unload event of the DOM. Calling event.preventDefault() will cancel the close. Usually you would want to use the beforeunload handler to decide whether the window should be closed, which will also be called when the window is reloaded. In Electron, returning any value other than undefined would cancel the close. For example:
	EvtBrowserWindowClose = "close"
	// Emitted when the window is closed. After you have received this event you should remove the reference to the window and avoid using it any more.
	EvtBrowserWindowClosed = "closed"
	// Emitted when the web page becomes unresponsive.
	EvtBrowserWindowUnresponsive = "unresponsive"
	// Emitted when the unresponsive web page becomes responsive again.
	EvtBrowserWindowResponsive = "responsive"
	// Emitted when the window loses focus.
	EvtBrowserWindowBlur = "blur"
	// Emitted when the window gains focus.
	EvtBrowserWindowFocus = "focus"
	// Emitted when the window is shown.
	EvtBrowserWindowShow = "show"
	// Emitted when the window is hidden.
	EvtBrowserWindowHide = "hide"
	// Emitted when the web page has been rendered and window can be displayed without a visual flash.
	EvtBrowserWindowReadyToShow = "ready-to-show"
	// Emitted when window is maximized.
	EvtBrowserWindowMaximize = "maximize"
	// Emitted when the window exits from a maximized state.
	EvtBrowserWindowUnmaximize = "unmaximize"
	// Emitted when the window is minimized.
	EvtBrowserWindowMinimize = "minimize"
	// Emitted when the window is restored from a minimized state.
	EvtBrowserWindowRestore = "restore"
	// Emitted when the window is being resized.
	EvtBrowserWindowResize = "resize"
	// Emitted when the window is being moved to a new position. Note: On macOS this event is just an alias of moved.
	EvtBrowserWindowMove = "move"
	// Emitted once when the window is moved to a new position.
	EvtBrowserWindowMoved = "moved"
	// Emitted when the window enters a full-screen state.
	EvtBrowserWindowEnterFullScreen = "enter-full-screen"
	// Emitted when the window leaves a full-screen state.
	EvtBrowserWindowLeaveFullScreen = "leave-full-screen"
	// Emitted when the window enters a full-screen state triggered by HTML API.
	EvtBrowserWindowEnterHtmlFullScreen = "enter-html-full-screen"
	// Emitted when the window leaves a full-screen state triggered by HTML API.
	EvtBrowserWindowLeaveHtmlFullScreen = "leave-html-full-screen"
	// Emitted when an App Command is invoked. These are typically related to keyboard media keys or browser commands, as well as the "Back" button built into some mice on Windows. Commands are lowercased, underscores are replaced with hyphens, and the APPCOMMAND_ prefix is stripped off. e.g. APPCOMMAND_BROWSER_BACKWARD is emitted as browser-backward.
	EvtBrowserWindowAppCommand = "app-command"
	// Emitted when scroll wheel event phase has begun.
	EvtBrowserWindowScrollTouchBegin = "scroll-touch-begin"
	// Emitted when scroll wheel event phase has ended.
	EvtBrowserWindowScrollTouchEnd = "scroll-touch-end"
	// Emitted when scroll wheel event phase filed upon reaching the edge of element.
	EvtBrowserWindowScrollTouchEdge = "scroll-touch-edge"
	// Emitted on 3-finger swipe. Possible directions are up, right, down, left.
	EvtBrowserWindowSwipe = "swipe"
)

// BrowserWindow version@1.4.15
//
// Create and control browser windows.
type BrowserWindow struct {
	*events.Emitter
	// A WebContents object this window owns. All web page related events and operations will be done via it. See the webContents documentation for its methods and events.
	WebContents *WebContents `js:"webContents"`
	// A Integer representing the unique ID of the window.
	Id int64 `js:"id"`
	// Force closing the window, the unload and beforeunload event won't be emitted for the web page, and close event will also not be emitted for this window, but it guarantees the closed event will be emitted.
	Destroy func() `js:"destroy"`
	// Try to close the window. This has the same effect as a user manually clicking the close button of the window. The web page may cancel the close though. See the close event.
	Close func() `js:"close"`
	// Focuses on the window.
	Focus func() `js:"focus"`
	// Removes focus from the window.
	Blur        func()            `js:"blur"`
	IsFocused   func() (Obj bool) `js:"isFocused"`
	IsDestroyed func() (Obj bool) `js:"isDestroyed"`
	// Shows and gives focus to the window.
	Show func() `js:"show"`
	// Shows the window but doesn't focus on it.
	ShowInactive func() `js:"showInactive"`
	// Hides the window.
	Hide      func()            `js:"hide"`
	IsVisible func() (Obj bool) `js:"isVisible"`
	IsModal   func() (Obj bool) `js:"isModal"`
	// Maximizes the window.
	Maximize func() `js:"maximize"`
	// Unmaximizes the window.
	Unmaximize  func()            `js:"unmaximize"`
	IsMaximized func() (Obj bool) `js:"isMaximized"`
	// Minimizes the window. On some platforms the minimized window will be shown in the Dock.
	Minimize func() `js:"minimize"`
	// Restores the window from minimized state to its previous state.
	Restore     func()            `js:"restore"`
	IsMinimized func() (Obj bool) `js:"isMinimized"`
	// Sets whether the window should be in fullscreen mode.
	SetFullScreen func(Flag bool)   `js:"setFullScreen"`
	IsFullScreen  func() (Obj bool) `js:"isFullScreen"`
	// This will make a window maintain an aspect ratio. The extra size allows a developer to have space, specified in pixels, not included within the aspect ratio calculations. This API already takes into account the difference between a window's size and its content size. Consider a normal window with an HD video player and associated controls. Perhaps there are 15 pixels of controls on the left edge, 25 pixels of controls on the right edge and 50 pixels of controls below the player. In order to maintain a 16:9 aspect ratio (standard aspect ratio for HD @1920x1080) within the player itself we would call this function with arguments of 16/9 and [ 40, 50 ]. The second argument doesn't care where the extra width and height are within the content view--only that they exist. Just sum any extra width and height areas you have within the overall content view.
	SetAspectRatio func(AspectRatio float64, ExtraSize *BrowserWindowSetAspectRatioExtraSize) `js:"setAspectRatio"`
	// Uses Quick Look to preview a file at a given path.
	PreviewFile func(Path string, DisplayName string) `js:"previewFile"`
	// Closes the currently open Quick Look panel.
	CloseFilePreview func() `js:"closeFilePreview"`
	// Resizes and moves the window to the supplied bounds
	SetBounds func(Bounds *js.Object, Animate bool) `js:"setBounds"`
	GetBounds func() (Obj *js.Object)               `js:"getBounds"`
	// Resizes and moves the window's client area (e.g. the web page) to the supplied bounds.
	SetContentBounds func(Bounds *js.Object, Animate bool) `js:"setContentBounds"`
	GetContentBounds func() (Obj *js.Object)               `js:"getContentBounds"`
	// Resizes the window to width and height.
	SetSize func(Width int64, Height int64, Animate bool) `js:"setSize"`
	GetSize func() (Obj *js.Object)                       `js:"getSize"`
	// Resizes the window's client area (e.g. the web page) to width and height.
	SetContentSize func(Width int64, Height int64, Animate bool) `js:"setContentSize"`
	GetContentSize func() (Obj *js.Object)                       `js:"getContentSize"`
	// Sets the minimum size of window to width and height.
	SetMinimumSize func(Width int64, Height int64) `js:"setMinimumSize"`
	GetMinimumSize func() (Obj *js.Object)         `js:"getMinimumSize"`
	// Sets the maximum size of window to width and height.
	SetMaximumSize func(Width int64, Height int64) `js:"setMaximumSize"`
	GetMaximumSize func() (Obj *js.Object)         `js:"getMaximumSize"`
	// Sets whether the window can be manually resized by user.
	SetResizable func(Resizable bool) `js:"setResizable"`
	IsResizable  func() (Obj bool)    `js:"isResizable"`
	// Sets whether the window can be moved by user. On Linux does nothing.
	SetMovable func(Movable bool) `js:"setMovable"`
	// On Linux always returns true.
	IsMovable func() (Obj bool) `js:"isMovable"`
	// Sets whether the window can be manually minimized by user. On Linux does nothing.
	SetMinimizable func(Minimizable bool) `js:"setMinimizable"`
	// On Linux always returns true.
	IsMinimizable func() (Obj bool) `js:"isMinimizable"`
	// Sets whether the window can be manually maximized by user. On Linux does nothing.
	SetMaximizable func(Maximizable bool) `js:"setMaximizable"`
	// On Linux always returns true.
	IsMaximizable func() (Obj bool) `js:"isMaximizable"`
	// Sets whether the maximize/zoom window button toggles fullscreen mode or maximizes the window.
	SetFullScreenable func(Fullscreenable bool) `js:"setFullScreenable"`
	IsFullScreenable  func() (Obj bool)         `js:"isFullScreenable"`
	// Sets whether the window can be manually closed by user. On Linux does nothing.
	SetClosable func(Closable bool) `js:"setClosable"`
	// On Linux always returns true.
	IsClosable func() (Obj bool) `js:"isClosable"`
	// Sets whether the window should show always on top of other windows. After setting this, the window is still a normal window, not a toolbox window which can not be focused on.
	SetAlwaysOnTop func(Flag bool, Level BrowserWindowSetAlwaysOnTopLevel) `js:"setAlwaysOnTop"`
	IsAlwaysOnTop  func() (Obj bool)                                       `js:"isAlwaysOnTop"`
	// Moves window to the center of the screen.
	Center func() `js:"center"`
	// Moves window to x and y.
	SetPosition func(X int64, Y int64, Animate bool) `js:"setPosition"`
	GetPosition func() (Obj *js.Object)              `js:"getPosition"`
	// Changes the title of native window to title.
	SetTitle func(Title string) `js:"setTitle"`
	// Note: The title of web page can be different from the title of the native window.
	GetTitle func() (Obj string) `js:"getTitle"`
	// Changes the attachment point for sheets on macOS. By default, sheets are attached just below the window frame, but you may want to display them beneath a HTML-rendered toolbar. For example:
	SetSheetOffset func(OffsetY float64, OffsetX float64) `js:"setSheetOffset"`
	// Starts or stops flashing the window to attract user's attention.
	FlashFrame func(Flag bool) `js:"flashFrame"`
	// Makes the window not show in the taskbar.
	SetSkipTaskbar func(Skip bool) `js:"setSkipTaskbar"`
	// Enters or leaves the kiosk mode.
	SetKiosk func(Flag bool)   `js:"setKiosk"`
	IsKiosk  func() (Obj bool) `js:"isKiosk"`
	// The native type of the handle is HWND on Windows, NSView* on macOS, and Window (unsigned long) on Linux.
	GetNativeWindowHandle func() (Obj *js.Object) `js:"getNativeWindowHandle"`
	// Hooks a windows message. The callback is called when the message is received in the WndProc.
	HookWindowMessage     func(Message int64, Callback BrowserWindowHookWindowMessageCallback) `js:"hookWindowMessage"`
	IsWindowMessageHooked func(Message int64) (Obj bool)                                       `js:"isWindowMessageHooked"`
	// Unhook the window message.
	UnhookWindowMessage func(Message int64) `js:"unhookWindowMessage"`
	// Unhooks all of the window messages.
	UnhookAllWindowMessages func() `js:"unhookAllWindowMessages"`
	// Sets the pathname of the file the window represents, and the icon of the file will show in window's title bar.
	SetRepresentedFilename func(Filename string) `js:"setRepresentedFilename"`
	GetRepresentedFilename func() (Obj string)   `js:"getRepresentedFilename"`
	// Specifies whether the window’s document has been edited, and the icon in title bar will become gray when set to true.
	SetDocumentEdited func(Edited bool) `js:"setDocumentEdited"`
	IsDocumentEdited  func() (Obj bool) `js:"isDocumentEdited"`
	FocusOnWebView    func()            `js:"focusOnWebView"`
	BlurWebView       func()            `js:"blurWebView"`
	// Same as webContents.capturePage([rect, ]callback).
	CapturePage func(Rect *js.Object, Callback BrowserWindowCapturePageCallback) `js:"capturePage"`
	// Same as webContents.loadURL(url[, options]). The url can be a remote address (e.g. http://) or a path to a local HTML file using the file:// protocol. To ensure that file URLs are properly formatted, it is recommended to use Node's url.format method: You can load a URL using a POST request with URL-encoded data by doing the following:
	LoadURL func(URL string, Options *BrowserWindowLoadURLOptions) `js:"loadURL"`
	// Same as webContents.reload.
	Reload func() `js:"reload"`
	// Sets the menu as the window's menu bar, setting it to null will remove the menu bar.
	SetMenu func(Menu *Menu) `js:"setMenu"`
	// Sets progress value in progress bar. Valid range is [0, 1.0]. Remove progress bar when progress < 0; Change to indeterminate mode when progress > 1. On Linux platform, only supports Unity desktop environment, you need to specify the *.desktop file name to desktopName field in package.json. By default, it will assume app.getName().desktop. On Windows, a mode can be passed. Accepted values are none, normal, indeterminate, error, and paused. If you call setProgressBar without a mode set (but with a value within the valid range), normal will be assumed.
	SetProgressBar func(Progress float64, Options *BrowserWindowSetProgressBarOptions) `js:"setProgressBar"`
	// Sets a 16 x 16 pixel overlay onto the current taskbar icon, usually used to convey some sort of application status or to passively notify the user.
	SetOverlayIcon func(Overlay *NativeImage, Description string) `js:"setOverlayIcon"`
	// Sets whether the window should have a shadow. On Windows and Linux does nothing.
	SetHasShadow func(HasShadow bool) `js:"setHasShadow"`
	// On Windows and Linux always returns true.
	HasShadow func() (Obj bool) `js:"hasShadow"`
	// Add a thumbnail toolbar with a specified set of buttons to the thumbnail image of a window in a taskbar button layout. Returns a Boolean object indicates whether the thumbnail has been added successfully. The number of buttons in thumbnail toolbar should be no greater than 7 due to the limited room. Once you setup the thumbnail toolbar, the toolbar cannot be removed due to the platform's limitation. But you can call the API with an empty array to clean the buttons. The buttons is an array of Button objects: The flags is an array that can include following Strings:
	SetThumbarButtons func(Buttons *js.Object) (Obj bool) `js:"setThumbarButtons"`
	// Sets the region of the window to show as the thumbnail image displayed when hovering over the window in the taskbar. You can reset the thumbnail to be the entire window by specifying an empty region: {x: 0, y: 0, width: 0, height: 0}.
	SetThumbnailClip func(Region *js.Object) `js:"setThumbnailClip"`
	// Sets the toolTip that is displayed when hovering over the window thumbnail in the taskbar.
	SetThumbnailToolTip func(ToolTip string) `js:"setThumbnailToolTip"`
	// Sets the properties for the window's taskbar button. Note: relaunchCommand and relaunchDisplayName must always be set together. If one of those properties is not set, then neither will be used.
	SetAppDetails func(Options *BrowserWindowSetAppDetailsOptions) `js:"setAppDetails"`
	// Same as webContents.showDefinitionForSelection().
	ShowDefinitionForSelection func() `js:"showDefinitionForSelection"`
	// Changes window icon.
	SetIcon func(Icon *NativeImage) `js:"setIcon"`
	// Sets whether the window menu bar should hide itself automatically. Once set the menu bar will only show when users press the single Alt key. If the menu bar is already visible, calling setAutoHideMenuBar(true) won't hide it immediately.
	SetAutoHideMenuBar func(Hide bool)   `js:"setAutoHideMenuBar"`
	IsMenuBarAutoHide  func() (Obj bool) `js:"isMenuBarAutoHide"`
	// Sets whether the menu bar should be visible. If the menu bar is auto-hide, users can still bring up the menu bar by pressing the single Alt key.
	SetMenuBarVisibility func(Visible bool) `js:"setMenuBarVisibility"`
	IsMenuBarVisible     func() (Obj bool)  `js:"isMenuBarVisible"`
	// Sets whether the window should be visible on all workspaces. Note: This API does nothing on Windows.
	SetVisibleOnAllWorkspaces func(Visible bool) `js:"setVisibleOnAllWorkspaces"`
	// Note: This API always returns false on Windows.
	IsVisibleOnAllWorkspaces func() (Obj bool) `js:"isVisibleOnAllWorkspaces"`
	// Makes the window ignore all mouse events. All mouse events happened in this window will be passed to the window below this window, but if this window has focus, it will still receive keyboard events.
	SetIgnoreMouseEvents func(Ignore bool) `js:"setIgnoreMouseEvents"`
	// Prevents the window contents from being captured by other apps. On macOS it sets the NSWindow's sharingType to NSWindowSharingNone. On Windows it calls SetWindowDisplayAffinity with WDA_MONITOR.
	SetContentProtection func(Enable bool) `js:"setContentProtection"`
	// Changes whether the window can be focused.
	SetFocusable func(Focusable bool) `js:"setFocusable"`
	// Sets parent as current window's parent window, passing null will turn current window into a top-level window.
	SetParentWindow func(Parent *BrowserWindow) `js:"setParentWindow"`
	GetParentWindow func() (Obj *BrowserWindow) `js:"getParentWindow"`
	GetChildWindows func() (Obj *js.Object)     `js:"getChildWindows"`
	// Controls whether to hide cursor when typing.
	SetAutoHideCursor func(AutoHide bool) `js:"setAutoHideCursor"`
	// Adds a vibrancy effect to the browser window. Passing null or an empty string will remove the vibrancy effect on the window.
	SetVibrancy func(Type BrowserWindowSetVibrancyType) `js:"setVibrancy"`
}

func WrapBrowserWindow(o *js.Object) *BrowserWindow {
	return &BrowserWindow{
		Emitter: events.New(o),
	}
}

func GetAllWindows() *js.Object {
	o := electron.Get("BrowserWindow")
	ret := o.Call("getAllWindows")
	return ret
}
func GetFocusedWindow() *js.Object {
	o := electron.Get("BrowserWindow")
	ret := o.Call("getFocusedWindow")
	return ret
}
func FromWebContents(WebContents *WebContents) *js.Object {
	o := electron.Get("BrowserWindow")
	ret := o.Call("fromWebContents", WebContents)
	return ret
}
func FromId(Id int64) *js.Object {
	o := electron.Get("BrowserWindow")
	ret := o.Call("fromId", Id)
	return ret
}
func AddDevToolsExtension(Path string) {
	o := electron.Get("BrowserWindow")
	o.Call("addDevToolsExtension", Path)
}
func RemoveDevToolsExtension(Name string) {
	o := electron.Get("BrowserWindow")
	o.Call("removeDevToolsExtension", Name)
}
func GetDevToolsExtensions() *js.Object {
	o := electron.Get("BrowserWindow")
	ret := o.Call("getDevToolsExtensions")
	return ret
}
func NewBrowserWindow(Options *BrowserWindowBrowserWindowOptions) *BrowserWindow {
	o := electron.Get("BrowserWindow")
	ret := o.New(Options)
	return WrapBrowserWindow(ret)
}

type BrowserWindowSetAspectRatioExtraSize struct {
	*js.Object
	Width  int64 `js:"width"`
	Height int64 `js:"height"`
}

type BrowserWindowHookWindowMessageCallback func()
type BrowserWindowCapturePageCallback func(Image *NativeImage)
type BrowserWindowLoadURLOptions struct {
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

type BrowserWindowSetProgressBarOptions struct {
	*js.Object
	// Mode for the progress bar. Can be , , , , or .
	Mode BrowserWindowOptionsMode `js:"mode"`
}

type BrowserWindowOptionsMode string

// consts
const (
	BrowserWindowOptionsModeNone          BrowserWindowOptionsMode = "none"
	BrowserWindowOptionsModeNormal        BrowserWindowOptionsMode = "normal"
	BrowserWindowOptionsModeIndeterminate BrowserWindowOptionsMode = "indeterminate"
	BrowserWindowOptionsModeError         BrowserWindowOptionsMode = "error"
)

type BrowserWindowSetAppDetailsOptions struct {
	*js.Object
	// Window's . It has to be set, otherwise the other options will have no effect.
	AppId string `js:"appId"`
	// Window's .
	AppIconPath string `js:"appIconPath"`
	// Index of the icon in . Ignored when is not set. Default is .
	AppIconIndex int64 `js:"appIconIndex"`
	// Window's .
	RelaunchCommand string `js:"relaunchCommand"`
	// Window's .
	RelaunchDisplayName string `js:"relaunchDisplayName"`
}

type BrowserWindowBrowserWindowOptions struct {
	*js.Object
	// Window's width in pixels. Default is .
	Width int64 `js:"width"`
	// Window's height in pixels. Default is .
	Height int64 `js:"height"`
	// ( if y is used) Window's left offset from screen. Default is to center the window.
	X int64 `js:"x"`
	// ( if x is used) Window's top offset from screen. Default is to center the window.
	Y int64 `js:"y"`
	// The and would be used as web page's size, which means the actual window's size will include window frame's size and be slightly larger. Default is .
	UseContentSize bool `js:"useContentSize"`
	// Show window in the center of the screen.
	Center bool `js:"center"`
	// Window's minimum width. Default is .
	MinWidth int64 `js:"minWidth"`
	// Window's minimum height. Default is .
	MinHeight int64 `js:"minHeight"`
	// Window's maximum width. Default is no limit.
	MaxWidth int64 `js:"maxWidth"`
	// Window's maximum height. Default is no limit.
	MaxHeight int64 `js:"maxHeight"`
	// Whether window is resizable. Default is .
	Resizable bool `js:"resizable"`
	// Whether window is movable. This is not implemented on Linux. Default is .
	Movable bool `js:"movable"`
	// Whether window is minimizable. This is not implemented on Linux. Default is .
	Minimizable bool `js:"minimizable"`
	// Whether window is maximizable. This is not implemented on Linux. Default is .
	Maximizable bool `js:"maximizable"`
	// Whether window is closable. This is not implemented on Linux. Default is .
	Closable bool `js:"closable"`
	// Whether the window can be focused. Default is . On Windows setting also implies setting . On Linux setting makes the window stop interacting with wm, so the window will always stay on top in all workspaces.
	Focusable bool `js:"focusable"`
	// Whether the window should always stay on top of other windows. Default is .
	AlwaysOnTop bool `js:"alwaysOnTop"`
	// Whether the window should show in fullscreen. When explicitly set to the fullscreen button will be hidden or disabled on macOS. Default is .
	Fullscreen bool `js:"fullscreen"`
	// Whether the window can be put into fullscreen mode. On macOS, also whether the maximize/zoom button should toggle full screen mode or maximize window. Default is .
	Fullscreenable bool `js:"fullscreenable"`
	// Whether to show the window in taskbar. Default is .
	SkipTaskbar bool `js:"skipTaskbar"`
	// The kiosk mode. Default is .
	Kiosk bool `js:"kiosk"`
	// Default window title. Default is .
	Title string `js:"title"`
	// The window icon. On Windows it is recommended to use icons to get best visual effects, you can also leave it undefined so the executable's icon will be used.
	Icon *NativeImage `js:"icon"`
	// Whether window should be shown when created. Default is .
	Show bool `js:"show"`
	// Specify to create a . Default is .
	Frame bool `js:"frame"`
	// Specify parent window. Default is .
	Parent *BrowserWindow `js:"parent"`
	// Whether this is a modal window. This only works when the window is a child window. Default is .
	Modal bool `js:"modal"`
	// Whether the web view accepts a single mouse-down event that simultaneously activates the window. Default is .
	AcceptFirstMouse bool `js:"acceptFirstMouse"`
	// Whether to hide cursor when typing. Default is .
	DisableAutoHideCursor bool `js:"disableAutoHideCursor"`
	// Auto hide the menu bar unless the key is pressed. Default is .
	AutoHideMenuBar bool `js:"autoHideMenuBar"`
	// Enable the window to be resized larger than screen. Default is .
	EnableLargerThanScreen bool `js:"enableLargerThanScreen"`
	// Window's background color as Hexadecimal value, like or or (alpha is supported). Default is (white).
	BackgroundColor string `js:"backgroundColor"`
	// Whether window should have a shadow. This is only implemented on macOS. Default is .
	HasShadow bool `js:"hasShadow"`
	// Forces using dark theme for the window, only works on some GTK+3 desktop environments. Default is .
	DarkTheme bool `js:"darkTheme"`
	// Makes the window . Default is .
	Transparent bool `js:"transparent"`
	// The type of window, default is normal window. See more about this below.
	Type string `js:"type"`
	// The style of window title bar. Default is . Possible values are:
	TitleBarStyle BrowserWindowOptionsTitleBarStyle `js:"titleBarStyle"`
	// Use style for frameless windows on Windows, which adds standard window frame. Setting it to will remove window shadow and window animations. Default is .
	ThickFrame bool `js:"thickFrame"`
	// Add a type of vibrancy effect to the window, only on macOS. Can be , , , , , , , , or .
	Vibrancy BrowserWindowOptionsVibrancy `js:"vibrancy"`
	// Controls the behavior on macOS when option-clicking the green stoplight button on the toolbar or by clicking the Window > Zoom menu item. If , the window will grow to the preferred width of the web page when zoomed, will cause it to zoom to the width of the screen. This will also affect the behavior when calling directly. Default is .
	ZoomToPageWidth bool `js:"zoomToPageWidth"`
	// Settings of web page's features.
	WebPreferences *BrowserWindowOptionsWebPreferences `js:"webPreferences"`
}

type BrowserWindowOptionsWebPreferences struct {
	*js.Object
	// Whether to enable DevTools. If it is set to , can not use to open DevTools. Default is .
	DevTools bool `js:"devTools"`
	// Whether node integration is enabled. Default is .
	NodeIntegration bool `js:"nodeIntegration"`
	// Specifies a script that will be loaded before other scripts run in the page. This script will always have access to node APIs no matter whether node integration is turned on or off. The value should be the absolute file path to the script. When node integration is turned off, the preload script can reintroduce Node global symbols back to the global scope. See example .
	Preload string `js:"preload"`
	// Sets the session used by the page. Instead of passing the Session object directly, you can also choose to use the option instead, which accepts a partition string. When both and are provided, will be preferred. Default is the default session.
	Session *Session `js:"session"`
	// Sets the session used by the page according to the session's partition string. If starts with , the page will use a persistent session available to all pages in the app with the same . If there is no prefix, the page will use an in-memory session. By assigning the same , multiple pages can share the same session. Default is the default session.
	Partition string `js:"partition"`
	// The default zoom factor of the page, represents . Default is .
	ZoomFactor float64 `js:"zoomFactor"`
	// Enables JavaScript support. Default is .
	Javascript bool `js:"javascript"`
	// When , it will disable the same-origin policy (usually using testing websites by people), and set and to if these two options are not set by user. Default is .
	WebSecurity bool `js:"webSecurity"`
	// Allow an https page to display content like images from http URLs. Default is .
	AllowDisplayingInsecureContent bool `js:"allowDisplayingInsecureContent"`
	// Allow an https page to run JavaScript, CSS or plugins from http URLs. Default is .
	AllowRunningInsecureContent bool `js:"allowRunningInsecureContent"`
	// Enables image support. Default is .
	Images bool `js:"images"`
	// Make TextArea elements resizable. Default is .
	TextAreasAreResizable bool `js:"textAreasAreResizable"`
	// Enables WebGL support. Default is .
	Webgl bool `js:"webgl"`
	// Enables WebAudio support. Default is .
	Webaudio bool `js:"webaudio"`
	// Whether plugins should be enabled. Default is .
	Plugins bool `js:"plugins"`
	// Enables Chromium's experimental features. Default is .
	ExperimentalFeatures bool `js:"experimentalFeatures"`
	// Enables Chromium's experimental canvas features. Default is .
	ExperimentalCanvasFeatures bool `js:"experimentalCanvasFeatures"`
	// Enables scroll bounce (rubber banding) effect on macOS. Default is .
	ScrollBounce bool `js:"scrollBounce"`
	// A list of feature strings separated by , like to enable. The full list of supported feature strings can be found in the file.
	BlinkFeatures string `js:"blinkFeatures"`
	// A list of feature strings separated by , like to disable. The full list of supported feature strings can be found in the file.
	DisableBlinkFeatures string `js:"disableBlinkFeatures"`
	// Sets the default font for the font-family.
	DefaultFontFamily *BrowserWindowWebPreferencesDefaultFontFamily `js:"defaultFontFamily"`
	// Defaults to .
	DefaultFontSize int64 `js:"defaultFontSize"`
	// Defaults to .
	DefaultMonospaceFontSize int64 `js:"defaultMonospaceFontSize"`
	// Defaults to .
	MinimumFontSize int64 `js:"minimumFontSize"`
	// Defaults to .
	DefaultEncoding string `js:"defaultEncoding"`
	// Whether to throttle animations and timers when the page becomes background. Defaults to .
	BackgroundThrottling bool `js:"backgroundThrottling"`
	// Whether to enable offscreen rendering for the browser window. Defaults to . See the for more details.
	Offscreen bool `js:"offscreen"`
	// Whether to enable Chromium OS-level sandbox.
	Sandbox bool `js:"sandbox"`
	// Whether to run Electron APIs and the specified script in a separate JavaScript context. Defaults to . The context that the script runs in will still have full access to the and globals but it will use its own set of JavaScript builtins (, , , etc.) and will be isolated from any changes made to the global environment by the loaded page. The Electron API will only be available in the script and not the loaded page. This option should be used when loading potentially untrusted remote content to ensure the loaded content cannot tamper with the script and any Electron APIs being used. This option uses the same technique used by . You can access this context in the dev tools by selecting the 'Electron Isolated Context' entry in the combo box at the top of the Console tab. This option is currently experimental and may change or be removed in future Electron releases.
	ContextIsolation bool `js:"contextIsolation"`
}

type BrowserWindowWebPreferencesDefaultFontFamily struct {
	*js.Object
	// Defaults to .
	Standard string `js:"standard"`
	// Defaults to .
	Serif string `js:"serif"`
	// Defaults to .
	SansSerif string `js:"sansSerif"`
	// Defaults to .
	Monospace string `js:"monospace"`
	// Defaults to .
	Cursive string `js:"cursive"`
	// Defaults to .
	Fantasy string `js:"fantasy"`
}

type BrowserWindowOptionsTitleBarStyle string

// consts
const (
	BrowserWindowOptionsTitleBarStyleDefault     BrowserWindowOptionsTitleBarStyle = "default"
	BrowserWindowOptionsTitleBarStyleHidden      BrowserWindowOptionsTitleBarStyle = "hidden"
	BrowserWindowOptionsTitleBarStyleHiddenInset BrowserWindowOptionsTitleBarStyle = "hidden-inset"
)

type BrowserWindowOptionsVibrancy string

// consts
const (
	BrowserWindowOptionsVibrancyAppearanceBased BrowserWindowOptionsVibrancy = "appearance-based"
	BrowserWindowOptionsVibrancyLight           BrowserWindowOptionsVibrancy = "light"
	BrowserWindowOptionsVibrancyDark            BrowserWindowOptionsVibrancy = "dark"
	BrowserWindowOptionsVibrancyTitlebar        BrowserWindowOptionsVibrancy = "titlebar"
	BrowserWindowOptionsVibrancySelection       BrowserWindowOptionsVibrancy = "selection"
	BrowserWindowOptionsVibrancyMenu            BrowserWindowOptionsVibrancy = "menu"
	BrowserWindowOptionsVibrancyPopover         BrowserWindowOptionsVibrancy = "popover"
	BrowserWindowOptionsVibrancySidebar         BrowserWindowOptionsVibrancy = "sidebar"
	BrowserWindowOptionsVibrancyMediumLight     BrowserWindowOptionsVibrancy = "medium-light"
	BrowserWindowOptionsVibrancyUltraDark       BrowserWindowOptionsVibrancy = "ultra-dark"
)

type BrowserWindowSetVibrancyType string

// consts
const (
	BrowserWindowSetVibrancyTypeAppearanceBased BrowserWindowSetVibrancyType = "appearance-based"
	BrowserWindowSetVibrancyTypeLight           BrowserWindowSetVibrancyType = "light"
	BrowserWindowSetVibrancyTypeDark            BrowserWindowSetVibrancyType = "dark"
	BrowserWindowSetVibrancyTypeTitlebar        BrowserWindowSetVibrancyType = "titlebar"
	BrowserWindowSetVibrancyTypeSelection       BrowserWindowSetVibrancyType = "selection"
	BrowserWindowSetVibrancyTypeMenu            BrowserWindowSetVibrancyType = "menu"
	BrowserWindowSetVibrancyTypePopover         BrowserWindowSetVibrancyType = "popover"
	BrowserWindowSetVibrancyTypeSidebar         BrowserWindowSetVibrancyType = "sidebar"
	BrowserWindowSetVibrancyTypeMediumLight     BrowserWindowSetVibrancyType = "medium-light"
	BrowserWindowSetVibrancyTypeUltraDark       BrowserWindowSetVibrancyType = "ultra-dark"
)

type BrowserWindowSetAlwaysOnTopLevel string

// consts
const (
	BrowserWindowSetAlwaysOnTopLevelNormal      BrowserWindowSetAlwaysOnTopLevel = "normal"
	BrowserWindowSetAlwaysOnTopLevelFloating    BrowserWindowSetAlwaysOnTopLevel = "floating"
	BrowserWindowSetAlwaysOnTopLevelTornOffMenu BrowserWindowSetAlwaysOnTopLevel = "torn-off-menu"
	BrowserWindowSetAlwaysOnTopLevelModalPanel  BrowserWindowSetAlwaysOnTopLevel = "modal-panel"
	BrowserWindowSetAlwaysOnTopLevelMainMenu    BrowserWindowSetAlwaysOnTopLevel = "main-menu"
	BrowserWindowSetAlwaysOnTopLevelStatus      BrowserWindowSetAlwaysOnTopLevel = "status"
	BrowserWindowSetAlwaysOnTopLevelPopUpMenu   BrowserWindowSetAlwaysOnTopLevel = "pop-up-menu"
	BrowserWindowSetAlwaysOnTopLevelScreenSaver BrowserWindowSetAlwaysOnTopLevel = "screen-saver"
)
