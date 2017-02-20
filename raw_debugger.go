package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when debugging session is terminated. This happens either when webContents is closed or devtools is invoked for the attached webContents.
	EvtDebuggerDetach = "detach"
	// Emitted whenever debugging target issues instrumentation event.
	EvtDebuggerMessage = "message"
)

// Debugger version@1.4.15
//
// An alternate transport for Chrome's remote debugging protocol.
type Debugger struct {
	*events.Emitter
	// Attaches the debugger to the webContents.
	Attach     func(ProtocolVersion string) `js:"attach"`
	IsAttached func() (Obj bool)            `js:"isAttached"`
	// Detaches the debugger from the webContents.
	Detach func() `js:"detach"`
	// Send given command to the debugging target.
	SendCommand func(Method string, CommandParams *DebuggerSendCommandCommandParams, Callback DebuggerSendCommandCallback) `js:"sendCommand"`
}

func WrapDebugger(o *js.Object) *Debugger {
	return &Debugger{
		Emitter: events.New(o),
	}
}

type DebuggerSendCommandCommandParams struct {
	*js.Object
}

type DebuggerSendCommandCallback func( // Error message indicating the failure of the command.
	Error *DebuggerCallbackError, // Response defined by the 'returns' attribute of the command description in the remote debugging protocol.
	Result *js.Object)
type DebuggerCallbackError struct {
	*js.Object
}
