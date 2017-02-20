package electron

import "github.com/gopherjs/gopherjs/js"

// ShellModule version@1.4.15
//
// Manage files and URLs using their default applications.
type ShellModule struct {
	*js.Object
	// Show the given file in a file manager. If possible, select the file.
	ShowItemInFolder func(FullPath string) (Obj bool) `js:"showItemInFolder"`
	// Open the given file in the desktop's default manner.
	OpenItem func(FullPath string) (Obj bool) `js:"openItem"`
	// Open the given external protocol URL in the desktop's default manner. (For example, mailto: URLs in the user's default mail agent).
	OpenExternal func(URL string, Options *ShellModuleOpenExternalOptions, Callback ShellModuleOpenExternalCallback) (Obj bool) `js:"openExternal"`
	// Move the given file to trash and returns a boolean status for the operation.
	MoveItemToTrash func(FullPath string) (Obj bool) `js:"moveItemToTrash"`
	// Play the beep sound.
	Beep func() `js:"beep"`
	// Creates or updates a shortcut link at shortcutPath.
	WriteShortcutLink func(ShortcutPath string, Operation ShellModuleWriteShortcutLinkOperation, Options *js.Object) (Obj bool) `js:"writeShortcutLink"`
	// Resolves the shortcut link at shortcutPath. An exception will be thrown when any error happens.
	ReadShortcutLink func(ShortcutPath string) (Obj *js.Object) `js:"readShortcutLink"`
}

func GetShellModule() *ShellModule {
	o := Get("shell")
	return &ShellModule{
		Object: o,
	}
}

type ShellModuleOpenExternalOptions struct {
	*js.Object
	// to bring the opened application to the foreground. The default is .
	Activate bool `js:"activate"`
}

type ShellModuleOpenExternalCallback func(Error *js.Object)
type ShellModuleWriteShortcutLinkOperation string

// consts
const (
	ShellModuleWriteShortcutLinkOperationCreate  ShellModuleWriteShortcutLinkOperation = "create"
	ShellModuleWriteShortcutLinkOperationUpdate  ShellModuleWriteShortcutLinkOperation = "update"
	ShellModuleWriteShortcutLinkOperationReplace ShellModuleWriteShortcutLinkOperation = "replace"
)
