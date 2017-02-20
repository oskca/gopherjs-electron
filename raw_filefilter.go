package electron

import "github.com/gopherjs/gopherjs/js"

// FileFilter a Structure
type FileFilter struct {
	*js.Object
	Name       string     `js:"name"`
	Extensions *js.Object `js:"extensions"`
}
