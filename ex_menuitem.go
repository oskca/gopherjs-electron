package electron

import "github.com/gopherjs/gopherjs/js"

// MenuItem roles
const (
	RoleUndo               = "undo"
	RoleRedo               = "redo"
	RoleCut                = "cut"
	RoleCopy               = "copy"
	RolePaste              = "paste"
	RolePasteandmatchstyle = "pasteandmatchstyle"
	RoleSelectall          = "selectall"
	RoleDelete             = "delete"
	// minimize - Minimize current window
	RoleMinimize = "minimize"
	// close - Close current window
	RoleClose = "close"
	// quit- Quit the application
	RoleQuit = "quit-"
	// reload - Reload the current window
	RoleReload = "reload"
	// toggledevtools - Toggle developer tools in the current window
	RoleToggledevtools = "toggledevtools"
	// togglefullscreen- Toggle full screen mode on the current window
	RoleTogglefullscreen = "togglefullscreen-"
	// resetzoom - Reset the focused page’s zoom level to the original size
	RoleResetzoom = "resetzoom"
	// zoomin - Zoom in the focused page by 10%
	RoleZoomin = "zoomin"
	// zoomout - Zoom out the focused page by 10%
	RoleZoomout = "zoomout"

	// On macOS role can also have following additional values:
	// When specifying role on macOS, label and accelerator are the only options that
	// will affect the MenuItem. All other options will be ignored.
	RoleCOS = "cOS"
	// about - Map to the orderFrontStandardAboutPanel action
	RoleAbout = "about"
	// hide - Map to the hide action
	RoleHide = "hide"
	// hideothers - Map to the hideOtherApplications action
	RoleHideothers = "hideothers"
	// unhide - Map to the unhideAllApplications action
	RoleUnhide = "unhide"
	// startspeaking - Map to the startSpeaking action
	RoleStartspeaking = "startspeaking"
	// stopspeaking - Map to the stopSpeaking action
	RoleStopspeaking = "stopspeaking"
	// front - Map to the arrangeInFront action
	RoleFront = "front"
	// zoom - Map to the performZoom action
	RoleZoom = "zoom"
	// window - The submenu is a “Window” menu
	RoleWindow = "window"
	// help - The submenu is a “Help” menu
	RoleHelp = "help"
	// services - The submenu is a “Services” menu
	RoleServices = "services "
)

type ExMenuItemOption struct {
	// *js.Object
	// click Function (optional) - Will be called with click(menuItem, browserWindow, event)
	// when the menu item is clicked.
	//  menuItem MenuItem
	//  browserWindow BrowserWindow
	//  event Event
	// Click func(item *Item, w *browserwindow.BrowserWindow, event *js.Object) `js:"click"`
	Click   func()
	ClickEx func(item *MenuItem)
	// role String (optional) - Define the action of the menu item,
	//      when specified the click property will be ignored.
	Role string
	// type String (optional) - Can be normal, separator, submenu, checkbox or radio.
	Type string
	// label String - (optional)
	Label string
	// sublabel String - (optional)
	Sublabel string
	// accelerator Accelerator (optional)
	// icon (NativeImage | String) (optional)
	Icon *NativeImage
	// enabled Boolean (optional) - If false, the menu item will be greyed out and unclickable.
	Enabled bool
	// visible Boolean (optional) - If false, the menu item will be entirely hidden.
	Visible bool
	// checked Boolean (optional) - Should only be specified for checkbox or radio type menu items.
	Checked bool
	// submenu (MenuItemConstructorOptions[] | Menu) (optional) -
	//      Should be specified for submenu type menu items.
	//      If submenu is specified, the type: 'submenu' can be omitted.
	//      If the value is not a Menu then it will be automatically converted to one using Menu.buildFromTemplate.
	SubMenuOptions []ExMenuItemOption
	SubMenu        *Menu
	// id String (optional) - Unique within a single menu.
	// If defined then it can be used as a reference to this item by the position attribute.
	ID string
	// position String (optional) - This field allows fine-grained definition of
	// the specific location within a given menu.
	Position string
}

func (o *ExMenuItemOption) toMap() js.M {
	m := make(js.M)
	if o.Click != nil {
		m["click"] = o.Click
	}
	if o.ClickEx != nil {
		m["click"] = o.ClickEx
	}
	if o.Role != "" {
		m["role"] = o.Role
	}
	if o.Type != "" {
		m["type"] = o.Type
	}
	if o.Label != "" {
		m["label"] = o.Label
	}
	if o.Sublabel != "" {
		m["sublabel"] = o.Sublabel
	}
	if o.Icon != nil {
		m["icon"] = o.Icon
	}
	if o.Enabled != false {
		m["enabled"] = o.Enabled
	}
	if o.Visible != false {
		m["visible"] = o.Visible
	}
	if o.Checked != false {
		m["checked"] = o.Checked
	}
	if o.SubMenuOptions != nil {
		subm := make(js.S, 0)
		for _, opt := range o.SubMenuOptions {
			subm = append(subm, opt.toMap())
		}
		m["submenu"] = subm
	}
	if o.SubMenu != nil {
		m["submenu"] = o.SubMenu
	}
	if o.ID != "" {
		m["id"] = o.ID
	}
	if o.Position != "" {
		m["position"] = o.Position
	}
	return m
}

func ExNewItem(opt ExMenuItemOption) *MenuItem {
	o := electron.Get("Menu").Call("MenuItem", opt.toMap())
	return &MenuItem{
		Object: o,
	}
}
