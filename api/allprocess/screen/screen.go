// Package screen retrieves information about screen size, displays, cursor position, etc.
// You cannot require or use this module until the ready event of the app module is emitted.
// screen is an EventEmitter.
// Note: In the renderer / DevTools, window.screen is a reserved DOM property,
// so writing let {screen} = require('electron') will not work.
package screen

import electron "github.com/oskca/gopherjs-electron"

var screen = electron.Get("screen")

// An example of creating a window that fills the whole screen:

// const electron = require('electron')
// const {app, BrowserWindow} = electron

// let win

// app.on('ready', () => {
//   const {width, height} = electron.screen.getPrimaryDisplay().workAreaSize
//   win = new BrowserWindow({width, height})
//   win.loadURL('https://github.com')
// })
// Another example of creating a window in the external display:

// const electron = require('electron')
// const {app, BrowserWindow} = require('electron')

// let win

// app.on('ready', () => {
//   let displays = electron.screen.getAllDisplays()
//   let externalDisplay = displays.find((display) => {
//     return display.bounds.x !== 0 || display.bounds.y !== 0
//   })

//   if (externalDisplay) {
//     win = new BrowserWindow({
//       x: externalDisplay.bounds.x + 50,
//       y: externalDisplay.bounds.y + 50
//     })
//     win.loadURL('https://github.com')
//   }
// })

// Events
const (
	// The screen module emits the following events:

	// Event: ‘display-added’
	// Returns:
	// event Event
	// newDisplay Display
	// Emitted when newDisplay has been added.
	EvtDisplayAdded = "display-added"

	// Event: ‘display-removed’
	// Returns:
	// event Event
	// oldDisplay Display
	// Emitted when oldDisplay has been removed.
	EvtDisplayRemoved = "display-removed"

	// Event: ‘display-metrics-changed’
	// Returns:
	// event Event
	// display Display
	// changedMetrics String[]
	// Emitted when one or more metrics change in a display.
	// The changedMetrics is an array of strings that describe the changes.
	// Possible changes are bounds, workArea, scaleFactor and rotation.
	EvtDisplayMetricsChanged = "display-remmetrics-changedoved"
)

// Methods

// The screen module has the following methods:

// screen.getCursorScreenPoint()
// Returns Object:
// x Integer
// y Integer
// The current absolute position of the mouse pointer.
func GetCursorScreenPoint() *electron.Point {
	return &electron.Point{
		Object: screen.Call("getCursorScreenPoint"),
	}
}

// screen.getPrimaryDisplay()
// Returns Display - The primary display.
func GetPrimaryDisplay() *electron.Display {
	return &electron.Display{
		Object: screen.Call("getPrimaryDisplay"),
	}
}

// screen.getAllDisplays()
// Returns Display[] - An array of displays that are currently available.
func GetAllDisplays() []*electron.Display {
	ret := []*electron.Display{}
	s := screen.Call("getAllDisplays")
	for index := 0; index < s.Length(); index++ {
		o := s.Index(index)
		ret = append(ret, &electron.Display{
			Object: o,
		})
	}
	return ret
}

// screen.getDisplayNearestPoint(point)
// point Object
// x Integer
// y Integer
// Returns Display - The display nearest the specified point.
func GetDisplayNearestPoint() *electron.Point {
	return &electron.Point{
		Object: screen.Call("getDisplayNearestPoint"),
	}
}

// screen.getDisplayMatching(rect)
// rect Rectangle
// Returns Display - The display that most closely intersects the provided bounds.
func GetDisplayMatching(rect *electron.Rect) *electron.Display {
	return &electron.Display{
		Object: screen.Call("getDisplayMatching", rect),
	}
}
