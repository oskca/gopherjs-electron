package electron

import "github.com/gopherjs/gopherjs/js"
import "github.com/oskca/gopherjs-nodejs/eventemitter"

var (
	app = &App{
		Object:       Get("app"),
		EventEmitter: eventemitter.New(Get("app")),
	}
)

// Event: ‘will-finish-launching’
// Emitted when the application has finished basic startup. On Windows and Linux, the will-finish-launching event is the same as the ready event; on macOS, this event represents the applicationWillFinishLaunching notification of NSApplication. You would usually set up listeners for the open-file and open-url events here, and start the crash reporter and auto updater.
const WillFinishLaunching = "will-finish-launching"

// In most cases, you should just do everything in the ready event handler.
// Event: ‘ready’
// Returns:
// launchInfo Object macOS
// Emitted when Electron has finished initializing. On macOS, launchInfo holds the userInfo of the NSUserNotification that was used to open the application, if it was launched from Notification Center. You can call app.isReady() to check if this event has already fired.
const Ready = "ready"

// Event: ‘window-all-closed’
// Emitted when all windows have been closed.
// If you do not subscribe to this event and all windows are closed, the default behavior is to quit the app; however, if you subscribe, you control whether the app quits or not. If the user pressed Cmd + Q, or the developer called app.quit(), Electron will first try to close all the windows and then emit the will-quit event, and in this case the window-all-closed event would not be emitted.
const WindowAllClosed = "window-all-closed"

// Event: ‘before-quit’
// Returns:
// event Event
// Emitted before the application starts closing its windows. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application.
//
// Note: If application quit was initiated by autoUpdater.quitAndInstall() then before-quit is emitted after emitting close event on all windows and closing them.
const BeforeQuit = "before-quit"

// Event: ‘will-quit’
// Returns:
// event Event
// Emitted when all windows have been closed and the application will quit. Calling event.preventDefault() will prevent the default behaviour, which is terminating the application.
//
// See the description of the window-all-closed event for the differences between the will-quit and window-all-closed events.
const WillQuit = "will-quit"

// Event: ‘quit’
// Returns:

// event Event
// exitCode Integer
// Emitted when the application is quitting.
const Quit = "quit"

// Event: ‘open-file’ macOS
// Returns:
//
// event Event
// path String
// Emitted when the user wants to open a file with the application. The open-file event is usually emitted when the application is already open and the OS wants to reuse the application to open the file. open-file is also emitted when a file is dropped onto the dock and the application is not yet running. Make sure to listen for the open-file event very early in your application startup to handle this case (even before the ready event is emitted).
//
// You should call event.preventDefault() if you want to handle this event.
//
// On Windows, you have to parse process.argv (in the main process) to get the filepath.
const OpenFile = "open-file"

// Event: ‘open-url’ macOS
// Returns:
//
// event Event
// url String
// Emitted when the user wants to open a URL with the application. Your application’s Info.plist file must define the url scheme within the CFBundleURLTypes key, and set NSPrincipalClass to AtomApplication.
//
// You should call event.preventDefault() if you want to handle this event.
const OpenUrl = "open-url"

// Event: ‘activate’ macOS
// Returns:
//
// event Event
// hasVisibleWindows Boolean
// Emitted when the application is activated, which usually happens when the user clicks on the application’s dock icon.
//
// Event: ‘continue-activity’ macOS
// Returns:
//
// event Event
// type String - A string identifying the activity. Maps to NSUserActivity.activityType.
// userInfo Object - Contains app-specific state stored by the activity on another device.
// Emitted during Handoff when an activity from a different device wants to be resumed. You should call event.preventDefault() if you want to handle this event.
//
// A user activity can be continued only in an app that has the same developer Team ID as the activity’s source app and that supports the activity’s type. Supported activity types are specified in the app’s Info.plist under the NSUserActivityTypes key.
const Activate = "activate"

// Event: ‘browser-window-blur’
// Returns:
//
// event Event
// window BrowserWindow
// Emitted when a browserWindow gets blurred.
const BrowserWindowBlur = "browser-window-blur"

// Event: ‘browser-window-focus’
// Returns:
//
// event Event
// window BrowserWindow
// Emitted when a browserWindow gets focused.
const BrowserWindowFocus = "browser-window-focus"

// Event: ‘browser-window-created’
// Returns:
//
// event Event
// window BrowserWindow
// Emitted when a new browserWindow is created.
const BrowserWindowCreated = "browser-window-created"

// Event: ‘web-contents-created’
// Returns:
//
// event Event
// webContents WebContents
// Emitted when a new webContents is created.
const WebContentsCreated = "web-contents-created"

// Event: ‘certificate-error’
// Returns:
//
// event Event
// webContents WebContents
// url String
// error String - The error code
// certificate Certificate
// callback Function
// isTrusted Boolean - Whether to consider the certificate as trusted
// Emitted when failed to verify the certificate for url, to trust the certificate you should prevent the default behavior with event.preventDefault() and call callback(true).
const CertificateError = "certificate-error"

// const {app} = require('electron')
//
// app.on('certificate-error', (event, webContents, url, error, certificate, callback) => {
//   if (url === 'https://github.com') {
//     // Verification logic.
//     event.preventDefault()
//     callback(true)
//   } else {
//     callback(false)
//   }
// })

// Event: ‘select-client-certificate’
// Returns:
//
// event Event
// webContents WebContents
// url URL
// certificateList Certificate[]
// callback Function
// certificate Certificate (optional)
// Emitted when a client certificate is requested.
//
// The url corresponds to the navigation entry requesting the client certificate and callback can be called with an entry filtered from the list. Using event.preventDefault() prevents the application from using the first certificate from the store.
//
// const {app} = require('electron')
//
// app.on('select-client-certificate', (event, webContents, url, list, callback) => {
//   event.preventDefault()
//   callback(list[0])
// })
const SelectClientCertificate = "select-client-certificate"

// Event: ‘login’
// Returns:
//
// event Event
// webContents WebContents
// request Object
// method String
// url URL
// referrer URL
// authInfo Object
// isProxy Boolean
// scheme String
// host String
// port Integer
// realm String
// callback Function
// username String
// password String
// Emitted when webContents wants to do basic auth.
//
// The default behavior is to cancel all authentications, to override this you should prevent the default behavior with event.preventDefault() and call callback(username, password) with the credentials.
//
// const {app} = require('electron')
//
// app.on('login', (event, webContents, request, authInfo, callback) => {
//   event.preventDefault()
//   callback('username', 'secret')
// })
const Login = "login"

// Event: ‘gpu-process-crashed’
// Returns:
//
// event Event
// killed Boolean
// Emitted when the gpu process crashes or is killed.
const GpuProcessCrashed = "gpu-process-crashed"

// Event: ‘accessibility-support-changed’ macOS Windows
// Returns:
//
// event Event
// accessibilitySupportEnabled Boolean - true when Chrome’s accessibility support is enabled, false otherwise.
// Emitted when Chrome’s accessibility support changes. This event fires when assistive technologies, such as screen readers, are enabled or disabled. See https://www.chromium.org/developers/design-documents/accessibility for more details.
const AccessibilitySupportChanged = "accessibility-support-changed"

func GetApp() *App {
	return app
}

type App struct {
	*js.Object
	*eventemitter.EventEmitter
	// app.quit()
	// Try to close all windows. The before-quit event will be emitted first. If all windows are successfully closed, the will-quit event will be emitted and by default the application will terminate.
	//
	// This method guarantees that all beforeunload and unload event handlers are correctly executed. It is possible that a window cancels the quitting by returning false in the beforeunload event handler.
	Quit func() `js:"quit"`

	// app.exit([exitCode])
	// exitCode Integer (optional)
	// Exits immediately with exitCode. exitCode defaults to 0.
	//
	// All windows will be closed immediately without asking user and the before-quit and will-quit events will not be emitted.
	Exit func(exitCode int) `js:"exit"`

	// app.focus()
	// On Linux, focuses on the first visible window. On macOS, makes the application the active app. On Windows, focuses on the application’s first window.
	Focus func() `js:"focus"`

	// app.hide() macOS
	// Hides all application windows without minimizing them.
	Hide func() `js:"hide"`

	// app.show() macOS
	// Shows application windows after they were hidden. Does not automatically focus them.
	Show func() `js:"show"`

	// app.getAppPath()
	// Returns String - The current application directory.
	//
	// app.getPath(name)
	// name String
	// Returns String - A path to a special directory or file associated with name. On failure an Error is thrown.
	//
	// You can request the following paths by the name:
	//
	// home User’s home directory.
	// appData Per-user application data directory, which by default points to:
	// %APPDATA% on Windows
	// $XDG_CONFIG_HOME or ~/.config on Linux
	// ~/Library/Application Support on macOS
	// userData The directory for storing your app’s configuration files, which by default it is the appData directory appended with your app’s name.
	// temp Temporary directory.
	// exe The current executable file.
	// module The libchromiumcontent library.
	// desktop The current user’s Desktop directory.
	// documents Directory for a user’s “My Documents”.
	// downloads Directory for a user’s downloads.
	// music Directory for a user’s music.
	// pictures Directory for a user’s pictures.
	// videos Directory for a user’s videos.
	// pepperFlashSystemPlugin Full path to the system version of the Pepper Flash plugin.
	GetPath func(name string) string `js:"getPath"`

	// app.getVersion()
	// Returns String - The version of the loaded application. If no version is found in the application’s package.json file, the version of the current bundle or executable is returned.
	GetVersion func() string `js:"getVersion"`

	// app.getName()
	// Returns String - The current application’s name, which is the name in the application’s package.json file.
	// Usually the name field of package.json is a short lowercased name, according to the npm modules spec. You should usually also specify a productName field, which is your application’s full capitalized name, and which will be preferred over name by Electron.
	GetName func() string `js:"getName"`

	// app.setName(name)
	// name String
	// Overrides the current application’s name.
	SetName func(name string) `js:"setName"`

	// app.relaunch([options])
	// options Object (optional)
	// args String[] - (optional)
	// execPath String (optional)
	// Relaunches the app when current instance exits.

	// By default the new instance will use the same working directory and command line arguments with current instance. When args is specified, the args will be passed as command line arguments instead. When execPath is specified, the execPath will be executed for relaunch instead of current app.

	// Note that this method does not quit the app when executed, you have to call app.quit or app.exit after calling app.relaunch to make the app restart.

	// When app.relaunch is called for multiple times, multiple instances will be started after current instance exited.

	// An example of restarting current instance immediately and adding a new command line argument to the new instance:

	// const {app} = require('electron')

	// app.relaunch({args: process.argv.slice(1).concat(['--relaunch'])})
	// app.exit(0)
	// app.isReady()
	// Returns Boolean - true if Electron has finished initializing, false otherwise.

	// app.setPath(name, path)
	// name String
	// path String
	// Overrides the path to a special directory or file associated with name. If the path specifies a directory that does not exist, the directory will be created by this method. On failure an Error is thrown.

	// You can only override paths of a name defined in app.getPath.

	// By default, web pages’ cookies and caches will be stored under the userData directory. If you want to change this location, you have to override the userData path before the ready event of the app module is emitted.

	// app.getLocale()
	// Returns String - The current application locale. Possible return values are documented here.

	// Note: When distributing your packaged app, you have to also ship the locales folder.

	// Note: On Windows you have to call it after the ready events gets emitted.

	// app.addRecentDocument(path) macOS Windows
	// path String
	// Adds path to the recent documents list.

	// This list is managed by the OS. On Windows you can visit the list from the task bar, and on macOS you can visit it from dock menu.

	// app.clearRecentDocuments() macOS Windows
	// Clears the recent documents list.

	// app.setAsDefaultProtocolClient(protocol[, path, args]) macOS Windows
	// protocol String - The name of your protocol, without ://. If you want your app to handle electron:// links, call this method with electron as the parameter.
	// path String (optional) Windows - Defaults to process.execPath
	// args String[] (optional) Windows - Defaults to an empty array
	// Returns Boolean - Whether the call succeeded.

	// This method sets the current executable as the default handler for a protocol (aka URI scheme). It allows you to integrate your app deeper into the operating system. Once registered, all links with your-protocol:// will be opened with the current executable. The whole link, including protocol, will be passed to your application as a parameter.

	// On Windows you can provide optional parameters path, the path to your executable, and args, an array of arguments to be passed to your executable when it launches.

	// Note: On macOS, you can only register protocols that have been added to your app’s info.plist, which can not be modified at runtime. You can however change the file with a simple text editor or script during build time. Please refer to Apple’s documentation for details.

	// The API uses the Windows Registry and LSSetDefaultHandlerForURLScheme internally.

	// app.removeAsDefaultProtocolClient(protocol[, path, args]) macOS Windows
	// protocol String - The name of your protocol, without ://.
	// path String (optional) Windows - Defaults to process.execPath
	// args String[] (optional) Windows - Defaults to an empty array
	// Returns Boolean - Whether the call succeeded.

	// This method checks if the current executable as the default handler for a protocol (aka URI scheme). If so, it will remove the app as the default handler.

	// app.isDefaultProtocolClient(protocol[, path, args]) macOS Windows
	// protocol String - The name of your protocol, without ://.
	// path String (optional) Windows - Defaults to process.execPath
	// args String[] (optional) Windows - Defaults to an empty array
	// Returns Boolean

	// This method checks if the current executable is the default handler for a protocol (aka URI scheme). If so, it will return true. Otherwise, it will return false.

	// Note: On macOS, you can use this method to check if the app has been registered as the default protocol handler for a protocol. You can also verify this by checking ~/Library/Preferences/com.apple.LaunchServices.plist on the macOS machine. Please refer to Apple’s documentation for details.

	// The API uses the Windows Registry and LSCopyDefaultHandlerForURLScheme internally.

	// app.setUserTasks(tasks) Windows
	// tasks Task[] - Array of Task objects
	// Adds tasks to the Tasks category of the JumpList on Windows.

	// tasks is an array of Task objects.

	// Returns Boolean - Whether the call succeeded.

	// Note: If you’d like to customize the Jump List even more use app.setJumpList(categories) instead.

	// app.getJumpListSettings() Windows
	// Returns Object:

	// minItems Integer - The minimum number of items that will be shown in the Jump List (for a more detailed description of this value see the MSDN docs).
	// removedItems JumpListItem[] - Array of JumpListItem objects that correspond to items that the user has explicitly removed from custom categories in the Jump List. These items must not be re-added to the Jump List in the next call to app.setJumpList(), Windows will not display any custom category that contains any of the removed items.
	// app.setJumpList(categories) Windows
	// categories JumpListCategory[] or null - Array of JumpListCategory objects.
	// Sets or removes a custom Jump List for the application, and returns one of the following strings:

	// ok - Nothing went wrong.
	// error - One or more errors occurred, enable runtime logging to figure out the likely cause.
	// invalidSeparatorError - An attempt was made to add a separator to a custom category in the Jump List. Separators are only allowed in the standard Tasks category.
	// fileTypeRegistrationError - An attempt was made to add a file link to the Jump List for a file type the app isn’t registered to handle.
	// customCategoryAccessDeniedError - Custom categories can’t be added to the Jump List due to user privacy or group policy settings.
	// If categories is null the previously set custom Jump List (if any) will be replaced by the standard Jump List for the app (managed by Windows).

	// Note: If a JumpListCategory object has neither the type nor the name property set then its type is assumed to be tasks. If the name property is set but the type property is omitted then the type is assumed to be custom.

	// Note: Users can remove items from custom categories, and Windows will not allow a removed item to be added back into a custom category until after the next successful call to app.setJumpList(categories). Any attempt to re-add a removed item to a custom category earlier than that will result in the entire custom category being omitted from the Jump List. The list of removed items can be obtained using app.getJumpListSettings().

	// Here’s a very simple example of creating a custom Jump List:

	// const {app} = require('electron')

	// app.setJumpList([
	//   {
	//     type: 'custom',
	//     name: 'Recent Projects',
	//     items: [
	//       { type: 'file', path: 'C:\\Projects\\project1.proj' },
	//       { type: 'file', path: 'C:\\Projects\\project2.proj' }
	//     ]
	//   },
	//   { // has a name so `type` is assumed to be "custom"
	//     name: 'Tools',
	//     items: [
	//       {
	//         type: 'task',
	//         title: 'Tool A',
	//         program: process.execPath,
	//         args: '--run-tool-a',
	//         icon: process.execPath,
	//         iconIndex: 0,
	//         description: 'Runs Tool A'
	//       },
	//       {
	//         type: 'task',
	//         title: 'Tool B',
	//         program: process.execPath,
	//         args: '--run-tool-b',
	//         icon: process.execPath,
	//         iconIndex: 0,
	//         description: 'Runs Tool B'
	//       }
	//     ]
	//   },
	//   { type: 'frequent' },
	//   { // has no name and no type so `type` is assumed to be "tasks"
	//     items: [
	//       {
	//         type: 'task',
	//         title: 'New Project',
	//         program: process.execPath,
	//         args: '--new-project',
	//         description: 'Create a new project.'
	//       },
	//       { type: 'separator' },
	//       {
	//         type: 'task',
	//         title: 'Recover Project',
	//         program: process.execPath,
	//         args: '--recover-project',
	//         description: 'Recover Project'
	//       }
	//     ]
	//   }
	// ])
	// app.makeSingleInstance(callback)
	// callback Function
	// argv String[] - An array of the second instance’s command line arguments
	// workingDirectory String - The second instance’s working directory
	// This method makes your application a Single Instance Application - instead of allowing multiple instances of your app to run, this will ensure that only a single instance of your app is running, and other instances signal this instance and exit.

	// callback will be called with callback(argv, workingDirectory) when a second instance has been executed. argv is an Array of the second instance’s command line arguments, and workingDirectory is its current working directory. Usually applications respond to this by making their primary window focused and non-minimized.

	// The callback is guaranteed to be executed after the ready event of app gets emitted.

	// This method returns false if your process is the primary instance of the application and your app should continue loading. And returns true if your process has sent its parameters to another instance, and you should immediately quit.

	// On macOS the system enforces single instance automatically when users try to open a second instance of your app in Finder, and the open-file and open-url events will be emitted for that. However when users start your app in command line the system’s single instance mechanism will be bypassed and you have to use this method to ensure single instance.

	// An example of activating the window of primary instance when a second instance starts:

	// const {app} = require('electron')
	// let myWindow = null

	// const shouldQuit = app.makeSingleInstance((commandLine, workingDirectory) => {
	//   // Someone tried to run a second instance, we should focus our window.
	//   if (myWindow) {
	//     if (myWindow.isMinimized()) myWindow.restore()
	//     myWindow.focus()
	//   }
	// })

	// if (shouldQuit) {
	//   app.quit()
	// }

	// // Create myWindow, load the rest of the app, etc...
	// app.on('ready', () => {
	// })
	// app.releaseSingleInstance()
	// Releases all locks that were created by makeSingleInstance. This will allow multiple instances of the application to once again run side by side.

	// app.setUserActivity(type, userInfo[, webpageURL]) macOS
	// type String - Uniquely identifies the activity. Maps to NSUserActivity.activityType.
	// userInfo Object - App-specific state to store for use by another device.
	// webpageURL String (optional) - The webpage to load in a browser if no suitable app is installed on the resuming device. The scheme must be http or https.
	// Creates an NSUserActivity and sets it as the current activity. The activity is eligible for Handoff to another device afterward.

	// app.getCurrentActivityType() macOS
	// Returns String - The type of the currently running activity.

	// app.setAppUserModelId(id) Windows
	// id String
	// Changes the Application User Model ID to id.

	// app.importCertificate(options, callback) LINUX
	// options Object
	// certificate String - Path for the pkcs12 file.
	// password String - Passphrase for the certificate.
	// callback Function
	// result Integer - Result of import.
	// Imports the certificate in pkcs12 format into the platform certificate store. callback is called with the result of import operation, a value of 0 indicates success while any other value indicates failure according to chromium net_error_list.

	// app.disableHardwareAcceleration()
	// Disables hardware acceleration for current app.

	// This method can only be called before app is ready.

	// app.setBadgeCount(count) Linux macOS
	// count Integer
	// Returns Boolean - Whether the call succeeded.

	// Sets the counter badge for current app. Setting the count to 0 will hide the badge.

	// On macOS it shows on the dock icon. On Linux it only works for Unity launcher,

	// Note: Unity launcher requires the exsistence of a .desktop file to work, for more information please read Desktop Environment Integration.

	// app.getBadgeCount() Linux macOCreateS
	// Returns Integer - The current value displayed in the counter badge.

	// app.isUnityRunning() Linux
	// Returns Boolean - Whether the current desktop environment is Unity launcher.

	// app.getLoginItemSettings() macOS Windows
	// Returns Object:

	// openAtLogin Boolean - true if the app is set to open at login.
	// openAsHidden Boolean - true if the app is set to open as hidden at login. This setting is only supported on macOS.
	// wasOpenedAtLogin Boolean - true if the app was opened at login automatically. This setting is only supported on macOS.
	// wasOpenedAsHidden Boolean - true if the app was opened as a hidden login item. This indicates that the app should not open any windows at startup. This setting is only supported on macOS.
	// restoreState Boolean - true if the app was opened as a login item that should restore the state from the previous session. This indicates that the app should restore the windows that were open the last time the app was closed. This setting is only supported on macOS.
	// Note: This API has no effect on MAS builds.

	// app.setLoginItemSettings(settings) macOS Windows
	// settings Object
	// openAtLogin Boolean (optional) - true to open the app at login, false to remove the app as a login item. Defaults to false.
	// openAsHidden Boolean (optional) - true to open the app as hidden. Defaults to false. The user can edit this setting from the System Preferences so app.getLoginItemStatus().wasOpenedAsHidden should be checked when the app is opened to know the current value. This setting is only supported on macOS.
	// Set the app’s login item settings.

	// Note: This API has no effect on MAS builds.

	// app.isAccessibilitySupportEnabled() macOS Windows
	// Returns Boolean - true if Chrome’s accessibility support is enabled, false otherwise. This API will return true if the use of assistive technologies, such as screen readers, has been detected. See https://www.chromium.org/developers/design-documents/accessibility for more details.

	// app.setAboutPanelOptions(options) macOS
	// options Object
	// applicationName String (optional) - The app’s name.
	// applicationVersion String (optional) - The app’s version.
	// copyright String (optional) - Copyright information.
	// credits String (optional) - Credit information.
	// version String (optional) - The app’s build version number.
	// Set the about panel options. This will override the values defined in the app’s .plist file. See the Apple docs for more details.

	// app.commandLine.appendSwitch(switch[, value])
	// switch String - A command-line switch
	// value String (optional) - A value for the given switch
	// Append a switch (with optional value) to Chromium’s command line.

	// Note: This will not affect process.argv, and is mainly used by developers to control some low-level Chromium behaviors.

	// app.commandLine.appendArgument(value)
	// value String - The argument to append to the command line
	// Append an argument to Chromium’s command line. The argument will be quoted correctly.

	// Note: This will not affect process.argv.

	// app.dock.bounce([type]) macOS
	// type String (optional) - Can be critical or informational. The default is informational
	// When critical is passed, the dock icon will bounce until either the application becomes active or the request is canceled.

	// When informational is passed, the dock icon will bounce for one second. However, the request remains active until either the application becomes active or the request is canceled.

	// Returns Integer an ID representing the request.

	// app.dock.cancelBounce(id) macOS
	// id Integer
	// Cancel the bounce of id.

	// app.dock.downloadFinished(filePath) macOS
	// filePath String
	// Bounces the Downloads stack if the filePath is inside the Downloads folder.

	// app.dock.setBadge(text) macOS
	// text String
	// Sets the string to be displayed in the dock’s badging area.

	// app.dock.getBadge() macOS
	// Returns String - The badge string of the dock.

	// app.dock.hide() macOS
	// Hides the dock icon.

	// app.dock.show() macOS
	// Shows the dock icon.

	// app.dock.isVisible() macOS
	// Returns Boolean - Whether the dock icon is visible. The app.dock.show() call is asynchronous so this method might not return true immediately after that call.

	// app.dock.setMenu(menu) macOS
	// menu Menu
	// Sets the application’s dock menu.

	// app.dock.setIcon(image) macOS
	// image (NativeImage | String)
	// Sets the image associated with this dock icon.
}