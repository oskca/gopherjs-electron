package electron

import "github.com/gopherjs/gopherjs/js"

// MimeTypedBuffer a Structure
type MimeTypedBuffer struct {
	*js.Object
	// The mimeType of the Buffer that you are sending
	MimeType string `js:"mimeType"`
	// The actual Buffer content
	Buffer *js.Object `js:"buffer"`
}
