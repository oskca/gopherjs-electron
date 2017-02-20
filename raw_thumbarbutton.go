package electron

import "github.com/gopherjs/gopherjs/js"

// ThumbarButton a Structure
type ThumbarButton struct {
	*js.Object
	// The icon showing in thumbnail toolbar.
	Icon  *NativeImage                    `js:"icon"`
	Click ThumbarButtonThumbarButtonClick `js:"click"`
	// The text of the button's tooltip.
	Tooltip string `js:"tooltip"`
	// Control specific states and behaviors of the button. By default, it is .
	Flags *js.Object `js:"flags"`
}

type ThumbarButtonThumbarButtonClick func()
