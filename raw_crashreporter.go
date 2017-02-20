package electron

import "github.com/gopherjs/gopherjs/js"

// CrashReporterModule version@1.4.15
//
// Submit crash reports to a remote server.
type CrashReporterModule struct {
	*js.Object
	// You are required to call this method before using any other crashReporter APIs and in each process (main/renderer) from which you want to collect crash reports. You can pass different options to crashReporter.start when calling from different processes. Note: On Windows and Linux, Electron uses breakpad for crash collection and reporting. Crashes can be collected from the main and renderer process, but not from the child processes created via the child_process module. Note: On macOS, Electron uses a new crashpad client for crash collection and reporting. Crashes can be collected from the main, renderer and any of the child processes created via the child_process module. If you want to enable crash reporting, initializing crashpad from the main process using crashReporter.start is required regardless of which process you want to collect crashes from. Once initialized this way, the crashpad handler collects crashes from all processes. You still have to call crashReporter.start from the renderer process, otherwise crashes from renderer processes will get reported without companyName, productName or any of the extra information.
	Start func(Options *CrashReporterModuleStartOptions) `js:"start"`
	// Returns the date and ID of the last crash report. If no crash reports have been sent or the crash reporter has not been started, null is returned.
	GetLastCrashReport func() (Obj *js.Object) `js:"getLastCrashReport"`
	// Returns all uploaded crash reports. Each report contains the date and uploaded ID.
	GetUploadedReports func() (Obj *js.Object) `js:"getUploadedReports"`
	// Note: This API can only be called from the main process.
	GetUploadToServer func() (Obj bool) `js:"getUploadToServer"`
	// This would normally be controlled by user preferences. This has no effect if called before start is called. Note: This API can only be called from the main process.
	SetUploadToServer func(UploadToServer bool) `js:"setUploadToServer"`
}

func GetCrashReporterModule() *CrashReporterModule {
	o := Get("crashReporter")
	return &CrashReporterModule{
		Object: o,
	}
}

type CrashReporterModuleStartOptions struct {
	*js.Object
	CompanyName string `js:"companyName"`
	// URL that crash reports will be sent to as POST.
	SubmitURL string `js:"submitURL"`
	// Defaults to .
	ProductName string `js:"productName"`
	// Whether crash reports should be sent to the server Default is .
	UploadToServer bool `js:"uploadToServer"`
	// Default is .
	IgnoreSystemCrashHandler bool `js:"ignoreSystemCrashHandler"`
	// An object you can define that will be sent along with the report. Only string properties are sent correctly, Nested objects are not supported.
	Extra *CrashReporterModuleOptionsExtra `js:"extra"`
}

type CrashReporterModuleOptionsExtra struct {
	*js.Object
}
