package electron

import "github.com/gopherjs/gopherjs/js"

// FileFilterEx wraps
type FileFilterEx struct {
	Name       string   // name String
	Extensions []string // extensions String[]
}

func (f FileFilterEx) toJs() js.M {
	m := make(js.M)
	if f.Name != "" {
		m["name"] = f.Name
	}
	if f.Extensions != nil {
		es := make(js.S, 0)
		for j := 0; j < len(f.Extensions); j++ {
			es = append(es, f.Extensions[j])
		}
		m["extensions"] = es
	}
	return m
}

// DialogOptionOpen wraps the dialog creating options
type DialogOptionOpen struct {
	Title       string         // title String (optional)
	DefaultPath string         // defaultPath String (optional)
	ButtonLabel string         // buttonLabel String (optional) - Custom label for the confirmation button, when left empty the default label will be used.
	Filters     []FileFilterEx // filters FileFilterEx[] (optional)
	Properties  []string       // properties String[] (optional) - Contains which features the dialog should use, can contain openFile, openDirectory, multiSelections, createDirectory and showHiddenFiles.
	// normalizeAccessKeys Boolean (optional) - Normalize the keyboard access keys across platforms.
	// Default is false. Enabling this assumes & is used in the button labels for the placement of
	// the keyboard shortcut access key and labels will be converted so they work correctly on
	// each platform, & characters are removed on macOS, converted to _ on Linux, and
	// left untouched on Windows. For example, a button label of Vie&w will be converted to
	// Vie_w on Linux and View on macOS and can be selected via Alt-W on Windows and Linux.
	NormalizeAccessKeys bool
}

func (o DialogOptionOpen) toJs() js.M {
	ret := make(js.M)
	if o.Title != "" {
		ret["title"] = o.Title
	}
	if o.DefaultPath != "" {
		ret["defaultPath"] = o.DefaultPath
	}
	if o.ButtonLabel != "" {
		ret["buttonLabel"] = o.ButtonLabel
	}
	if o.Filters != nil {
		fs := make(js.S, 0)
		for i := 0; i < len(o.Filters); i++ {
			fs = append(fs, o.Filters[i].toJs())
		}
		ret["filters"] = fs
	}
	if o.Properties != nil {
		p := make(js.S, 0)
		for i := 0; i < len(o.Properties); i++ {
			p = append(p, o.Properties[i])
		}
		ret["properties"] = p
	}
	return ret
}

// OpenDialog properties
const (
	DialogPropOpenFile        = "openFile"
	DialogPropOpenDirectory   = "openDirectory"
	DialogPropMultiSelections = "multiSelections"
	DialogPropCreateDirectory = "createDirectory"
	DialogPropShowHiddenFiles = "showHiddenFiles"
)

// ShowOpenDialogEx ([browserWindow, ]options[, callback])
// browserWindow BrowserWindow (optional)
// options Object
// 	title String (optional)
// 	defaultPath String (optional)
// 	buttonLabel String (optional) - Custom label for the confirmation button, when left empty the default label will be used.
// 	filters FileFilterEx[] (optional)
// 	properties String[] (optional) - Contains which features the dialog should use, can contain openFile, openDirectory, multiSelections, createDirectory and showHiddenFiles.
// 	normalizeAccessKeys Boolean (optional) - Normalize the keyboard access keys across platforms. Default is false. Enabling this assumes & is used in the button labels for the placement of the keyboard shortcut access key and labels will be converted so they work correctly on each platform, & characters are removed on macOS, converted to _ on Linux, and left untouched on Windows. For example, a button label of Vie&w will be converted to Vie_w on Linux and View on macOS and can be selected via Alt-W on Windows and Linux.
// callback Function (optional)
// 	filePaths String[] - An array of file paths chosen by the user
// Returns String[], an array of file paths chosen by the user, if the callback is provided it returns undefined.
//
// The browserWindow argument allows the dialog to attach itself to a parent window, making it modal.
//
// The filters specifies an array of file types that can be displayed or selected when you
// want to limit the user to a specific type. For example:
//
// {
//   filters: [
//     {name: 'Images', extensions: ['jpg', 'png', 'gif']},
//     {name: 'Movies', extensions: ['mkv', 'avi', 'mp4']},
//     {name: 'Custom File Type', extensions: ['as']},
//     {name: 'All Files', extensions: ['*']}
//   ]
// }
// The extensions array should contain extensions without wildcards or dots
// (e.g. 'png' is good but '.png' and '*.png' are bad).
// To show all files, use the '*' wildcard (no other wildcard is supported).
//
// If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filenames)
//
// Note: On Windows and Linux an open dialog can not be both a file selector and
// a directory selector, so if you set properties to ['openFile', 'openDirectory'] on
// these platforms, a directory selector will be shown.
func (d *DialogModule) ShowOpenDialogEx(opt DialogOptionOpen, bw ...*BrowserWindow) (filePaths []string) {
	var out *js.Object
	if len(bw) > 0 {
		out = d.Call("showOpenDialog", bw[0], opt.toJs())
	} else {
		out = d.Call("showOpenDialog", opt.toJs())
	}
	// check null value
	if out.String() == "undefined" {
		return
	}
	filePaths = []string{}
	for i := 0; i < out.Length(); i++ {
		filePaths = append(filePaths, out.Index(i).String())
	}
	return
	// return out.Interface().([]string)
	// js.Global.Get("Uint8Array").New(buffer).Interface().([]byte)
}

// DialogOptionSave wraps for ShowSaveDialog
type DialogOptionSave struct {
	Title       string         // title String (optional)
	DefaultPath string         // defaultPath String (optional)
	ButtonLabel string         // buttonLabel String (optional) - Custom label for the confirmation button, when left empty the default label will be used.
	Filters     []FileFilterEx // filters FileFilterEx[] (optional)
}

func (o DialogOptionSave) toJs() js.M {
	ret := make(js.M)
	if o.Title != "" {
		ret["title"] = o.Title
	}
	if o.DefaultPath != "" {
		ret["defaultPath"] = o.DefaultPath
	}
	if o.ButtonLabel != "" {
		ret["buttonLabel"] = o.ButtonLabel
	}
	if o.Filters != nil {
		fs := make(js.S, 0)
		for i := 0; i < len(o.Filters); i++ {
			fs = append(fs, o.Filters[i].toJs())
		}
		ret["filters"] = fs
	}
	return ret
}

// ShowSaveDialogEx dialog.showSaveDialog([browserWindow, ]options[, callback])
// browserWindow BrowserWindow (optional)
// options Object
// 	title String (optional)
// 	defaultPath String (optional)
// 	buttonLabel String (optional) - Custom label for the confirmation button, when left empty the default label will be used.
// 	filters FileFilterEx[] (optional)
// callback Function (optional)
// 	filename String
// Returns String, the path of the file chosen by the user, if a callback is provided it returns undefined.
//
// The browserWindow argument allows the dialog to attach itself to a parent window, making it modal.
//
// The filters specifies an array of file types that can be displayed, see dialog.showOpenDialog for an example.
//
// If a callback is passed, the API call will be asynchronous and the result will be passed via callback(filename)
func (d *DialogModule) ShowSaveDialogEx(opt DialogOptionSave, bw ...*BrowserWindow) (filepath string) {
	var out *js.Object
	if len(bw) > 0 {
		out = d.Call("showSaveDialog", bw[0], opt.toJs())
	} else {
		out = d.Call("showSaveDialog", opt.toJs())
	}
	// check null value
	return out.String()
}

// const for Message Type
const (
	DialogTypeNone     = "none"
	DialogTypeInfo     = "info"
	DialogTypeError    = "error"
	DialogTypeQuestion = "question"
	DialogTypeWarning  = "warning"
)

// DialogOptionMessage options
type DialogOptionMessage struct {
	Type      string       // type String (optional) - Can be "none", "info", "error", "question" or "warning". On Windows, “question” displays the same icon as “info”, unless you set an icon using the “icon” option.
	Buttons   []string     // buttons String[] (optional) - Array of texts for buttons. On Windows, an empty array will result in one button labeled “OK”.
	DefaultID int          // defaultId Integer (optional) - Index of the button in the buttons array which will be selected by default when the message box opens.
	Title     string       // title String (optional) - Title of the message box, some platforms will not show it.
	Message   string       // message String - Content of the message box.
	Detail    string       // detail String (optional) - Extra information of the message.
	Icon      *NativeImage // icon NativeImage (optional)
	CancelID  int          // cancelId Integer (optional) - The value will be returned when user cancels the dialog instead of clicking the buttons of the dialog. By default it is the index of the buttons that have “cancel” or “no” as label, or 0 if there is no such buttons. On macOS and Windows the index of the “Cancel” button will always be used as cancelId even if it is specified.
	NoLink    bool         // noLink Boolean (optional) - On Windows Electron will try to figure out which one of the buttons are common buttons (like “Cancel” or “Yes”), and show the others as command links in the dialog. This can make the dialog appear in the style of modern Windows apps. If you don’t like this behavior, you can set noLink to true.
}

func (o DialogOptionMessage) toJs() js.M {
	m := make(js.M)
	if o.Type != "" {
		m["type"] = o.Type
	}
	if o.Buttons != nil {
		s := make(js.S, 0)
		for i := 0; i < len(o.Buttons); i++ {
			s = append(s, o.Buttons[i])
		}
		m["buttons"] = s
	}
	if o.DefaultID != 0 {
		m["defaultID"] = o.DefaultID
	}
	if o.Title != "" {
		m["title"] = o.Title
	}
	if o.Message != "" {
		m["message"] = o.Message
	}
	if o.Detail != "" {
		m["detail"] = o.Detail
	}
	if o.Icon != nil {
		m["icon"] = o.Icon
	}
	if o.NoLink != false {
		m["noLink"] = o.NoLink
	}
	return m
}

// ShowMessageBoxEx dialog.showMessageBox([browserWindow, ]options[, callback])
// browserWindow BrowserWindow (optional)
// options Object
// 	type String (optional) - Can be "none", "info", "error", "question" or "warning". On Windows, “question” displays the same icon as “info”, unless you set an icon using the “icon” option.
// 	buttons String[] (optional) - Array of texts for buttons. On Windows, an empty array will result in one button labeled “OK”.
// 	defaultId Integer (optional) - Index of the button in the buttons array which will be selected by default when the message box opens.
// 	title String (optional) - Title of the message box, some platforms will not show it.
// 	message String - Content of the message box.
// 	detail String (optional) - Extra information of the message.
// 	icon NativeImage (optional)
// 	cancelId Integer (optional) - The value will be returned when user cancels the dialog instead of clicking the buttons of the dialog. By default it is the index of the buttons that have “cancel” or “no” as label, or 0 if there is no such buttons. On macOS and Windows the index of the “Cancel” but	ton will always be used as cancelId even if it is specified.
// 	noLink Boolean (optional) - On Windows Electron will try to figure out which one of the buttons are common buttons (like “Cancel” or “Yes”), and show the others as command links in the dialog. This can make the dialog appear in the style of modern Windows apps. If you don’t like this behavior, you can set noLink to true.
// callback Function (optional)
// 	response Number - The index of the button that was clicked
// Returns Integer, the index of the clicked button, if a callback is provided it returns undefined.
//
// Shows a message box, it will block the process until the message box is closed. It returns the index of the clicked button.
//
// The browserWindow argument allows the dialog to attach itself to a parent window, making it modal.
//
// If a callback is passed, the API call will be asynchronous and the result will be passed via callback(response).
func (d *DialogModule) ShowMessageBoxEx(opt DialogOptionMessage, bw ...*BrowserWindow) (buttonIndex int) {
	var out *js.Object
	if len(bw) > 0 {
		out = d.Call("showMessageBox", bw[0], opt.toJs())
	} else {
		out = d.Call("showMessageBox", opt.toJs())
	}
	// check null value
	return out.Int()
}

// ExShowMessage easy form for ShowMessageBox
func (d *DialogModule) ShowMessageEx(title, msg string) {
	d.ShowMessageBoxEx(DialogOptionMessage{
		Title:   title,
		Type:    DialogTypeInfo,
		Message: msg,
	})
}
