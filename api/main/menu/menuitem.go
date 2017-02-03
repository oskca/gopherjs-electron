package menu

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-electron/api/allprocess/nativeimage"
)

// Item types
const (
	TypeNormal    = "normal"
	TypeSeparator = "separator"
	TypeSubmenu   = "submenu"
	TypeCheckbox  = "checkbox "
	TypeRadio     = "radio"
)

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

type Option struct {
	// *js.Object
	// click Function (optional) - Will be called with click(menuItem, browserWindow, event)
	// when the menu item is clicked.
	//  menuItem MenuItem
	//  browserWindow BrowserWindow
	//  event Event
	// Click func(item *Item, w *browserwindow.BrowserWindow, event *js.Object) `js:"click"`
	Click   func()
	ClickEx func(item *Item)
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
	Icon *nativeimage.NativeImage
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
	SubMenuOptions []Option
	SubMenu        *Menu
	// id String (optional) - Unique within a single menu.
	// If defined then it can be used as a reference to this item by the position attribute.
	ID string
	// position String (optional) - This field allows fine-grained definition of
	// the specific location within a given menu.
	Position string
}

func (o *Option) toMap() js.M {
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

type Item struct {
	*js.Object
	// menuItem.enabled
	// A Boolean indicating whether the item is enabled, this property can be dynamically changed.
	Enabled bool `js:"enabled"`

	// menuItem.visible
	// A Boolean indicating whether the item is visible, this property can be dynamically changed.
	Visible bool `js:"visible"`

	// menuItem.checked
	// A Boolean indicating whether the item is checked, this property can be dynamically changed.
	// A checkbox menu item will toggle the checked property on and off when selected.
	// A radio menu item will turn on its checked property when clicked,
	// and will turn off that property for all adjacent items in the same menu.
	// You can add a click function for additional behavior.
	Checked bool `js:"checked"`

	// menuItem.label
	// A String representing the menu items visible label
	Label string `js:"label"`

	// menuItem.click
	// A Function that is fired when the MenuItem recieves a click event
	Click func() `js:"click"`
}

func NewItem(opt Option) *Item {
	o := menu.Call("MenuItem", opt.toMap())
	return &Item{
		Object: o,
	}
}
