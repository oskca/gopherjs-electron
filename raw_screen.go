package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

const (
	// Emitted when newDisplay has been added.
	EvtScreenDisplayAdded = "display-added"
	// Emitted when oldDisplay has been removed.
	EvtScreenDisplayRemoved = "display-removed"
	// Emitted when one or more metrics change in a display. The changedMetrics is an array of strings that describe the changes. Possible changes are bounds, workArea, scaleFactor and rotation.
	EvtScreenDisplayMetricsChanged = "display-metrics-changed"
)

// ScreenModule version@1.4.15
//
// Retrieve information about screen size, displays, cursor position, etc.
type ScreenModule struct {
	*events.Emitter
	// The current absolute position of the mouse pointer.
	GetCursorScreenPoint   func() (Obj *ScreenModuleGetCursorScreenPointObj)                     `js:"getCursorScreenPoint"`
	GetPrimaryDisplay      func() (Obj *js.Object)                                               `js:"getPrimaryDisplay"`
	GetAllDisplays         func() (Obj *js.Object)                                               `js:"getAllDisplays"`
	GetDisplayNearestPoint func(Point *ScreenModuleGetDisplayNearestPointPoint) (Obj *js.Object) `js:"getDisplayNearestPoint"`
	GetDisplayMatching     func(Rect *js.Object) (Obj *js.Object)                                `js:"getDisplayMatching"`
}

func GetScreenModule() *ScreenModule {
	o := Get("screen")
	return &ScreenModule{
		Emitter: events.New(o),
	}
}

type ScreenModuleGetCursorScreenPointObj struct {
	*js.Object
	X int64 `js:"x"`
	Y int64 `js:"y"`
}

type ScreenModuleGetDisplayNearestPointPoint struct {
	*js.Object
	X int64 `js:"x"`
	Y int64 `js:"y"`
}
