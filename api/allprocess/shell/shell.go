// Package shell Manage files and URLs using their default applications.
// Process: Main, Renderer
// The shell module provides functions related to desktop integration.
// An example of opening a URL in the user’s default browser:
//
//  const {shell} = require('electron')
//  shell.openExternal('https://github.com')
package shell

import (
	"github.com/gopherjs/gopherjs/js"
)

// operations
const (
	OpCreate  = "create"  // create - Creates a new shortcut, overwriting if necessary.
	OpUpdate  = "update"  // update - Updates specified properties only on an existing shortcut.
	OpReplace = "replace" // replace - Overwrites an existing shortcut, fails if the shortcut doesn’t exist.
)

// ShortcutDetails Object
type ShortcutDetails struct {
	*js.Object
	Target         string `js:"target "`         // target String - The target to launch from this shortcut.
	Cwd            string `js:"cwd "`            // cwd String (optional) - The working directory. Default is empty.
	Args           string `js:"args "`           // args String (optional) - The arguments to be applied to target when launching from this shortcut. Default is empty.
	Description    string `js:"description "`    // description String (optional) - The description of the shortcut. Default is empty.
	Icon           string `js:"icon "`           // icon String (optional) - The path to the icon, can be a DLL or EXE. icon and iconIndex have to be set together. Default is empty, which uses the target’s icon.
	IconIndex      int    `js:"iconIndex "`      // iconIndex Number (optional) - The resource ID of icon when icon is a DLL or EXE. Default is 0.
	AppUserModelID string `js:"appUserModelId "` // appUserModelId String (optional) - The Application User Model ID. Default is empty.
}

func (s ShortcutDetails) toJs() js.M {
	m := make(js.M)
	if s.Target != "" {
		m["target"] = s.Target
	}
	if s.Cwd != "" {
		m["cwd"] = s.Cwd
	}
	if s.Args != "" {
		m["args"] = s.Args
	}
	if s.Description != "" {
		m["description"] = s.Description
	}
	if s.Icon != "" {
		m["icon"] = s.Icon
	}
	m["iconIndex"] = s.IconIndex
	if s.AppUserModelID != "" {
		m["appUserModelID"] = s.AppUserModelID
	}
	return m
}

// Methods
// The shell module has the following methods:

// Shell wraps
type Shell struct {
	*js.Object
	// shell.showItemInFolder(fullPath)
	// fullPath String
	// Returns Boolean - Whether the item was successfully shown
	// Show the given file in a file manager. If possible, select the file.
	ShowItemInFolder func(fullPath string) bool `js:"showItemInFolder"`

	// shell.openItem(fullPath)
	// fullPath String
	// Returns Boolean - Whether the item was successfully opened.
	// Open the given file in the desktop’s default manner.
	OpenItem func(fullPath string) bool `js:"openItem"`

	// shell.openExternal(url[, options, callback])
	// url String
	// options Object (optional) macOS
	// activate Boolean - true to bring the opened application to the foreground. The default is true.
	// callback Function (optional) - If specified will perform the open asynchronously. macOS
	// error Error
	// Returns Boolean - Whether an application was available to open the URL. If callback is specified, always returns true.
	// Open the given external protocol URL in the desktop’s default manner. (For example, mailto: URLs in the user’s default mail agent).
	OpenExternal func(url string) bool `js:"openExternal"`

	// shell.moveItemToTrash(fullPath)
	// fullPath String
	// Returns Boolean - Whether the item was successfully moved to the trash
	// Move the given file to trash and returns a boolean status for the operation.
	MoveItemToTrash func(fullPath string) bool `js:"moveItemToTrash"`

	// shell.beep()
	// Play the beep sound.
	Beep func() `js:"beep"`

	// shell.writeShortcutLink(shortcutPath[, operation], options) Windows
	// shortcutPath String
	// operation String (optional) - Default is create, can be one of following:
	//  create - Creates a new shortcut, overwriting if necessary.
	//  update - Updates specified properties only on an existing shortcut.
	//  replace - Overwrites an existing shortcut, fails if the shortcut doesn’t exist.
	// options ShortcutDetails
	// Returns Boolean - Whether the shortcut was created successfully
	// Creates or updates a shortcut link at shortcutPath.

	// shell.readShortcutLink(shortcutPath) Windows
	// shortcutPath String
	// Returns ShortcutDetails

	// Resolves the shortcut link at shortcutPath.

	// An exception will be thrown when any error happens.
}

// WriteShortcutLink Creates or updates a shortcut link at shortcutPath.
func (s *Shell) WriteShortcutLink(shortcutPath string, detail ShortcutDetails, operation ...string) bool {
	op := OpCreate
	if len(operation) > 0 {
		op = operation[0]
	}
	out := s.Call("writeShortcutLink", shortcutPath, op, detail.toJs())
	return out.Bool()
}

// ReadShortcutLink Resolves the shortcut link at shortcutPath.
func (s *Shell) ReadShortcutLink(shortcutPath string) *ShortcutDetails {
	out := s.Call("readShortcutLink", shortcutPath)
	return &ShortcutDetails{
		Object: out,
	}
}
