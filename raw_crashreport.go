package electron

import "github.com/gopherjs/gopherjs/js"

// CrashReport a Structure
type CrashReport struct {
	*js.Object
	Date string `js:"date"`
	ID   int64  `js:"ID"`
}
