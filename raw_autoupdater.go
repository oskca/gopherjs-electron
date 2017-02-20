package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

const (
	// Emitted when there is an error while updating.
	EvtAutoUpdaterError = "error"
	// Emitted when checking if an update has started.
	EvtAutoUpdaterCheckingForUpdate = "checking-for-update"
	// Emitted when there is an available update. The update is downloaded automatically.
	EvtAutoUpdaterUpdateAvailable = "update-available"
	// Emitted when there is no available update.
	EvtAutoUpdaterUpdateNotAvailable = "update-not-available"
	// Emitted when an update has been downloaded. On Windows only releaseName is available.
	EvtAutoUpdaterUpdateDownloaded = "update-downloaded"
)

// AutoUpdaterModule version@1.4.15
//
// Enable apps to automatically update themselves.
type AutoUpdaterModule struct {
	*events.Emitter
	// Sets the url and initialize the auto updater.
	SetFeedURL func(URL string, RequestHeaders *AutoUpdaterModuleSetFeedURLRequestHeaders) `js:"setFeedURL"`
	GetFeedURL func() (Obj string)                                                         `js:"getFeedURL"`
	// Asks the server whether there is an update. You must call setFeedURL before using this API.
	CheckForUpdates func() `js:"checkForUpdates"`
	// Restarts the app and installs the update after it has been downloaded. It should only be called after update-downloaded has been emitted. Note: autoUpdater.quitAndInstall() will close all application windows first and only emit before-quit event on app after that. This is different from the normal quit event sequence.
	QuitAndInstall func() `js:"quitAndInstall"`
}

func GetAutoUpdaterModule() *AutoUpdaterModule {
	o := Get("autoUpdater")
	return &AutoUpdaterModule{
		Emitter: events.New(o),
	}
}

type AutoUpdaterModuleSetFeedURLRequestHeaders struct {
	*js.Object
}
