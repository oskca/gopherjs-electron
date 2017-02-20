package electron

import "github.com/gopherjs/gopherjs/js"

// ContentTracingModule version@1.4.15
//
// Collect tracing data from Chromium's content module for finding performance
//bottlenecks and slow operations.
type ContentTracingModule struct {
	*js.Object
	// Get a set of category groups. The category groups can change as new code paths are reached. Once all child processes have acknowledged the getCategories request the callback is invoked with an array of category groups.
	GetCategories func(Callback ContentTracingModuleGetCategoriesCallback) `js:"getCategories"`
	// Start recording on all processes. Recording begins immediately locally and asynchronously on child processes as soon as they receive the EnableRecording request. The callback will be called once all child processes have acknowledged the startRecording request. categoryFilter is a filter to control what category groups should be traced. A filter can have an optional - prefix to exclude category groups that contain a matching category. Having both included and excluded category patterns in the same list is not supported. Examples: traceOptions controls what kind of tracing is enabled, it is a comma-delimited list. Possible options are: The first 3 options are trace recoding modes and hence mutually exclusive. If more than one trace recording modes appear in the traceOptions string, the last one takes precedence. If none of the trace recording modes are specified, recording mode is record-until-full. The trace option will first be reset to the default option (record_mode set to record-until-full, enable_sampling and enable_systrace set to false) before options parsed from traceOptions are applied on it.
	StartRecording func(Options *ContentTracingModuleStartRecordingOptions, Callback ContentTracingModuleStartRecordingCallback) `js:"startRecording"`
	// Stop recording on all processes. Child processes typically cache trace data and only rarely flush and send trace data back to the main process. This helps to minimize the runtime overhead of tracing since sending trace data over IPC can be an expensive operation. So, to end tracing, we must asynchronously ask all child processes to flush any pending trace data. Once all child processes have acknowledged the stopRecording request, callback will be called with a file that contains the traced data. Trace data will be written into resultFilePath if it is not empty or into a temporary file. The actual file path will be passed to callback if it's not null.
	StopRecording func(ResultFilePath string, Callback ContentTracingModuleStopRecordingCallback) `js:"stopRecording"`
	// Start monitoring on all processes. Monitoring begins immediately locally and asynchronously on child processes as soon as they receive the startMonitoring request. Once all child processes have acknowledged the startMonitoring request the callback will be called.
	StartMonitoring func(Options *ContentTracingModuleStartMonitoringOptions, Callback ContentTracingModuleStartMonitoringCallback) `js:"startMonitoring"`
	// Stop monitoring on all processes. Once all child processes have acknowledged the stopMonitoring request the callback is called.
	StopMonitoring func(Callback ContentTracingModuleStopMonitoringCallback) `js:"stopMonitoring"`
	// Get the current monitoring traced data. Child processes typically cache trace data and only rarely flush and send trace data back to the main process. This is because it may be an expensive operation to send the trace data over IPC and we would like to avoid unneeded runtime overhead from tracing. So, to end tracing, we must asynchronously ask all child processes to flush any pending trace data. Once all child processes have acknowledged the captureMonitoringSnapshot request the callback will be called with a file that contains the traced data.
	CaptureMonitoringSnapshot func(ResultFilePath string, Callback ContentTracingModuleCaptureMonitoringSnapshotCallback) `js:"captureMonitoringSnapshot"`
	// Get the maximum usage across processes of trace buffer as a percentage of the full state. When the TraceBufferUsage value is determined the callback is called.
	GetTraceBufferUsage func(Callback ContentTracingModuleGetTraceBufferUsageCallback) `js:"getTraceBufferUsage"`
	// callback will be called every time the given event occurs on any process.
	SetWatchEvent func(CategoryName string, EventName string, Callback ContentTracingModuleSetWatchEventCallback) `js:"setWatchEvent"`
	// Cancel the watch event. This may lead to a race condition with the watch event callback if tracing is enabled.
	CancelWatchEvent func() `js:"cancelWatchEvent"`
}

func GetContentTracingModule() *ContentTracingModule {
	o := Get("contentTracing")
	return &ContentTracingModule{
		Object: o,
	}
}

type ContentTracingModuleStartRecordingCallback func()
type ContentTracingModuleStartMonitoringCallback func()
type ContentTracingModuleStopMonitoringCallback func()
type ContentTracingModuleSetWatchEventCallback func()
type ContentTracingModuleGetCategoriesCallback func(Categories *js.Object)
type ContentTracingModuleStartRecordingOptions struct {
	*js.Object
	CategoryFilter string `js:"categoryFilter"`
	TraceOptions   string `js:"traceOptions"`
}

type ContentTracingModuleStopRecordingCallback func(ResultFilePath string)
type ContentTracingModuleStartMonitoringOptions struct {
	*js.Object
	CategoryFilter string `js:"categoryFilter"`
	TraceOptions   string `js:"traceOptions"`
}

type ContentTracingModuleCaptureMonitoringSnapshotCallback func(ResultFilePath string)
type ContentTracingModuleGetTraceBufferUsageCallback func(Value float64, Percentage float64)
