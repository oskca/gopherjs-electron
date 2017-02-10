package electron

import (
	"github.com/gopherjs/gopherjs/js"
)

// Menu.buildFromTemplate(template)
// template MenuItemConstructorOptions[]
// Returns Menu
// Generally, the template is just an array of options for constructing a MenuItem.
// The usage can be referenced above.
// You can also attach other fields to the element of the template and they will become
// properties of the constructed menu items.
//
// !!! You can use the go style struct literal here
func ExBuildFromTemplate(opts []ExMenuItemOption) *Menu {
	o := make(js.S, 0)
	for _, opt := range opts {
		o = append(o, opt.toMap())
	}
	return WrapMenu(electron.Get("Menu").Call("buildFromTemplate", o))
}
