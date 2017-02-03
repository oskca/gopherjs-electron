package menu

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
)

var menu = electron.Get("Menu")

// Menu The menu object has the following instance methods:
type Menu struct {
	*js.Object
	// Instance Methods
	// menu.popup([browserWindow, x, y, positioningItem])
	// browserWindow BrowserWindow (optional) - Default is BrowserWindow.getFocusedWindow().
	// x Number (optional) - Default is the current mouse cursor position.
	// y Number (required if x is used) - Default is the current mouse cursor position.
	// positioningItem Number (optional) macOS - The index of the menu item to be positioned under the mouse cursor at the specified coordinates. Default is -1.
	// Pops up this menu as a context menu in the browserWindow.
	Popup   func()         `js:"popup"`
	PopupEx func(x, y int) `js:"popup"`

	// menu.append(menuItem)
	// menuItem MenuItem
	// Appends the menuItem to the menu.
	Append func(item *Item) `js:"append"`

	// menu.insert(pos, menuItem)
	// pos Integer
	// menuItem MenuItem
	// Inserts the menuItem to the pos position of the menu.
	Insert func(pos int, item *Item) `js:"insert"`

	// Instance Properties
	// menu objects also have the following properties:

	// menu.items
	// A MenuItem[] array containing the menu’s items.
	Items []*Item `js:"items"`

	// Each Menu consists of multiple MenuItems and each MenuItem can have a submenu.
}

// New Creates a new menu.
func New() *Menu {
	return &Menu{
		Object: menu.New(),
	}
}

// Static Methods
// The menu class has the following static methods:

// Menu.setApplicationMenu(menu)
// menu Menu
// Sets menu as the application menu on macOS. On Windows and Linux, the menu will be set as each window’s top menu.
// Note: This API has to be called after the ready event of app module.
func SetApplicationMenu(m *Menu) {
	menu.Call("setApplicationMenu", m)
}

// GetApplicationMenu Returns Menu - The application menu, if set, or null, if not set.
func GetApplicationMenu() *Menu {
	return &Menu{
		Object: menu.Call("getApplicationMenu"),
	}
}

// Menu.sendActionToFirstResponder(action) macOS
// action String
// Sends the action to the first responder of application. This is used for emulating default Cocoa menu behaviors, usually you would just use the role property of MenuItem.

// See the macOS Cocoa Event Handling Guide for more information on macOS’ native actions.

// Menu.buildFromTemplate(template)
// template MenuItemConstructorOptions[]
// Returns Menu
// Generally, the template is just an array of options for constructing a MenuItem.
// The usage can be referenced above.
// You can also attach other fields to the element of the template and they will become
// properties of the constructed menu items.
//
// !!! You can use the go style struct literal here
func BuildFromTemplate(opts []Option) *Menu {
	o := make(js.S, 0)
	for _, opt := range opts {
		o = append(o, opt.toMap())
	}
	return &Menu{
		Object: menu.Call("buildFromTemplate", o),
	}
}
