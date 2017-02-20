package electron

import "github.com/gopherjs/gopherjs/js"

// UploadBlob a Structure
type UploadBlob struct {
	*js.Object
	// .
	Type string `js:"type"`
	// UUID of blob data to upload.
	BlobUUID string `js:"blobUUID"`
}
