package electron

import "github.com/gopherjs/gopherjs/js"

// DialogModule version@1.4.15
//
// Display native system dialogs for opening and saving files, alerting, etc.
type DialogModule struct {
	*js.Object
	// The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed or selected when you want to limit the user to a specific type. For example: The extensions array should contain extensions without wildcards or dots (e.g. 'png' is good but '.png' and '*.png' are bad). To show all files, use the '*' wildcard (no other wildcard is supported). If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filenames) Note: On Windows and Linux an open dialog can not be both a file selector and a directory selector, so if you set properties to ['openFile', 'openDirectory'] on these platforms, a directory selector will be shown.
	ShowOpenDialog func(BrowserWindow *BrowserWindow, Options *DialogModuleShowOpenDialogOptions, Callback DialogModuleShowOpenDialogCallback) (Obj *js.Object) `js:"showOpenDialog"`
	// The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. The filters specifies an array of file types that can be displayed, see dialog.showOpenDialog for an example. If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filename)
	ShowSaveDialog func(BrowserWindow *BrowserWindow, Options *DialogModuleShowSaveDialogOptions, Callback DialogModuleShowSaveDialogCallback) (Obj string) `js:"showSaveDialog"`
	// Shows a message box, it will block the process until the message box is closed. It returns the index of the clicked button. The browserWindow argument allows the dialog to attach itself to a parent window, making it modal. If a callback is passed, the API call will be asynchronous and the result will be passed via callback(response).
	ShowMessageBox func(BrowserWindow *BrowserWindow, Options *DialogModuleShowMessageBoxOptions, Callback DialogModuleShowMessageBoxCallback) (Obj int64) `js:"showMessageBox"`
	// Displays a modal dialog that shows an error message. This API can be called safely before the ready event the app module emits, it is usually used to report errors in early stage of startup.  If called before the app readyevent on Linux, the message will be emitted to stderr, and no GUI dialog will appear.
	ShowErrorBox func(Title string, Content string) `js:"showErrorBox"`
}

func GetDialogModule() *DialogModule {
	o := Get("dialog")
	return &DialogModule{
		Object: o,
	}
}

type DialogModuleShowOpenDialogOptions struct {
	*js.Object
	Title       string `js:"title"`
	DefaultPath string `js:"defaultPath"`
	// Custom label for the confirmation button, when left empty the default label will be used.
	ButtonLabel string     `js:"buttonLabel"`
	Filters     *js.Object `js:"filters"`
	// Contains which features the dialog should use, can contain , , , and .
	Properties *js.Object `js:"properties"`
	// Normalize the keyboard access keys across platforms. Default is . Enabling this assumes is used in the button labels for the placement of the keyboard shortcut access key and labels will be converted so they work correctly on each platform, characters are removed on macOS, converted to on Linux, and left untouched on Windows. For example, a button label of will be converted to on Linux and on macOS and can be selected via on Windows and Linux.
	NormalizeAccessKeys bool `js:"normalizeAccessKeys"`
}

type DialogModuleShowOpenDialogCallback func( // An array of file paths chosen by the user
	FilePaths *js.Object)
type DialogModuleShowSaveDialogOptions struct {
	*js.Object
	Title       string `js:"title"`
	DefaultPath string `js:"defaultPath"`
	// Custom label for the confirmation button, when left empty the default label will be used.
	ButtonLabel string     `js:"buttonLabel"`
	Filters     *js.Object `js:"filters"`
}

type DialogModuleShowSaveDialogCallback func(Filename string)
type DialogModuleShowMessageBoxOptions struct {
	*js.Object
	// Can be , , , or . On Windows, "question" displays the same icon as "info", unless you set an icon using the "icon" option.
	Type string `js:"type"`
	// Array of texts for buttons. On Windows, an empty array will result in one button labeled "OK".
	Buttons *js.Object `js:"buttons"`
	// Index of the button in the buttons array which will be selected by default when the message box opens.
	DefaultId int64 `js:"defaultId"`
	// Title of the message box, some platforms will not show it.
	Title string `js:"title"`
	// Content of the message box.
	Message string `js:"message"`
	// Extra information of the message.
	Detail string       `js:"detail"`
	Icon   *NativeImage `js:"icon"`
	// The value will be returned when user cancels the dialog instead of clicking the buttons of the dialog. By default it is the index of the buttons that have "cancel" or "no" as label, or 0 if there is no such buttons. On macOS and Windows the index of the "Cancel" button will always be used as even if it is specified.
	CancelId int64 `js:"cancelId"`
	// On Windows Electron will try to figure out which one of the are common buttons (like "Cancel" or "Yes"), and show the others as command links in the dialog. This can make the dialog appear in the style of modern Windows apps. If you don't like this behavior, you can set to .
	NoLink bool `js:"noLink"`
}

type DialogModuleShowMessageBoxCallback func( // The index of the button that was clicked
	Response float64)
