// Package ipc Communicate asynchronously from a renderer process to the main process.
// Process: Renderer
//
// The ipcRenderer module is an instance of the EventEmitter class.
// It provides a few methods so you can send synchronous and asynchronous messages from
// the render process (web page) to the main process. You can also receive replies from the main process.
package ipc

import (
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api"
)

var ipcRenderer = electron.Get("ipcRenderer")

// ipcRenderer

// See ipcMain for code examples.

// Methods

// The ipcRenderer module has the following method to listen for events and send messages:

// ipcRenderer.on(channel, listener)
// channel String
// listener Function
// Listens to channel, when a new message arrives listener would be called with listener(event, args...).
func On(channel string, listener api.IpcListener) {
	ipcRenderer.Call("on", channel, listener)
}

// ipcRenderer.once(channel, listener)
// channel String
// listener Function
// Adds a one time listener function for the event. This listener is invoked only the next time a message is sent to channel, after which it is removed.
func Once(channel string, listener api.IpcListener) {
	ipcRenderer.Call("once", channel, listener)
}

// ipcRenderer.removeListener(channel, listener)
// channel String
// listener Function
// Removes the specified listener from the listener array for the specified channel.
func RemoveListener(channel string, listener api.IpcListener) {
	ipcRenderer.Call("removeListener", channel, listener)
}

// ipcRenderer.removeAllListeners([channel])
// channel String (optional)
// Removes all listeners, or those of the specified channel.
func RemoveAllListeners(channel ...string) {
	if len(channel) > 0 {
		ipcRenderer.Call("removeAllListeners", channel[0])
	} else {
		ipcRenderer.Call("removeAllListeners")
	}
}

// ipcRenderer.send(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Send a message to the main process asynchronously via channel,
// you can also send arbitrary arguments. Arguments will be serialized in JSON internally and
// hence no functions or prototype chain will be included.
// The main process handles it by listening for channel with ipcMain module.
func Send(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipcRenderer.Call("send", n...)
}

// ipcRenderer.sendSync(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Send a message to the main process synchronously via channel, you can also send arbitrary arguments.
// Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included.
// The main process handles it by listening for channel with ipcMain module,
// and replies by setting event.returnValue.
// Note: Sending a synchronous message will block the whole renderer process, unless
// you know what you are doing you should never use it.
func SendSync(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipcRenderer.Call("sendSync", n...)
}

// ipcRenderer.sendToHost(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Like ipcRenderer.send but the event will be sent to the <webview> element in
// the host page instead of the main process.
func SendToHost(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipcRenderer.Call("sendToHost", n...)
}
