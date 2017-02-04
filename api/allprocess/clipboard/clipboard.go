// Package clipboard performs copy and paste operations on the system clipboard.
// Process: Main, Renderer
package clipboard

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/allprocess/nativeimage"
)

// const types
const (
	TypeSelection = "selection"
)

// The following example shows how to write a string to the clipboard:

// const {clipboard} = require('electron')
// clipboard.writeText('Example String')
// On X Window systems, there is also a selection clipboard. To manipulate it you need to pass selection to each method:

// const {clipboard} = require('electron')
// clipboard.writeText('Example String', 'selection')
// console.log(clipboard.readText('selection'))
// Methods

// Get returns electron module clipboard
func Get() *Clipboard {
	return &Clipboard{
		Object: electron.Get("clipboard"),
	}
}

// Clipboard wrapper
type Clipboard struct {
	*js.Object

	// The clipboard module has the following methods:

	// Note: Experimental APIs are marked as such and could be removed in future.

	// clipboard.readText([type])
	// type String (optional)
	// Returns String - The content in the clipboard as plain text.
	ReadText func(typ ...string) string `js:"readText"`

	// clipboard.writeText(text[, type])
	// text String
	// type String (optional)
	// Writes the text into the clipboard as plain text.
	WriteText func(text string, typ ...string) `js:"writeText"`

	// clipboard.readHTML([type])
	// type String (optional)
	// Returns String - The content in the clipboard as markup.
	ReadHTML func(typ ...string) string `js:"readHTML"`

	// clipboard.writeHTML(markup[, type])
	// markup String
	// type String (optional)
	// Writes markup to the clipboard.
	WriteHTML func(markup string, typ ...string) `js:"writeHTML"`

	// clipboard.readImage([type])
	// type String (optional)
	// Returns NativeImage - The image content in the clipboard.
	ReadImage func(typ ...string) (image *nativeimage.NativeImage) `js:"readImage"`

	// clipboard.writeImage(image[, type])
	// image NativeImage
	// type String (optional)
	// Writes image to the clipboard.
	WriteImage func(image *nativeimage.NativeImage, typ ...string) `js:"writeImage"`

	// clipboard.readRTF([type])
	// type String (optional)
	// Returns String - The content in the clipboard as RTF.
	readRTF func(typ ...string) string `js:"readRTF"`

	// clipboard.writeRTF(text[, type])
	// text String
	// type String (optional)
	// Writes the text into the clipboard in RTF.
	WriteRTF func(text string, typ ...string) `js:"writeRTF"`

	// clipboard.readBookmark() macOS Windows
	// Returns Object:
	//
	// title String
	// url String
	// Returns an Object containing title and url keys representing the bookmark in the clipboard. The title and url values will be empty strings when the bookmark is unavailable.
	ReadBookmark func() (obj *js.Object) `js:"readBookmark"`

	// clipboard.writeBookmark(title, url[, type]) macOS Windows
	// title String
	// url String
	// type String (optional)
	// Writes the title and url into the clipboard as a bookmark.
	WriteBookmark func(title, url string, typ ...string) `js:"writeBookmark"`

	// Note: Most apps on Windows don’t support pasting bookmarks into them so you can use clipboard.write to write both a bookmark and fallback text to the clipboard.
	//
	// clipboard.write({
	//   text: 'http://electron.atom.io',
	//   bookmark: 'Electron Homepage'
	// })

	// clipboard.readFindText() macOS
	// Returns String - The text on the find pasteboard. This method uses synchronous IPC when called from the renderer process. The cached value is reread from the find pasteboard whenever the application is activated.
	ReadFindText func() string `js:"readFindText"`

	// clipboard.writeFindText(text) macOS
	// text String
	// Writes the text into the find pasteboard as plain text.
	// This method uses synchronous IPC when called from the renderer process.
	WriteFindText func(text string) `js:"writeFindText"`

	// clipboard.clear([type])
	// type String (optional)
	// Clears the clipboard content.
	Clear func(typ ...string) string `js:"clear"`

	// clipboard.availableFormats([type])
	// type String (optional)
	// Returns String[] - An array of supported formats for the clipboard type.
	AvailableFormats func(typ ...string) []string `js:"availableFormats"`

	// clipboard.has(data[, type]) Experimental
	// data String
	// type String (optional)
	// Returns Boolean - Whether the clipboard supports the format of specified data.
	Has func(data string, typ ...string) string `js:"has"`

	// const {clipboard} = require('electron')
	// console.log(clipboard.has('<p>selection</p>'))

	// clipboard.read(data[, type]) Experimental
	// data String
	// type String (optional)
	// Returns String - Reads data from the clipboard.
	Read func(data string, typ ...string) string `js:"read"`

	// clipboard.write(data[, type])
	// data Object
	// text String (optional)
	// html String (optional)
	// image NativeImage (optional)
	// rtf String (optional)
	// bookmark String (optional) - The title of the url at text.
	// type String (optional)
	// const {clipboard} = require('electron')
	// clipboard.write({text: 'test', html: '<b>test</b>'})
	// Writes data to the clipboard.
	Write func(data *js.Object, typ ...string) `js:"write"`
}
