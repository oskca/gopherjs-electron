package electron

import "github.com/gopherjs/gopherjs/js"

// IpcMainModule version@1.4.15
//
// Communicate asynchronously from the main process to renderer processes.
type IpcMainModule struct {
	*js.Object
	// Listens to channel, when a new message arrives listener would be called with listener(event, args...).
	On func(Channel string, Listener IpcMainModuleOnListener) `js:"on"`
	// Adds a one time listener function for the event. This listener is invoked only the next time a message is sent to channel, after which it is removed.
	Once func(Channel string, Listener IpcMainModuleOnceListener) `js:"once"`
	// Removes the specified listener from the listener array for the specified channel.
	RemoveListener func(Channel string, Listener IpcMainModuleRemoveListenerListener) `js:"removeListener"`
	// Removes all listeners, or those of the specified channel.
	RemoveAllListeners func(Channel string) `js:"removeAllListeners"`
}

func GetIpcMainModule() *IpcMainModule {
	o := Get("ipcMain")
	return &IpcMainModule{
		Object: o,
	}
}

type IpcMainModuleOnceListener func()
type IpcMainModuleRemoveListenerListener func()
type IpcMainModuleOnListener func()
