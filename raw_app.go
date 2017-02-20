package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

const (
	// Emitted when the application has finished basic startup. On Windows and Linux, the will-finish-launching event is the same as the ready event; on macOS, this event represents the applicationWillFinishLaunching notification of NSApplication. You would usually set up listeners for the open-file and open-url events here, and start the crash reporter and auto updater. In most cases, you should just do everything in the ready event handler.
	EvtAppWillFinishLaunching = "will-finish-launching"
	// Emitted when Electron has finished initializing. On macOS, launchInfo holds the userInfo of the NSUserNotification that was used to open the application, if it was launched from Notification Center. You can call app.isReady() to check if this event has already fired.
	EvtAppReady = "ready"
	// Emitted when all windows have been closed. If you do not subscribe to this event and all windows are closed, the default behavior is to quit the app; however, if you subscribe, you control whether the app quits or not. If the user pressed Cmd + Q, or the developer called app.quit(), Electron will first try to close all the windows and then emit the will-quit event, and in this case the window-all-closed event would not be emitted.
	EvtAppWindowAllClosed = "window-all-closed"
	// Emitted before the application starts closing its windows. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application. Note: If application quit was initiated by autoUpdater.quitAndInstall() then before-quit is emitted after emitting close event on all windows and closing them.
	EvtAppBeforeQuit = "before-quit"
	// Emitted when all windows have been closed and the application will quit. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application. See the description of the window-all-closed event for the differences between the will-quit and window-all-closed events.
	EvtAppWillQuit = "will-quit"
	// Emitted when the application is quitting.
	EvtAppQuit = "quit"
	// Emitted when the user wants to open a file with the application. The open-file event is usually emitted when the application is already open and the OS wants to reuse the application to open the file. open-file is also emitted when a file is dropped onto the dock and the application is not yet running. Make sure to listen for the open-file event very early in your application startup to handle this case (even before the ready event is emitted). You should call event.preventDefault() if you want to handle this event. On Windows, you have to parse process.argv (in the main process) to get the filepath.
	EvtAppOpenFile = "open-file"
	// Emitted when the user wants to open a URL with the application. Your application's Info.plist file must define the url scheme within the CFBundleURLTypes key, and set NSPrincipalClass to AtomApplication. You should call event.preventDefault() if you want to handle this event.
	EvtAppOpenURL = "open-url"
	// Emitted when the application is activated, which usually happens when the user clicks on the application's dock icon.
	EvtAppActivate = "activate"
	// Emitted during Handoff when an activity from a different device wants to be resumed. You should call event.preventDefault() if you want to handle this event. A user activity can be continued only in an app that has the same developer Team ID as the activity's source app and that supports the activity's type. Supported activity types are specified in the app's Info.plist under the NSUserActivityTypes key.
	EvtAppContinueActivity = "continue-activity"
	// Emitted when a browserWindow gets blurred.
	EvtAppBrowserWindowBlur = "browser-window-blur"
	// Emitted when a browserWindow gets focused.
	EvtAppBrowserWindowFocus = "browser-window-focus"
	// Emitted when a new browserWindow is created.
	EvtAppBrowserWindowCreated = "browser-window-created"
	// Emitted when a new webContents is created.
	EvtAppWebContentsCreated = "web-contents-created"
	// Emitted when failed to verify the certificate for url, to trust the certificate you should prevent the default behavior with event.preventDefault() and call callback(true).
	EvtAppCertificateError = "certificate-error"
	// Emitted when a client certificate is requested. The url corresponds to the navigation entry requesting the client certificate and callback can be called with an entry filtered from the list. Using event.preventDefault() prevents the application from using the first certificate from the store.
	EvtAppSelectClientCertificate = "select-client-certificate"
	// Emitted when webContents wants to do basic auth. The default behavior is to cancel all authentications, to override this you should prevent the default behavior with event.preventDefault() and call callback(username, password) with the credentials.
	EvtAppLogin = "login"
	// Emitted when the gpu process crashes or is killed.
	EvtAppGpuProcessCrashed = "gpu-process-crashed"
	// Emitted when Chrome's accessibility support changes. This event fires when assistive technologies, such as screen readers, are enabled or disabled. See https://www.chromium.org/developers/design-documents/accessibility for more details.
	EvtAppAccessibilitySupportChanged = "accessibility-support-changed"
)

// AppModule version@1.4.15
//
// Control your application's event lifecycle.
type AppModule struct {
	*events.Emitter
	CommandLine *AppModuleAppModuleCommandLine `js:"commandLine"`
	Dock        *AppModuleAppModuleDock        `js:"dock"`
	// Try to close all windows. The before-quit event will be emitted first. If all windows are successfully closed, the will-quit event will be emitted and by default the application will terminate. This method guarantees that all beforeunload and unload event handlers are correctly executed. It is possible that a window cancels the quitting by returning false in the beforeunload event handler.
	Quit func() `js:"quit"`
	// Exits immediately with exitCode.  exitCode defaults to 0. All windows will be closed immediately without asking user and the before-quit and will-quit events will not be emitted.
	Exit func(ExitCode int64) `js:"exit"`
	// Relaunches the app when current instance exits. By default the new instance will use the same working directory and command line arguments with current instance. When args is specified, the args will be passed as command line arguments instead. When execPath is specified, the execPath will be executed for relaunch instead of current app. Note that this method does not quit the app when executed, you have to call app.quit or app.exit after calling app.relaunch to make the app restart. When app.relaunch is called for multiple times, multiple instances will be started after current instance exited. An example of restarting current instance immediately and adding a new command line argument to the new instance:
	Relaunch func(Options *AppModuleRelaunchOptions) `js:"relaunch"`
	IsReady  func() (Obj bool)                       `js:"isReady"`
	// On Linux, focuses on the first visible window. On macOS, makes the application the active app. On Windows, focuses on the application's first window.
	Focus func() `js:"focus"`
	// Hides all application windows without minimizing them.
	Hide func() `js:"hide"`
	// Shows application windows after they were hidden. Does not automatically focus them.
	Show       func()              `js:"show"`
	GetAppPath func() (Obj string) `js:"getAppPath"`
	// You can request the following paths by the name:
	GetPath func(Name string) (Obj string) `js:"getPath"`
	// Overrides the path to a special directory or file associated with name. If the path specifies a directory that does not exist, the directory will be created by this method. On failure an Error is thrown. You can only override paths of a name defined in app.getPath. By default, web pages' cookies and caches will be stored under the userData directory. If you want to change this location, you have to override the userData path before the ready event of the app module is emitted.
	SetPath    func(Name string, Path string) `js:"setPath"`
	GetVersion func() (Obj string)            `js:"getVersion"`
	// Usually the name field of package.json is a short lowercased name, according to the npm modules spec. You should usually also specify a productName field, which is your application's full capitalized name, and which will be preferred over name by Electron.
	GetName func() (Obj string) `js:"getName"`
	// Overrides the current application's name.
	SetName func(Name string) `js:"setName"`
	// Note: When distributing your packaged app, you have to also ship the locales folder. Note: On Windows you have to call it after the ready events gets emitted.
	GetLocale func() (Obj string) `js:"getLocale"`
	// Adds path to the recent documents list. This list is managed by the OS. On Windows you can visit the list from the task bar, and on macOS you can visit it from dock menu.
	AddRecentDocument func(Path string) `js:"addRecentDocument"`
	// Clears the recent documents list.
	ClearRecentDocuments func() `js:"clearRecentDocuments"`
	// This method sets the current executable as the default handler for a protocol (aka URI scheme). It allows you to integrate your app deeper into the operating system. Once registered, all links with your-protocol:// will be opened with the current executable. The whole link, including protocol, will be passed to your application as a parameter. On Windows you can provide optional parameters path, the path to your executable, and args, an array of arguments to be passed to your executable when it launches. Note: On macOS, you can only register protocols that have been added to your app's info.plist, which can not be modified at runtime. You can however change the file with a simple text editor or script during build time. Please refer to Apple's documentation for details. The API uses the Windows Registry and LSSetDefaultHandlerForURLScheme internally.
	SetAsDefaultProtocolClient func(Protocol string, Path string, Args *js.Object) (Obj bool) `js:"setAsDefaultProtocolClient"`
	// This method checks if the current executable as the default handler for a protocol (aka URI scheme). If so, it will remove the app as the default handler.
	RemoveAsDefaultProtocolClient func(Protocol string, Path string, Args *js.Object) (Obj bool) `js:"removeAsDefaultProtocolClient"`
	// This method checks if the current executable is the default handler for a protocol (aka URI scheme). If so, it will return true. Otherwise, it will return false. Note: On macOS, you can use this method to check if the app has been registered as the default protocol handler for a protocol. You can also verify this by checking ~/Library/Preferences/com.apple.LaunchServices.plist on the macOS machine. Please refer to Apple's documentation for details. The API uses the Windows Registry and LSCopyDefaultHandlerForURLScheme internally.
	IsDefaultProtocolClient func(Protocol string, Path string, Args *js.Object) (Obj bool) `js:"isDefaultProtocolClient"`
	// Adds tasks to the Tasks category of the JumpList on Windows. tasks is an array of Task objects. Note: If you'd like to customize the Jump List even more use app.setJumpList(categories) instead.
	SetUserTasks        func(Tasks *js.Object) (Obj bool)             `js:"setUserTasks"`
	GetJumpListSettings func() (Obj *AppModuleGetJumpListSettingsObj) `js:"getJumpListSettings"`
	// Sets or removes a custom Jump List for the application, and returns one of the following strings: If categories is null the previously set custom Jump List (if any) will be replaced by the standard Jump List for the app (managed by Windows). Note: If a JumpListCategory object has neither the type nor the name property set then its type is assumed to be tasks. If the name property is set but the type property is omitted then the type is assumed to be custom. Note: Users can remove items from custom categories, and Windows will not allow a removed item to be added back into a custom category until after the next successful call to app.setJumpList(categories). Any attempt to re-add a removed item to a custom category earlier than that will result in the entire custom category being omitted from the Jump List. The list of removed items can be obtained using app.getJumpListSettings(). Here's a very simple example of creating a custom Jump List:
	SetJumpList func(Categories *js.Object) `js:"setJumpList"`
	// This method makes your application a Single Instance Application - instead of allowing multiple instances of your app to run, this will ensure that only a single instance of your app is running, and other instances signal this instance and exit. callback will be called with callback(argv, workingDirectory) when a second instance has been executed. argv is an Array of the second instance's command line arguments, and workingDirectory is its current working directory. Usually applications respond to this by making their primary window focused and non-minimized. The callback is guaranteed to be executed after the ready event of app gets emitted. This method returns false if your process is the primary instance of the application and your app should continue loading. And returns true if your process has sent its parameters to another instance, and you should immediately quit. On macOS the system enforces single instance automatically when users try to open a second instance of your app in Finder, and the open-file and open-url events will be emitted for that. However when users start your app in command line the system's single instance mechanism will be bypassed and you have to use this method to ensure single instance. An example of activating the window of primary instance when a second instance starts:
	MakeSingleInstance func(Callback AppModuleMakeSingleInstanceCallback) `js:"makeSingleInstance"`
	// Releases all locks that were created by makeSingleInstance. This will allow multiple instances of the application to once again run side by side.
	ReleaseSingleInstance func() `js:"releaseSingleInstance"`
	// Creates an NSUserActivity and sets it as the current activity. The activity is eligible for Handoff to another device afterward.
	SetUserActivity        func(Type string, UserInfo *AppModuleSetUserActivityUserInfo, WebpageURL string) `js:"setUserActivity"`
	GetCurrentActivityType func() (Obj string)                                                              `js:"getCurrentActivityType"`
	// Changes the Application User Model ID to id.
	SetAppUserModelId func(Id string) `js:"setAppUserModelId"`
	// Imports the certificate in pkcs12 format into the platform certificate store. callback is called with the result of import operation, a value of 0 indicates success while any other value indicates failure according to chromium net_error_list.
	ImportCertificate func(Options *AppModuleImportCertificateOptions, Callback AppModuleImportCertificateCallback) `js:"importCertificate"`
	// Disables hardware acceleration for current app. This method can only be called before app is ready.
	DisableHardwareAcceleration func() `js:"disableHardwareAcceleration"`
	// Sets the counter badge for current app. Setting the count to 0 will hide the badge. On macOS it shows on the dock icon. On Linux it only works for Unity launcher, Note: Unity launcher requires the exsistence of a .desktop file to work, for more information please read Desktop Environment Integration.
	SetBadgeCount  func(Count int64) (Obj bool) `js:"setBadgeCount"`
	GetBadgeCount  func() (Obj int64)           `js:"getBadgeCount"`
	IsUnityRunning func() (Obj bool)            `js:"isUnityRunning"`
	// Note: This API has no effect on MAS builds.
	GetLoginItemSettings func() (Obj *AppModuleGetLoginItemSettingsObj) `js:"getLoginItemSettings"`
	// Set the app's login item settings. Note: This API has no effect on MAS builds.
	SetLoginItemSettings          func(Settings *AppModuleSetLoginItemSettingsSettings) `js:"setLoginItemSettings"`
	IsAccessibilitySupportEnabled func() (Obj bool)                                     `js:"isAccessibilitySupportEnabled"`
	// Set the about panel options. This will override the values defined in the app's .plist file. See the Apple docs for more details.
	SetAboutPanelOptions func(Options *AppModuleSetAboutPanelOptionsOptions) `js:"setAboutPanelOptions"`
}

func GetAppModule() *AppModule {
	o := Get("app")
	return &AppModule{
		Emitter: events.New(o),
	}
}

type AppModuleSetAboutPanelOptionsOptions struct {
	*js.Object
	// The app's name.
	ApplicationName string `js:"applicationName"`
	// The app's version.
	ApplicationVersion string `js:"applicationVersion"`
	// Copyright information.
	Copyright string `js:"copyright"`
	// Credit information.
	Credits string `js:"credits"`
	// The app's build version number.
	Version string `js:"version"`
}

type AppModuleAppModuleCommandLine struct {
	*js.Object
	// Append a switch (with optional value) to Chromium's command line. Note: This will not affect process.argv, and is mainly used by developers to control some low-level Chromium behaviors.
	AppendSwitch AppModuleCommandLineAppendSwitch `js:"appendSwitch"`
	// Append an argument to Chromium's command line. The argument will be quoted correctly. Note: This will not affect process.argv.
	AppendArgument AppModuleCommandLineAppendArgument `js:"appendArgument"`
}

type AppModuleCommandLineAppendSwitch func( // A command-line switch
	Switch string, // A value for the given switch
						Value string)
type AppModuleCommandLineAppendArgument func( // The argument to append to the command line
	Value string)
type AppModuleAppModuleDock struct {
	*js.Object
	// When critical is passed, the dock icon will bounce until either the application becomes active or the request is canceled. When informational is passed, the dock icon will bounce for one second. However, the request remains active until either the application becomes active or the request is canceled.
	Bounce AppModuleDockBounce `js:"bounce"`
	// Cancel the bounce of id.
	CancelBounce AppModuleDockCancelBounce `js:"cancelBounce"`
	// Bounces the Downloads stack if the filePath is inside the Downloads folder.
	DownloadFinished AppModuleDockDownloadFinished `js:"downloadFinished"`
	// Sets the string to be displayed in the dock’s badging area.
	SetBadge AppModuleDockSetBadge `js:"setBadge"`
	GetBadge AppModuleDockGetBadge `js:"getBadge"`
	// Hides the dock icon.
	Hide AppModuleDockHide `js:"hide"`
	// Shows the dock icon.
	Show      AppModuleDockShow      `js:"show"`
	IsVisible AppModuleDockIsVisible `js:"isVisible"`
	// Sets the application's dock menu.
	SetMenu AppModuleDockSetMenu `js:"setMenu"`
	// Sets the image associated with this dock icon.
	SetIcon AppModuleDockSetIcon `js:"setIcon"`
}

type AppModuleDockDownloadFinished func(FilePath string)
type AppModuleDockSetIcon func(Image *NativeImage)
type AppModuleDockCancelBounce func(Id int64)
type AppModuleDockSetBadge func(Text string)
type AppModuleDockGetBadge func()
type AppModuleDockHide func()
type AppModuleDockShow func()
type AppModuleDockIsVisible func()
type AppModuleDockSetMenu func(Menu *Menu)
type AppModuleDockBounce func( // Can be `critical` or `informational`. The default is `informational`
	Type AppModuleBounceType)
type AppModuleBounceType string

// consts
const (
	AppModuleBounceTypeCritical      AppModuleBounceType = "critical"
	AppModuleBounceTypeInformational AppModuleBounceType = "informational"
)

type AppModuleGetJumpListSettingsObj struct {
	*js.Object
	// The minimum number of items that will be shown in the Jump List (for a more detailed description of this value see the ).
	MinItems int64 `js:"minItems"`
	// Array of objects that correspond to items that the user has explicitly removed from custom categories in the Jump List. These items must not be re-added to the Jump List in the call to , Windows will not display any custom category that contains any of the removed items.
	RemovedItems *js.Object `js:"removedItems"`
}

type AppModuleSetUserActivityUserInfo struct {
	*js.Object
}

type AppModuleGetLoginItemSettingsObj struct {
	*js.Object
	// if the app is set to open at login.
	OpenAtLogin bool `js:"openAtLogin"`
	// if the app is set to open as hidden at login. This setting is only supported on macOS.
	OpenAsHidden bool `js:"openAsHidden"`
	// if the app was opened at login automatically. This setting is only supported on macOS.
	WasOpenedAtLogin bool `js:"wasOpenedAtLogin"`
	// if the app was opened as a hidden login item. This indicates that the app should not open any windows at startup. This setting is only supported on macOS.
	WasOpenedAsHidden bool `js:"wasOpenedAsHidden"`
	// if the app was opened as a login item that should restore the state from the previous session. This indicates that the app should restore the windows that were open the last time the app was closed. This setting is only supported on macOS.
	RestoreState bool `js:"restoreState"`
}

type AppModuleSetLoginItemSettingsSettings struct {
	*js.Object
	// to open the app at login, to remove the app as a login item. Defaults to .
	OpenAtLogin bool `js:"openAtLogin"`
	// to open the app as hidden. Defaults to . The user can edit this setting from the System Preferences so should be checked when the app is opened to know the current value. This setting is only supported on macOS.
	OpenAsHidden bool `js:"openAsHidden"`
}

type AppModuleRelaunchOptions struct {
	*js.Object
	// (optional)
	Args     *js.Object `js:"args"`
	ExecPath string     `js:"execPath"`
}

type AppModuleMakeSingleInstanceCallback func( // An array of the second instance's command line arguments
	Argv *js.Object, // The second instance's working directory
	WorkingDirectory string)
type AppModuleImportCertificateOptions struct {
	*js.Object
	// Path for the pkcs12 file.
	Certificate string `js:"certificate"`
	// Passphrase for the certificate.
	Password string `js:"password"`
}

type AppModuleImportCertificateCallback func( // Result of import.
	Result int64)
