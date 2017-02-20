package electron

import "github.com/gopherjs/gopherjs/js"

// Menu version@1.4.15
//
// Create native application menus and context menus.
type Menu struct {
	*js.Object
	// A MenuItem[] array containing the menu's items. Each Menu consists of multiple MenuItems and each MenuItem can have a submenu.
	Items *js.Object `js:"items"`
	// Pops up this menu as a context menu in the browserWindow.
	Popup func(BrowserWindow *BrowserWindow, X float64, Y float64, PositioningItem float64) `js:"popup"`
	// Appends the menuItem to the menu.
	Append func(MenuItem *MenuItem) `js:"append"`
	// Inserts the menuItem to the pos position of the menu.
	Insert func(Pos int64, MenuItem *MenuItem) `js:"insert"`
}

func WrapMenu(o *js.Object) *Menu {
	return &Menu{
		Object: o,
	}
}

func SetApplicationMenu(Menu *Menu) {
	o := electron.Get("Menu")
	o.Call("setApplicationMenu", Menu)
}
func GetApplicationMenu() *js.Object {
	o := electron.Get("Menu")
	ret := o.Call("getApplicationMenu")
	return ret
}
func SendActionToFirstResponder(Action string) {
	o := electron.Get("Menu")
	o.Call("sendActionToFirstResponder", Action)
}
func BuildFromTemplate(Template *js.Object) *js.Object {
	o := electron.Get("Menu")
	ret := o.Call("buildFromTemplate", Template)
	return ret
}
func NewMenu() *Menu {
	o := electron.Get("Menu")
	ret := o.New()
	return WrapMenu(ret)
}
