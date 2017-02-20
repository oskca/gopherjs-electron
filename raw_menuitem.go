package electron

import "github.com/gopherjs/gopherjs/js"

// MenuItem version@1.4.15
//
// Add items to native application menus and context menus.
type MenuItem struct {
	*js.Object
	// A Boolean indicating whether the item is enabled, this property can be dynamically changed.
	Enabled bool `js:"enabled"`
	// A Boolean indicating whether the item is visible, this property can be dynamically changed.
	Visible bool `js:"visible"`
	// A Boolean indicating whether the item is checked, this property can be dynamically changed. A checkbox menu item will toggle the checked property on and off when selected. A radio menu item will turn on its checked property when clicked, and will turn off that property for all adjacent items in the same menu. You can add a click function for additional behavior.
	Checked bool `js:"checked"`
	// A String representing the menu items visible label
	Label string `js:"label"`
	// A Function that is fired when the MenuItem recieves a click event
	Click MenuItemMenuItemClick `js:"click"`
}

func WrapMenuItem(o *js.Object) *MenuItem {
	return &MenuItem{
		Object: o,
	}
}

func NewMenuItem(Options *MenuItemMenuItemOptions) *MenuItem {
	o := electron.Get("MenuItem")
	ret := o.New(Options)
	return WrapMenuItem(ret)
}

type MenuItemMenuItemOptions struct {
	*js.Object
	// Will be called with when the menu item is clicked.
	Click MenuItemOptionsClick `js:"click"`
	// Define the action of the menu item, when specified the property will be ignored.
	Role string `js:"role"`
	// Can be , , , or .
	Type MenuItemOptionsType `js:"type"`
	// (optional)
	Label string `js:"label"`
	// (optional)
	Sublabel    string       `js:"sublabel"`
	Accelerator *js.Object   `js:"accelerator"`
	Icon        *NativeImage `js:"icon"`
	// If false, the menu item will be greyed out and unclickable.
	Enabled bool `js:"enabled"`
	// If false, the menu item will be entirely hidden.
	Visible bool `js:"visible"`
	// Should only be specified for or type menu items.
	Checked bool `js:"checked"`
	// Should be specified for type menu items. If is specified, the can be omitted. If the value is not a then it will be automatically converted to one using .
	Submenu *js.Object `js:"submenu"`
	// Unique within a single menu. If defined then it can be used as a reference to this item by the position attribute.
	Id string `js:"id"`
	// This field allows fine-grained definition of the specific location within a given menu.
	Position string `js:"position"`
}

type MenuItemOptionsClick func(MenuItem *MenuItem, BrowserWindow *BrowserWindow, Event *js.Object)
type MenuItemOptionsType string

// consts
const (
	MenuItemOptionsTypeNormal    MenuItemOptionsType = "normal"
	MenuItemOptionsTypeSeparator MenuItemOptionsType = "separator"
	MenuItemOptionsTypeSubmenu   MenuItemOptionsType = "submenu"
	MenuItemOptionsTypeCheckbox  MenuItemOptionsType = "checkbox"
	MenuItemOptionsTypeRadio     MenuItemOptionsType = "radio"
)

type MenuItemMenuItemClick func()
