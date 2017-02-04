// Package remote Use main process modules from the renderer process.
// Process: Renderer
//
// The remote module provides a simple way to do inter-process communication (IPC) between the renderer process (web page) and the main process.
//
// In Electron, GUI-related modules (such as dialog, menu etc.) are only available in the main process,
// not in the renderer process. In order to use them from the renderer process, the ipc module is necessary to
// send inter-process messages to the main process. With the remote module, you can invoke methods of
// the main process object without explicitly sending inter-process messages, similar to Java’s RMI.
// An example of creating a browser window from a renderer process:
//
//  const {BrowserWindow} = require('electron').remote
//  let win = new BrowserWindow({width: 800, height: 600})
//  win.loadURL('https://github.com')
//
// Note: For the reverse (access the renderer process from the main process),
// you can use webContents.executeJavascript.
package remote

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/browserwindow"
	"github.com/oskca/gopherjs-electron/api/main/webcontents"
)

var (
	remote = js.Global.Get("electron").Get("remote")
)

// GetCurrentWindow remote.getCurrentWindow()
// Returns BrowserWindow - The window to which this web page belongs.
func GetCurrentWindow() *browserwindow.BrowserWindow {
	return &browserwindow.BrowserWindow{
		Object: remote.Call("getCurrentWindow"),
	}
}

// GetCurrentWebContents remote.getCurrentWebContents()
// Returns WebContents - The web contents of this web page.
func GetCurrentWebContents() *webcontents.WebContents {
	return &webcontents.WebContents{
		Object: remote.Call("getCurrentWebContents"),
	}
}

// GetGlobal remote.getGlobal(name)
// name String
// Returns any - The global variable of name (e.g. global[name]) in the main process.
func GetGlobal(name string) *js.Object {
	return remote.Call("getGlobal", name)
}

func init() {
	electron.UseRemote()
}
