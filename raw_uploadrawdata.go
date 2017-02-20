package electron

import "github.com/gopherjs/gopherjs/js"

// UploadRawData a Structure
type UploadRawData struct {
	*js.Object
	// .
	Type string `js:"type"`
	// Data to be uploaded.
	Bytes *js.Object `js:"bytes"`
}
