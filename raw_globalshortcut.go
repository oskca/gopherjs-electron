package electron

import "github.com/gopherjs/gopherjs/js"

// GlobalShortcutModule version@1.4.15
//
// Detect keyboard events when the application does not have keyboard focus.
type GlobalShortcutModule struct {
	*js.Object
	// Registers a global shortcut of accelerator. The callback is called when the registered shortcut is pressed by the user. When the accelerator is already taken by other applications, this call will silently fail. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
	Register func(Accelerator *js.Object, Callback GlobalShortcutModuleRegisterCallback) `js:"register"`
	// When the accelerator is already taken by other applications, this call will still return false. This behavior is intended by operating systems, since they don't want applications to fight for global shortcuts.
	IsRegistered func(Accelerator *js.Object) (Obj bool) `js:"isRegistered"`
	// Unregisters the global shortcut of accelerator.
	Unregister func(Accelerator *js.Object) `js:"unregister"`
	// Unregisters all of the global shortcuts.
	UnregisterAll func() `js:"unregisterAll"`
}

func GetGlobalShortcutModule() *GlobalShortcutModule {
	o := Get("globalShortcut")
	return &GlobalShortcutModule{
		Object: o,
	}
}

type GlobalShortcutModuleRegisterCallback func()
