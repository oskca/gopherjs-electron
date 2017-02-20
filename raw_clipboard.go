package electron

import "github.com/gopherjs/gopherjs/js"

// ClipboardModule version@1.4.15
//
// Perform copy and paste operations on the system clipboard.
type ClipboardModule struct {
	*js.Object
	ReadText func(Type string) (Obj string) `js:"readText"`
	// Writes the text into the clipboard as plain text.
	WriteText func(Text string, Type string) `js:"writeText"`
	ReadHTML  func(Type string) (Obj string) `js:"readHTML"`
	// Writes markup to the clipboard.
	WriteHTML func(Markup string, Type string)     `js:"writeHTML"`
	ReadImage func(Type string) (Obj *NativeImage) `js:"readImage"`
	// Writes image to the clipboard.
	WriteImage func(Image *NativeImage, Type string) `js:"writeImage"`
	ReadRTF    func(Type string) (Obj string)        `js:"readRTF"`
	// Writes the text into the clipboard in RTF.
	WriteRTF func(Text string, Type string) `js:"writeRTF"`
	// Returns an Object containing title and url keys representing the bookmark in the clipboard. The title and url values will be empty strings when the bookmark is unavailable.
	ReadBookmark func() (Obj *ClipboardModuleReadBookmarkObj) `js:"readBookmark"`
	// Writes the title and url into the clipboard as a bookmark. Note: Most apps on Windows don't support pasting bookmarks into them so you can use clipboard.write to write both a bookmark and fallback text to the clipboard.
	WriteBookmark func(Title string, URL string, Type string) `js:"writeBookmark"`
	ReadFindText  func() (Obj string)                         `js:"readFindText"`
	// Writes the text into the find pasteboard as plain text. This method uses synchronous IPC when called from the renderer process.
	WriteFindText func(Text string) `js:"writeFindText"`
	// Clears the clipboard content.
	Clear            func(Type string)                           `js:"clear"`
	AvailableFormats func(Type string) (Obj *js.Object)          `js:"availableFormats"`
	Has              func(Data string, Type string) (Obj bool)   `js:"has"`
	Read             func(Data string, Type string) (Obj string) `js:"read"`
	// Writes data to the clipboard.
	Write func(Data *ClipboardModuleWriteData, Type string) `js:"write"`
}

func GetClipboardModule() *ClipboardModule {
	o := Get("clipboard")
	return &ClipboardModule{
		Object: o,
	}
}

type ClipboardModuleWriteData struct {
	*js.Object
	Text  string       `js:"text"`
	Html  string       `js:"html"`
	Image *NativeImage `js:"image"`
	Rtf   string       `js:"rtf"`
	// The title of the url at .
	Bookmark string `js:"bookmark"`
}

type ClipboardModuleReadBookmarkObj struct {
	*js.Object
	Title string `js:"title"`
	URL   string `js:"url"`
}
