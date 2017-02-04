// Package ipc Communicate asynchronously from a renderer process to the main process.
// Process: Renderer/Main
//
// The ipc module is an instance of the EventEmitter class.
// It provides a few methods so you can send synchronous and asynchronous messages from
// the render process (web page) to the main process. You can also receive replies from the main process.
package ipc

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/webcontents"
	"github.com/oskca/gopherjs-nodejs/eventemitter"
)

type IPC struct {
	*js.Object
	*eventemitter.EventEmitter
}

// Event is usd in IPC
type Event struct {
	*js.Object
	// Event object
	// The event object passed to the callback has the following methods:
	// event.returnValue
	// Set this to the value to be returned in a synchronous message.
	ReturnValue interface{} `js:"returnValue"`

	// event.sender
	// Returns the webContents that sent the message, you can call event.sender.send to reply to the asynchronous message, see webContents.send for more information.
	Sender *webcontents.WebContents `js:"sender"`
}

// Listener is used in ipc
type Listener func(evt *Event, args ...*js.Object)

func newIPC(obj *js.Object) *IPC {
	ipc := new(IPC)
	ipc.Object = obj
	ipc.EventEmitter = eventemitter.New(obj)
	return ipc
}

// Main returns ipcMain
func Main() *IPC {
	return newIPC(electron.Get("ipcMain"))
}

// Renderer returns ipcRenderer
func Renderer() *IPC {
	return newIPC(electron.Get("ipcRenderer"))
}

// ipc

// See ipcMain for code examples.

// Methods

// The ipc module has the following method to listen for events and send messages:

// ipc.on(channel, listener)
// channel String
// listener Function
// Listens to channel, when a new message arrives listener would be called with listener(event, args...).
func (ipc *IPC) On(channel string, listener Listener) {
	ipc.Call("on", channel, listener)
}

// ipc.once(channel, listener)
// channel String
// listener Function
// Adds a one time listener function for the event. This listener is invoked only the next time a message is sent to channel, after which it is removed.
func (ipc *IPC) Once(channel string, listener Listener) {
	ipc.Call("once", channel, listener)
}

// ipc.removeListener(channel, listener)
// channel String
// listener Function
// Removes the specified listener from the listener array for the specified channel.
func (ipc *IPC) RemoveListener(channel string, listener Listener) {
	ipc.Call("removeListener", channel, listener)
}

// ipc.removeAllListeners([channel])
// channel String (optional)
// Removes all listeners, or those of the specified channel.
func (ipc *IPC) RemoveAllListeners(channel ...string) {
	if len(channel) > 0 {
		ipc.Call("removeAllListeners", channel[0])
	} else {
		ipc.Call("removeAllListeners")
	}
}

// ipc.send(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Send a message to the main process asynchronously via channel,
// you can also send arbitrary arguments. Arguments will be serialized in JSON internally and
// hence no functions or prototype chain will be included.
// The main process handles it by listening for channel with ipcMain module.
func (ipc *IPC) Send(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipc.Call("send", n...)
}

// ipc.sendSync(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Send a message to the main process synchronously via channel, you can also send arbitrary arguments.
// Arguments will be serialized in JSON internally and hence no functions or prototype chain will be included.
// The main process handles it by listening for channel with ipcMain module,
// and replies by setting event.returnValue.
// Note: Sending a synchronous message will block the whole renderer process, unless
// you know what you are doing you should never use it.
func (ipc *IPC) SendSync(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipc.Call("sendSync", n...)
}

// ipc.sendToHost(channel[, arg1][, arg2][, ...])
// channel String
// ...args any[]
// Like ipc.send but the event will be sent to the <webview> element in
// the host page instead of the main process.
func (ipc *IPC) SendToHost(channel string, args ...interface{}) {
	n := []interface{}{channel}
	for _, arg := range args {
		n = append(n, arg)
	}
	ipc.Call("sendToHost", n...)
}
