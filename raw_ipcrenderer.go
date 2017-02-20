package electron

import "github.com/gopherjs/gopherjs/js"

// IpcRendererModule version@1.4.15
//
// Communicate asynchronously from a renderer process to the main process.
type IpcRendererModule struct {
	*js.Object
	// Listens to channel, when a new message arrives listener would be called with listener(event, args...).
	On func(Channel string, Listener IpcRendererModuleOnListener) `js:"on"`
	// Adds a one time listener function for the event. This listener is invoked only the next time a message is sent to channel, after which it is removed.
	Once func(Channel string, Listener IpcRendererModuleOnceListener) `js:"once"`
	// Removes the specified listener from the listener array for the specified channel.
	RemoveListener func(Channel string, Listener IpcRendererModuleRemoveListenerListener) `js:"removeListener"`
	// Removes all listeners, or those of the specified channel.
	RemoveAllListeners func(Channel string) `js:"removeAllListeners"`
	// Send a message to the main process asynchronously via channel, you can also send arbitrary arguments. Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included. The main process handles it by listening for channel with ipcMain module.
	Send func(Channel string, Args *js.Object) `js:"send"`
	// Send a message to the main process synchronously via channel, you can also send arbitrary arguments. Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included. The main process handles it by listening for channel with ipcMain module, and replies by setting event.returnValue. Note: Sending a synchronous message will block the whole renderer process, unless you know what you are doing you should never use it.
	SendSync func(Channel string, Args *js.Object) `js:"sendSync"`
	// Like ipcRenderer.send but the event will be sent to the  element in the host page instead of the main process.
	SendToHost func(Channel string, Args *js.Object) `js:"sendToHost"`
}

func GetIpcRendererModule() *IpcRendererModule {
	o := Get("ipcRenderer")
	return &IpcRendererModule{
		Object: o,
	}
}

type IpcRendererModuleOnListener func()
type IpcRendererModuleOnceListener func()
type IpcRendererModuleRemoveListenerListener func()
