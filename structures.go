package electron

import (
	"github.com/gopherjs/gopherjs/js"
)

// Rect defines rectangle for electron
type Rect struct {
	*js.Object
	X      int `js:"x"`
	Y      int `js:"y"`
	Width  int `js:"width"`
	Height int `js:"height"`
}
