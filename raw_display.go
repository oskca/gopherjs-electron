package electron

import "github.com/gopherjs/gopherjs/js"

// Display a Structure
type Display struct {
	*js.Object
	// Unique identifier associated with the display.
	Id float64 `js:"id"`
	// Can be 0, 90, 180, 270, represents screen rotation in clock-wise degrees.
	Rotation float64 `js:"rotation"`
	// Output device's pixel scale factor.
	ScaleFactor float64 `js:"scaleFactor"`
	// Can be , , .
	TouchSupport DisplayDisplayTouchSupport  `js:"touchSupport"`
	Bounds       *js.Object                  `js:"bounds"`
	Size         *DisplayDisplaySize         `js:"size"`
	WorkArea     *js.Object                  `js:"workArea"`
	WorkAreaSize *DisplayDisplayWorkAreaSize `js:"workAreaSize"`
}

type DisplayDisplaySize struct {
	*js.Object
	Height float64 `js:"height"`
	Width  float64 `js:"width"`
}

type DisplayDisplayWorkAreaSize struct {
	*js.Object
	Height float64 `js:"height"`
	Width  float64 `js:"width"`
}

type DisplayDisplayTouchSupport string

// consts
const (
	DisplayDisplayTouchSupportAvailable   DisplayDisplayTouchSupport = "available"
	DisplayDisplayTouchSupportUnavailable DisplayDisplayTouchSupport = "unavailable"
	DisplayDisplayTouchSupportUnknown     DisplayDisplayTouchSupport = "unknown"
)
