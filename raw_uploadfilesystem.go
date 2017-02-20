package electron

import "github.com/gopherjs/gopherjs/js"

// UploadFileSystem a Structure
type UploadFileSystem struct {
	*js.Object
	// .
	Type string `js:"type"`
	// FileSystem url to read data for upload.
	FilsSystemURL string `js:"filsSystemURL"`
	// Defaults to .
	Offset int64 `js:"offset"`
	// Number of bytes to read from . Defaults to .
	Length int64 `js:"length"`
	// Last Modification time in number of seconds sine the UNIX epoch.
	ModificationTime float64 `js:"modificationTime"`
}
