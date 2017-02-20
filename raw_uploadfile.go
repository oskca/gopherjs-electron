package electron

import "github.com/gopherjs/gopherjs/js"

// UploadFile a Structure
type UploadFile struct {
	*js.Object
	// .
	Type string `js:"type"`
	// Path of file to be uploaded.
	FilePath string `js:"filePath"`
	// Defaults to .
	Offset int64 `js:"offset"`
	// Number of bytes to read from . Defaults to .
	Length int64 `js:"length"`
	// Last Modification time in number of seconds sine the UNIX epoch.
	ModificationTime float64 `js:"modificationTime"`
}
