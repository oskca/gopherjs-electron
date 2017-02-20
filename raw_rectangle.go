package electron

import "github.com/gopherjs/gopherjs/js"

// Rectangle a Structure
type Rectangle struct {
	*js.Object
	// The x coordinate of the origin of the rectangle
	X float64 `js:"x"`
	// The y coordinate of the origin of the rectangle
	Y      float64 `js:"y"`
	Width  float64 `js:"width"`
	Height float64 `js:"height"`
}
