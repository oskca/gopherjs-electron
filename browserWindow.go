package electron

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	browserWindow = electron.Get("BrowserWindow")
)

const (
	// Instance Events
	// Objects created with new BrowserWindow emit the following events:

	// Note: Some events are only available on specific operating systems and are labeled as such.

	// Event: ‘page-title-updated’
	// Returns:
	//
	// event Event
	// title String
	// Emitted when the document changed its title, calling event.preventDefault() will prevent the native window’s title from changing.
	BwEvtPageTitleUpdated = "page-title-updated"

	// Event: ‘close’
	// Returns:
	//
	// event Event
	// Emitted when the window is going to be closed. It’s emitted before the beforeunload and unload event of the DOM. Calling event.preventDefault() will cancel the close.
	BwEvtClose = "close"

// Usually you would want to use the beforeunload handler to decide whether the window should be closed, which will also be called when the window is reloaded. In Electron, returning any value other than undefined would cancel the close. For example:

// window.onbeforeunload = (e) => {
//   console.log('I do not want to be closed')

//   // Unlike usual browsers that a message box will be prompted to users, returning
//   // a non-void value will silently cancel the close.
//   // It is recommended to use the dialog API to let the user confirm closing the
//   // application.
//   e.returnValue = false
// }
// Event: ‘closed’
// Emitted when the window is closed. After you have received this event you should remove the reference to the window and avoid using it any more.

// Event: ‘unresponsive’
// Emitted when the web page becomes unresponsive.

// Event: ‘responsive’
// Emitted when the unresponsive web page becomes responsive again.

// Event: ‘blur’
// Emitted when the window loses focus.

// Event: ‘focus’
// Emitted when the window gains focus.

// Event: ‘show’
// Emitted when the window is shown.

// Event: ‘hide’
// Emitted when the window is hidden.

// Event: ‘ready-to-show’
// Emitted when the web page has been rendered and window can be displayed without a visual flash.

// Event: ‘maximize’
// Emitted when window is maximized.

// Event: ‘unmaximize’
// Emitted when the window exits from a maximized state.

// Event: ‘minimize’
// Emitted when the window is minimized.

// Event: ‘restore’
// Emitted when the window is restored from a minimized state.

// Event: ‘resize’
// Emitted when the window is being resized.

// Event: ‘move’
// Emitted when the window is being moved to a new position.

// Note: On macOS this event is just an alias of moved.

// Event: ‘moved’ macOS
// Emitted once when the window is moved to a new position.

// Event: ‘enter-full-screen’
// Emitted when the window enters a full-screen state.

// Event: ‘leave-full-screen’
// Emitted when the window leaves a full-screen state.

// Event: ‘enter-html-full-screen’
// Emitted when the window enters a full-screen state triggered by HTML API.

// Event: ‘leave-html-full-screen’
// Emitted when the window leaves a full-screen state triggered by HTML API.

// Event: ‘app-command’ Windows
// Returns:

// event Event
// command String
// Emitted when an App Command is invoked. These are typically related to keyboard media keys or browser commands, as well as the “Back” button built into some mice on Windows.

// Commands are lowercased, underscores are replaced with hyphens, and the APPCOMMAND_ prefix is stripped off. e.g. APPCOMMAND_BROWSER_BACKWARD is emitted as browser-backward.

// const {BrowserWindow} = require('electron')
// let win = new BrowserWindow()
// win.on('app-command', (e, cmd) => {
//   // Navigate the window back when the user hits their mouse back button
//   if (cmd === 'browser-backward' && win.webContents.canGoBack()) {
//     win.webContents.goBack()
//   }
// })
// Event: ‘scroll-touch-begin’ macOS
// Emitted when scroll wheel event phase has begun.

// Event: ‘scroll-touch-end’ macOS
// Emitted when scroll wheel event phase has ended.

// Event: ‘scroll-touch-edge’ macOS
// Emitted when scroll wheel event phase filed upon reaching the edge of element.

// Event: ‘swipe’ macOS
// Returns:

// event Event
// direction String
// Emitted on 3-finger swipe. Possible directions are up, right, down, left.
)

type BrowserWindowOptions struct {
	*js.Object
	Width                  int         `js:"width"`          //Window's width in pixels. Default is 800
	Height                 int         `js:"height"`         //Window's height in pixels. Default is 600
	X                      int         `js:"x"`              //Window's left offset from screen. Default is to center the window.
	Y                      int         `js:"y"`              //Window's top offset from screen. Default is to center the window.
	UseContentSize         bool        `js:"useContentSize"` //The width and height would be used as web page's size, which means the actual window's size will include window frame's size and be slightly larger. Default is false.
	Center                 bool        `js:"center"`         //Show window in the center of the screen.
	MinWidth               int         `js:"minWidth"`       //Window's minimum width. Default is 0.
	MinHeight              int         `js:"minHeight"`      //Window's minimum height. Default is 0.
	MaxWidth               int         `js:"maxWidth"`       //Window's maximum width. Default is no limit.
	MaxHeight              int         `js:"maxHeight"`      //Window's maximum height. Default is no limit.
	Resizable              bool        `js:"resizable"`      //Whether window is resizable. Default is true.
	AlwaysOnTop            bool        `js:"alwaysOnTop"`    //Whether the window should always stay on top of other windows. Default is false.
	Fullscreen             bool        `js:"fullscreen"`     //Whether the window should show in fullscreen. When set to false the fullscreen button will be hidden or disabled on OS X. Default is false.
	SkipTaskbar            bool        `js:"skipTaskbar"`    //Whether to show the window in taskbar. Default is false.
	Kiosk                  bool        `js:"kiosk"`          //The kiosk mode. Default is false.
	Title                  string      `js:"title"`          //Default window title. Default is "Electron".
	Icon                   interface{} `js:"icon"`           //The window icon, when omitted on Windows the executable's icon would be used as window icon.
	Show                   bool        `js:"show"`           //Whether window should be shown when created. Default is true.
	Frame                  bool        `js:"frame"`          //Specify false to create a Frameless Window. Default is true
	AcceptFirstMouse       bool        `js:"acceptFirstMouse"`
	DisableAutoHideCursor  bool        `js:"disableAutoHideCursor"`
	AutoHideMenuBar        bool        `js:"autoHideMenuBar"`
	EnableLargerThanScreen bool        `js:"enableLargerThanScreen"`
	BackgroundColor        string      `js:"backgroundColor"`
	DarkTheme              bool        `js:"darkTheme"`
	Transparent            bool        `js:"transparent"`
	Type                   string      `js:"type"`
	TitleBarStyle          string      `js:"titleBarStyle"`
	WebPreferences         interface{} `js:"webPreferences"`
	NodeIntegration        bool        `js:"nodeIntegration"`
}

type BrowserWindow struct {
	*js.Object
	// win.webContents
	// A WebContents object this window owns. All web page related events and operations will be done via it.
	// See the webContents documentation for its methods and events.
	WebContents *js.Object `js:"webContents"`

	// win.id
	// A Integer representing the unique ID of the window.
	Id *js.Object `js:"id"`

	// Instance Methods
	// Objects created with new BrowserWindow have the following instance methods:

	// Note: Some methods are only available on specific operating systems and are labeled as such.

	// win.destroy()
	// Force closing the window, the unload and beforeunload event won’t be emitted for the web page, and close event will also not be emitted for this window, but it guarantees the closed event will be emitted.
	Destroy func() `js:"destroy"`

	// win.close()
	// Try to close the window. This has the same effect as a user manually clicking the close button of the window. The web page may cancel the close though. See the close event.
	Close func() `js:"close"`

	// win.focus()
	// Focuses on the window.

	// win.blur()
	// Removes focus from the window.

	// win.isFocused()
	// Returns Boolean - Whether the window is focused.

	// win.isDestroyed()
	// Returns Boolean - Whether the window is destroyed.

	// win.show()
	// Shows and gives focus to the window.

	// win.showInactive()
	// Shows the window but doesn’t focus on it.

	// win.hide()
	// Hides the window.

	// win.isVisible()
	// Returns Boolean - Whether the window is visible to the user.

	// win.isModal()
	// Returns Boolean - Whether current window is a modal window.

	// win.maximize()
	// Maximizes the window.

	// win.unmaximize()
	// Unmaximizes the window.

	// win.isMaximized()
	// Returns Boolean - Whether the window is maximized.

	// win.minimize()
	// Minimizes the window. On some platforms the minimized window will be shown in the Dock.

	// win.restore()
	// Restores the window from minimized state to its previous state.

	// win.isMinimized()
	// Returns Boolean - Whether the window is minimized.

	// win.setFullScreen(flag)
	// flag Boolean
	// Sets whether the window should be in fullscreen mode.

	// win.isFullScreen()
	// Returns Boolean - Whether the window is in fullscreen mode.

	// win.setAspectRatio(aspectRatio[, extraSize]) macOS
	// aspectRatio Float - The aspect ratio to maintain for some portion of the content view.
	// extraSize Object (optional) - The extra size not to be included while maintaining the aspect ratio.
	// width Integer
	// height Integer
	// This will make a window maintain an aspect ratio. The extra size allows a developer to have space, specified in pixels, not included within the aspect ratio calculations. This API already takes into account the difference between a window’s size and its content size.

	// Consider a normal window with an HD video player and associated controls. Perhaps there are 15 pixels of controls on the left edge, 25 pixels of controls on the right edge and 50 pixels of controls below the player. In order to maintain a 16:9 aspect ratio (standard aspect ratio for HD @1920x1080) within the player itself we would call this function with arguments of 16/9 and [ 40, 50 ]. The second argument doesn’t care where the extra width and height are within the content view–only that they exist. Just sum any extra width and height areas you have within the overall content view.

	// win.previewFile(path[, displayName]) macOS
	// path String - The absolute path to the file to preview with QuickLook. This is important as Quick Look uses the file name and file extension on the path to determine the content type of the file to open.
	// displayName String (Optional) - The name of the file to display on the Quick Look modal view. This is purely visual and does not affect the content type of the file. Defaults to path.
	// Uses Quick Look to preview a file at a given path.

	// win.closeFilePreview() macOS
	// Closes the currently open Quick Look panel.

	// win.setBounds(bounds[, animate])
	// bounds Rectangle
	// animate Boolean (optional) macOS
	// Resizes and moves the window to the supplied bounds

	// win.getBounds()
	// Returns Rectangle

	// win.setContentBounds(bounds[, animate])
	// bounds Rectangle
	// animate Boolean (optional) macOS
	// Resizes and moves the window’s client area (e.g. the web page) to the supplied bounds.

	// win.getContentBounds()
	// Returns Rectangle

	// win.setSize(width, height[, animate])
	// width Integer
	// height Integer
	// animate Boolean (optional) macOS
	// Resizes the window to width and height.

	// win.getSize()
	// Returns Integer[] - Contains the window’s width and height.

	// win.setContentSize(width, height[, animate])
	// width Integer
	// height Integer
	// animate Boolean (optional) macOS
	// Resizes the window’s client area (e.g. the web page) to width and height.

	// win.getContentSize()
	// Returns Integer[] - Contains the window’s client area’s width and height.

	// win.setMinimumSize(width, height)
	// width Integer
	// height Integer
	// Sets the minimum size of window to width and height.

	// win.getMinimumSize()
	// Returns Integer[] - Contains the window’s minimum width and height.

	// win.setMaximumSize(width, height)
	// width Integer
	// height Integer
	// Sets the maximum size of window to width and height.

	// win.getMaximumSize()
	// Returns Integer[] - Contains the window’s maximum width and height.

	// win.setResizable(resizable)
	// resizable Boolean
	// Sets whether the window can be manually resized by user.

	// win.isResizable()
	// Returns Boolean - Whether the window can be manually resized by user.

	// win.setMovable(movable) macOS Windows
	// movable Boolean
	// Sets whether the window can be moved by user. On Linux does nothing.

	// win.isMovable() macOS Windows
	// Returns Boolean - Whether the window can be moved by user.

	// On Linux always returns true.

	// win.setMinimizable(minimizable) macOS Windows
	// minimizable Boolean
	// Sets whether the window can be manually minimized by user. On Linux does nothing.

	// win.isMinimizable() macOS Windows
	// Returns Boolean - Whether the window can be manually minimized by user

	// On Linux always returns true.

	// win.setMaximizable(maximizable) macOS Windows
	// maximizable Boolean
	// Sets whether the window can be manually maximized by user. On Linux does nothing.

	// win.isMaximizable() macOS Windows
	// Returns Boolean - Whether the window can be manually maximized by user.

	// On Linux always returns true.

	// win.setFullScreenable(fullscreenable)
	// fullscreenable Boolean
	// Sets whether the maximize/zoom window button toggles fullscreen mode or maximizes the window.

	// win.isFullScreenable()
	// Returns Boolean - Whether the maximize/zoom window button toggles fullscreen mode or maximizes the window.

	// win.setClosable(closable) macOS Windows
	// closable Boolean
	// Sets whether the window can be manually closed by user. On Linux does nothing.

	// win.isClosable() macOS Windows
	// Returns Boolean - Whether the window can be manually closed by user.

	// On Linux always returns true.

	// win.setAlwaysOnTop(flag[, level])
	// flag Boolean
	// level String (optional) macOS - Values include normal, floating, torn-off-menu, modal-panel, main-menu, status, pop-up-menu, screen-saver, and dock (Deprecated). The default is floating. See the macOS docs for more details.
	// Sets whether the window should show always on top of other windows. After setting this, the window is still a normal window, not a toolbox window which can not be focused on.

	// win.isAlwaysOnTop()
	// Returns Boolean - Whether the window is always on top of other windows.

	// win.center()
	// Moves window to the center of the screen.

	// win.setPosition(x, y[, animate])
	// x Integer
	// y Integer
	// animate Boolean (optional) macOS
	// Moves window to x and y.

	// win.getPosition()
	// Returns Integer[] - Contains the window’s current position.

	// win.setTitle(title)
	// title String
	// Changes the title of native window to title.

	// win.getTitle()
	// Returns String - The title of the native window.

	// Note: The title of web page can be different from the title of the native window.

	// win.setSheetOffset(offsetY[, offsetX]) macOS
	// offsetY Float
	// offsetX Float (optional)
	// Changes the attachment point for sheets on macOS. By default, sheets are attached just below the window frame, but you may want to display them beneath a HTML-rendered toolbar. For example:

	// const {BrowserWindow} = require('electron')
	// let win = new BrowserWindow()

	// let toolbarRect = document.getElementById('toolbar').getBoundingClientRect()
	// win.setSheetOffset(toolbarRect.height)
	// win.flashFrame(flag)
	// flag Boolean
	// Starts or stops flashing the window to attract user’s attention.

	// win.setSkipTaskbar(skip)
	// skip Boolean
	// Makes the window not show in the taskbar.

	// win.setKiosk(flag)
	// flag Boolean
	// Enters or leaves the kiosk mode.

	// win.isKiosk()
	// Returns Boolean - Whether the window is in kiosk mode.

	// win.getNativeWindowHandle()
	// Returns Buffer - The platform-specific handle of the window.

	// The native type of the handle is HWND on Windows, NSView* on macOS, and Window (unsigned long) on Linux.

	// win.hookWindowMessage(message, callback) Windows
	// message Integer
	// callback Function
	// Hooks a windows message. The callback is called when the message is received in the WndProc.

	// win.isWindowMessageHooked(message) Windows
	// message Integer
	// Returns Boolean - true or false depending on whether the message is hooked.

	// win.unhookWindowMessage(message) Windows
	// message Integer
	// Unhook the window message.

	// win.unhookAllWindowMessages() Windows
	// Unhooks all of the window messages.

	// win.setRepresentedFilename(filename) macOS
	// filename String
	// Sets the pathname of the file the window represents, and the icon of the file will show in window’s title bar.

	// win.getRepresentedFilename() macOS
	// Returns String - The pathname of the file the window represents.

	// win.setDocumentEdited(edited) macOS
	// edited Boolean
	// Specifies whether the window’s document has been edited, and the icon in title bar will become gray when set to true.

	// win.isDocumentEdited() macOS
	// Returns Boolean - Whether the window’s document has been edited.

	// win.focusOnWebView()
	// win.blurWebView()
	// win.capturePage([rect, ]callback)
	// rect Rectangle (optional) - The bounds to capture
	// callback Function
	// image NativeImage
	// Same as webContents.capturePage([rect, ]callback).

	// win.loadURL(url[, options])
	// url String
	// options Object (optional)
	// httpReferrer String (optional) - A HTTP Referrer url.
	// userAgent String (optional) - A user agent originating the request.
	// extraHeaders String (optional) - Extra headers separated by “\n”
	// postData (UploadRawData | UploadFile | UploadFileSystem | UploadBlob)[] - (optional)
	// Same as webContents.loadURL(url[, options]).
	//
	// The url can be a remote address (e.g. http://) or a path to a local HTML file using the file:// protocol.
	//
	// To ensure that file URLs are properly formatted, it is recommended to use Node’s url.format method:
	//
	// let url = require('url').format({
	//   protocol: 'file',
	//   slashes: true,
	//   pathname: require('path').join(__dirname, 'index.html')
	// })
	//
	// win.loadURL(url)
	// You can load a URL using a POST request with URL-encoded data by doing the following:
	//
	// win.loadURL('http://localhost:8000/post', {
	//   postData: [{
	//     type: 'rawData',
	//     bytes: Buffer.from('hello=world')
	//   }],
	//   extraHeaders: 'Content-Type: application/x-www-form-urlencoded'
	// })
	// win.reload()
	// Same as webContents.reload.
	LoadURL   func(url string)                  `js:"loadURL"`
	LoadURLEx func(url string, opts *js.Object) `js:"loadURL"`

	// win.setMenu(menu) Linux Windows
	// menu Menu
	// Sets the menu as the window’s menu bar, setting it to null will remove the menu bar.

	// win.setProgressBar(progress[, options])
	// progress Double
	// options Object (optional)
	// mode String Windows - Mode for the progress bar. Can be none, normal, indeterminate, error, or paused.
	// Sets progress value in progress bar. Valid range is [0, 1.0].

	// Remove progress bar when progress < 0; Change to indeterminate mode when progress > 1.

	// On Linux platform, only supports Unity desktop environment, you need to specify the *.desktop file name to desktopName field in package.json. By default, it will assume app.getName().desktop.

	// On Windows, a mode can be passed. Accepted values are none, normal, indeterminate, error, and paused. If you call setProgressBar without a mode set (but with a value within the valid range), normal will be assumed.

	// win.setOverlayIcon(overlay, description) Windows
	// overlay NativeImage - the icon to display on the bottom right corner of the taskbar icon. If this parameter is null, the overlay is cleared
	// description String - a description that will be provided to Accessibility screen readers
	// Sets a 16 x 16 pixel overlay onto the current taskbar icon, usually used to convey some sort of application status or to passively notify the user.

	// win.setHasShadow(hasShadow) macOS
	// hasShadow Boolean
	// Sets whether the window should have a shadow. On Windows and Linux does nothing.

	// win.hasShadow() macOS
	// Returns Boolean - Whether the window has a shadow.

	// On Windows and Linux always returns true.

	// win.setThumbarButtons(buttons) Windows
	// buttons ThumbarButton[]
	// Returns Boolean - Whether the buttons were added successfully

	// Add a thumbnail toolbar with a specified set of buttons to the thumbnail image of a window in a taskbar button layout. Returns a Boolean object indicates whether the thumbnail has been added successfully.

	// The number of buttons in thumbnail toolbar should be no greater than 7 due to the limited room. Once you setup the thumbnail toolbar, the toolbar cannot be removed due to the platform’s limitation. But you can call the API with an empty array to clean the buttons.

	// The buttons is an array of Button objects:

	// Button Object
	// icon NativeImage - The icon showing in thumbnail toolbar.
	// click Function
	// tooltip String (optional) - The text of the button’s tooltip.
	// flags String[] (optional) - Control specific states and behaviors of the button. By default, it is ['enabled'].
	// The flags is an array that can include following Strings:

	// enabled - The button is active and available to the user.
	// disabled - The button is disabled. It is present, but has a visual state indicating it will not respond to user action.
	// dismissonclick - When the button is clicked, the thumbnail window closes immediately.
	// nobackground - Do not draw a button border, use only the image.
	// hidden - The button is not shown to the user.
	// noninteractive - The button is enabled but not interactive; no pressed button state is drawn. This value is intended for instances where the button is used in a notification.
	// win.setThumbnailClip(region) Windows
	// region Rectangle - Region of the window
	// Sets the region of the window to show as the thumbnail image displayed when hovering over the window in the taskbar. You can reset the thumbnail to be the entire window by specifying an empty region: {x: 0, y: 0, width: 0, height: 0}.

	// win.setThumbnailToolTip(toolTip) Windows
	// toolTip String
	// Sets the toolTip that is displayed when hovering over the window thumbnail in the taskbar.

	// win.setAppDetails(options) Windows
	// options Object
	// appId String (optional) - Window’s App User Model ID. It has to be set, otherwise the other options will have no effect.
	// appIconPath String (optional) - Window’s Relaunch Icon.
	// appIconIndex Integer (optional) - Index of the icon in appIconPath. Ignored when appIconPath is not set. Default is 0.
	// relaunchCommand String (optional) - Window’s Relaunch Command.
	// relaunchDisplayName String (optional) - Window’s Relaunch Display Name.
	// Sets the properties for the window’s taskbar button.

	// Note: relaunchCommand and relaunchDisplayName must always be set together. If one of those properties is not set, then neither will be used.

	// win.showDefinitionForSelection() macOS
	// Same as webContents.showDefinitionForSelection().

	// win.setIcon(icon) Windows Linux
	// icon NativeImage
	// Changes window icon.

	// win.setAutoHideMenuBar(hide)
	// hide Boolean
	// Sets whether the window menu bar should hide itself automatically. Once set the menu bar will only show when users press the single Alt key.

	// If the menu bar is already visible, calling setAutoHideMenuBar(true) won’t hide it immediately.

	// win.isMenuBarAutoHide()
	// Returns Boolean - Whether menu bar automatically hides itself.

	// win.setMenuBarVisibility(visible) Windows Linux
	// visible Boolean
	// Sets whether the menu bar should be visible. If the menu bar is auto-hide, users can still bring up the menu bar by pressing the single Alt key.

	// win.isMenuBarVisible()
	// Returns Boolean - Whether the menu bar is visible.

	// win.setVisibleOnAllWorkspaces(visible)
	// visible Boolean
	// Sets whether the window should be visible on all workspaces.

	// Note: This API does nothing on Windows.

	// win.isVisibleOnAllWorkspaces()
	// Returns Boolean - Whether the window is visible on all workspaces.

	// Note: This API always returns false on Windows.

	// win.setIgnoreMouseEvents(ignore)
	// ignore Boolean
	// Makes the window ignore all mouse events.

	// All mouse events happened in this window will be passed to the window below this window, but if this window has focus, it will still receive keyboard events.

	// win.setContentProtection(enable) macOS Windows
	// enable Boolean
	// Prevents the window contents from being captured by other apps.

	// On macOS it sets the NSWindow’s sharingType to NSWindowSharingNone. On Windows it calls SetWindowDisplayAffinity with WDA_MONITOR.

	// win.setFocusable(focusable) Windows
	// focusable Boolean
	// Changes whether the window can be focused.

	// win.setParentWindow(parent) Linux macOS
	// parent BrowserWindow
	// Sets parent as current window’s parent window, passing null will turn current window into a top-level window.

	// win.getParentWindow()
	// Returns BrowserWindow - The parent window.

	// win.getChildWindows()
	// Returns BrowserWindow[] - All child windows.

	// win.setAutoHideCursor(autoHide) macOS
	// autoHide Boolean
	// Controls whether to hide cursor when typing.

	// win.setVibrancy(type) macOS
	// type String - Can be appearance-based, light, dark, titlebar, selection, menu, popover, sidebar, medium-light or ultra-dark. See the macOS documentation for more details.
	// Adds a vibrancy effect to the browser window. Passing null or an empty string will remove the vibrancy effect on the window.
}

func NewBrowserWindow(opts *BrowserWindowOptions) *BrowserWindow {
	return &BrowserWindow{
		Object: browserWindow.New(opts),
	}
}

// Static Methods
// The BrowserWindow class has the following static methods:

// BrowserWindow.getAllWindows()
// Returns BrowserWindow[] - An array of all opened browser windows.
func GetAllWindows() []*BrowserWindow {
	ws := browserWindow.Call("getAllWindows")
	ret := []*BrowserWindow{}
	for i := 0; i < ws.Length(); i++ {
		ret = append(ret, &BrowserWindow{
			Object: ws.Index(i),
		})
	}
	return ret
}

// BrowserWindow.getFocusedWindow()
// Returns BrowserWindow - The window that is focused in this application, otherwise returns null.
func GetFocusedWindow() *BrowserWindow {
	fw := browserWindow.Get("getFocusedWindow")
	if fw.String() == "null" {
		return nil
	}
	return &BrowserWindow{
		Object: fw,
	}
}

// BrowserWindow.fromWebContents(webContents)
// webContents WebContents
// Returns BrowserWindow - The window that owns the given webContents.

// BrowserWindow.fromId(id)
// id Integer
// Returns BrowserWindow - The window with the given id.

// BrowserWindow.addDevToolsExtension(path)
// path String
// Adds DevTools extension located at path, and returns extension’s name.

// The extension will be remembered so you only need to call this API once, this API is not for programming use. If you try to add an extension that has already been loaded, this method will not return and instead log a warning to the console.

// The method will also not return if the extension’s manifest is missing or incomplete.

// Note: This API cannot be called before the ready event of the app module is emitted.

// BrowserWindow.removeDevToolsExtension(name)
// name String
// Remove a DevTools extension by name.

// Note: This API cannot be called before the ready event of the app module is emitted.

// BrowserWindow.getDevToolsExtensions()
// Returns Object - The keys are the extension names and each value is an Object containing name and version properties.

// To check if a DevTools extension is installed you can run the following:

// const {BrowserWindow} = require('electron')

// let installed = BrowserWindow.getDevToolsExtensions().hasOwnProperty('devtron')
// console.log(installed)
// Note: This API cannot be called before the ready event of the app module is emitted.
