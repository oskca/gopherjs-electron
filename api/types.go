package api

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-electron/api/main/webcontents"
)

// Point defines point for electron
type Point struct {
	*js.Object
	X int `js:"x"`
	Y int `js:"y"`
}

// Size defines size for electron
type Size struct {
	*js.Object
	Width  int `js:"width"`
	Height int `js:"height"`
}

// Rect defines rectangle for electron
type Rect struct {
	*js.Object
	X      int `js:"x"`
	Y      int `js:"y"`
	Width  int `js:"width"`
	Height int `js:"height"`
}

// Display defines screen.Display for electron
type Display struct {
	*js.Object
	ID           int    `js:"id "`           // id Number - Unique identifier associated with the display.
	Rotation     int    `js:"rotation "`     // rotation Number - Can be 0, 90, 180, 270, represents screen rotation in clock-wise degrees.
	ScaleFactor  int    `js:"scaleFactor "`  // scaleFactor Number - Output device’s pixel scale factor.
	TouchSupport string `js:"touchSupport "` // touchSupport String - Can be available, unavailable, unknown.
	Bounds       *Rect  `js:"bounds "`       // bounds Rectangle
	Size         *Size  `js:"size "`         // size Object
	WorkArea     *Rect  `js:"workArea "`     // workArea Rectangle
	WorkAreaSize *Size  `js:"workAreaSize "` // workAreaSize Object
}

// IpcEvent is usd in IPC
type IpcEvent struct {
	*js.Object
	// Event object
	// The event object passed to the callback has the following methods:
	// event.returnValue
	// Set this to the value to be returned in a synchronous message.
	ReturnValue interface{} `js:"returnValue"`

	// event.sender
	// Returns the webContents that sent the message, you can call event.sender.send to reply to the asynchronous message, see webContents.send for more information.
	Sender *webcontents.WebContents `js:"sender"`
}

// IpcListener is used in ipc
type IpcListener func(evt *IpcEvent, args ...*js.Object)
