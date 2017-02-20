package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when the tray icon is clicked.
	EvtTrayClick = "click"
	// Emitted when the tray icon is right clicked.
	EvtTrayRightClick = "right-click"
	// Emitted when the tray icon is double clicked.
	EvtTrayDoubleClick = "double-click"
	// Emitted when the tray balloon shows.
	EvtTrayBalloonShow = "balloon-show"
	// Emitted when the tray balloon is clicked.
	EvtTrayBalloonClick = "balloon-click"
	// Emitted when the tray balloon is closed because of timeout or user manually closes it.
	EvtTrayBalloonClosed = "balloon-closed"
	// Emitted when any dragged items are dropped on the tray icon.
	EvtTrayDrop = "drop"
	// Emitted when dragged files are dropped in the tray icon.
	EvtTrayDropFiles = "drop-files"
	// Emitted when dragged text is dropped in the tray icon.
	EvtTrayDropText = "drop-text"
	// Emitted when a drag operation enters the tray icon.
	EvtTrayDragEnter = "drag-enter"
	// Emitted when a drag operation exits the tray icon.
	EvtTrayDragLeave = "drag-leave"
	// Emitted when a drag operation ends on the tray or ends at another location.
	EvtTrayDragEnd = "drag-end"
)

// Tray version@1.4.15
//
// Add icons and context menus to the system's notification area.
type Tray struct {
	*events.Emitter
	// Destroys the tray icon immediately.
	Destroy func() `js:"destroy"`
	// Sets the image associated with this tray icon.
	SetImage func(Image *NativeImage) `js:"setImage"`
	// Sets the image associated with this tray icon when pressed on macOS.
	SetPressedImage func(Image *NativeImage) `js:"setPressedImage"`
	// Sets the hover text for this tray icon.
	SetToolTip func(ToolTip string) `js:"setToolTip"`
	// Sets the title displayed aside of the tray icon in the status bar.
	SetTitle func(Title string) `js:"setTitle"`
	// Sets when the tray's icon background becomes highlighted (in blue). Note: You can use highlightMode with a BrowserWindow by toggling between 'never' and 'always' modes when the window visibility changes.
	SetHighlightMode func(Mode TraySetHighlightModeMode) `js:"setHighlightMode"`
	// Displays a tray balloon.
	DisplayBalloon func(Options *TrayDisplayBalloonOptions) `js:"displayBalloon"`
	// Pops up the context menu of the tray icon. When menu is passed, the menu will be shown instead of the tray icon's context menu. The position is only available on Windows, and it is (0, 0) by default.
	PopUpContextMenu func(Menu *Menu, Position *TrayPopUpContextMenuPosition) `js:"popUpContextMenu"`
	// Sets the context menu for this icon.
	SetContextMenu func(Menu *Menu) `js:"setContextMenu"`
	// The bounds of this tray icon as Object.
	GetBounds   func() (Obj *js.Object) `js:"getBounds"`
	IsDestroyed func() (Obj bool)       `js:"isDestroyed"`
}

func WrapTray(o *js.Object) *Tray {
	return &Tray{
		Emitter: events.New(o),
	}
}

func NewTray(Image *NativeImage) *Tray {
	o := electron.Get("Tray")
	ret := o.New(Image)
	return WrapTray(ret)
}

type TrayDisplayBalloonOptions struct {
	*js.Object
	// (optional)
	Icon *NativeImage `js:"icon"`
	// (optional)
	Title string `js:"title"`
	// (optional)
	Content string `js:"content"`
}

type TrayPopUpContextMenuPosition struct {
	*js.Object
	X int64 `js:"x"`
	Y int64 `js:"y"`
}

type TraySetHighlightModeMode string

// consts
const (
	TraySetHighlightModeModeSelection TraySetHighlightModeMode = "selection"
	TraySetHighlightModeModeAlways    TraySetHighlightModeMode = "always"
	TraySetHighlightModeModeNever     TraySetHighlightModeMode = "never"
)
