package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when the download has been updated and is not done. The state can be one of following:
	EvtDownloadItemUpdated = "updated"
	// Emitted when the download is in a terminal state. This includes a completed download, a cancelled download (via downloadItem.cancel()), and interrupted download that can't be resumed. The state can be one of following:
	EvtDownloadItemDone = "done"
)

// DownloadItem version@1.4.15
//
// Control file downloads from remote sources.
type DownloadItem struct {
	*events.Emitter
	// The API is only available in session's will-download callback function. If user doesn't set the save path via the API, Electron will use the original routine to determine the save path(Usually prompts a save dialog).
	SetSavePath func(Path string)   `js:"setSavePath"`
	GetSavePath func() (Obj string) `js:"getSavePath"`
	// Pauses the download.
	Pause    func()            `js:"pause"`
	IsPaused func() (Obj bool) `js:"isPaused"`
	// Resumes the download that has been paused.
	Resume func() `js:"resume"`
	// Resumes Boolean - Whether the download can resume.
	CanResume func() `js:"canResume"`
	// Cancels the download operation.
	Cancel         func()              `js:"cancel"`
	GetURL         func() (Obj string) `js:"getURL"`
	GetMimeType    func() (Obj string) `js:"getMimeType"`
	HasUserGesture func() (Obj bool)   `js:"hasUserGesture"`
	// Note: The file name is not always the same as the actual one saved in local disk. If user changes the file name in a prompted download saving dialog, the actual name of saved file will be different.
	GetFilename func() (Obj string) `js:"getFilename"`
	// If the size is unknown, it returns 0.
	GetTotalBytes         func() (Obj int64)  `js:"getTotalBytes"`
	GetReceivedBytes      func() (Obj int64)  `js:"getReceivedBytes"`
	GetContentDisposition func() (Obj string) `js:"getContentDisposition"`
	// Note: The following methods are useful specifically to resume a cancelled item when session is restarted.
	GetState            func() (Obj string)     `js:"getState"`
	GetURLChain         func() (Obj *js.Object) `js:"getURLChain"`
	GetLastModifiedTime func() (Obj string)     `js:"getLastModifiedTime"`
	GetETag             func() (Obj string)     `js:"getETag"`
	GetStartTime        func() (Obj float64)    `js:"getStartTime"`
}

func WrapDownloadItem(o *js.Object) *DownloadItem {
	return &DownloadItem{
		Emitter: events.New(o),
	}
}
