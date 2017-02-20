package electron

import "github.com/gopherjs/gopherjs/js"

// UploadData a Structure
type UploadData struct {
	*js.Object
	// Content being sent.
	Bytes *js.Object `js:"bytes"`
	// Path of file being uploaded.
	File string `js:"file"`
	// UUID of blob data. Use method to retrieve the data.
	BlobUUID string `js:"blobUUID"`
}
