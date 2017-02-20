package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

const (
	// Emitted when Electron has loaded its internal initialization script and is beginning to load the web page or the main script. It can be used by the preload script to add removed Node global symbols back to the global scope when node integration is turned off:
	EvtProcessLoaded = "loaded"
)

// ProcessModule version@1.4.15
//
// Extensions to process object.
type ProcessModule struct {
	*events.Emitter
	// Setting this to true can disable the support for asar archives in Node's built-in modules.
	NoAsar string `js:"noAsar"`
	// Current process's type, can be "browser" (i.e. main process) or "renderer".
	Type string `js:"type"`
	// Electron's version string.
	Electron string `js:"electron"`
	// Chrome's version string.
	Chrome string `js:"chrome"`
	// Path to the resources directory.
	ResourcesPath string `js:"resourcesPath"`
	// For Mac App Store build, this property is true, for other builds it is undefined.
	Mas string `js:"mas"`
	// If the app is running as a Windows Store app (appx), this property is true, for otherwise it is undefined.
	WindowsStore string `js:"windowsStore"`
	// When app is started by being passed as parameter to the default app, this property is true in the main process, otherwise it is undefined.
	DefaultApp string `js:"defaultApp"`
	// Causes the main thread of the current process crash.
	Crash func() `js:"crash"`
	// Causes the main thread of the current process hang.
	Hang func() `js:"hang"`
	// Sets the file descriptor soft limit to maxDescriptors or the OS hard limit, whichever is lower for the current process.
	SetFdLimit func(MaxDescriptors int64) `js:"setFdLimit"`
	// Returns an object giving memory usage statistics about the current process. Note that all statistics are reported in Kilobytes.
	GetProcessMemoryInfo func() (Obj *ProcessModuleGetProcessMemoryInfoObj) `js:"getProcessMemoryInfo"`
	// Returns an object giving memory usage statistics about the entire system. Note that all statistics are reported in Kilobytes.
	GetSystemMemoryInfo func() (Obj *ProcessModuleGetSystemMemoryInfoObj) `js:"getSystemMemoryInfo"`
}

func GetProcessModule() *ProcessModule {
	o := Get("process")
	return &ProcessModule{
		Emitter: events.New(o),
	}
}

type ProcessModuleGetProcessMemoryInfoObj struct {
	*js.Object
	// The amount of memory currently pinned to actual physical RAM.
	WorkingSetSize int64 `js:"workingSetSize"`
	// The maximum amount of memory that has ever been pinned to actual physical RAM.
	PeakWorkingSetSize int64 `js:"peakWorkingSetSize"`
	// The amount of memory not shared by other processes, such as JS heap or HTML content.
	PrivateBytes int64 `js:"privateBytes"`
	// The amount of memory shared between processes, typically memory consumed by the Electron code itself
	SharedBytes int64 `js:"sharedBytes"`
}

type ProcessModuleGetSystemMemoryInfoObj struct {
	*js.Object
	// The total amount of physical memory in Kilobytes available to the system.
	Total int64 `js:"total"`
	// The total amount of memory not being used by applications or disk cache.
	Free int64 `js:"free"`
	// The total amount of swap memory in Kilobytes available to the system.
	SwapTotal int64 `js:"swapTotal"`
	// The free amount of swap memory in Kilobytes available to the system.
	SwapFree int64 `js:"swapFree"`
}
