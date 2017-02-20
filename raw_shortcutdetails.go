package electron

import "github.com/gopherjs/gopherjs/js"

// ShortcutDetails a Structure
type ShortcutDetails struct {
	*js.Object
	// The target to launch from this shortcut.
	Target string `js:"target"`
	// The working directory. Default is empty.
	Cwd string `js:"cwd"`
	// The arguments to be applied to when launching from this shortcut. Default is empty.
	Args string `js:"args"`
	// The description of the shortcut. Default is empty.
	Description string `js:"description"`
	// The path to the icon, can be a DLL or EXE. and have to be set together. Default is empty, which uses the target's icon.
	Icon string `js:"icon"`
	// The resource ID of icon when is a DLL or EXE. Default is 0.
	IconIndex float64 `js:"iconIndex"`
	// The Application User Model ID. Default is empty.
	AppUserModelId string `js:"appUserModelId"`
}
