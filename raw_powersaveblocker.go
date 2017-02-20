package electron

import "github.com/gopherjs/gopherjs/js"

// PowerSaveBlockerModule version@1.4.15
//
// Block the system from entering low-power (sleep) mode.
type PowerSaveBlockerModule struct {
	*js.Object
	// Starts preventing the system from entering lower-power mode. Returns an integer identifying the power save blocker. Note: prevent-display-sleep has higher precedence over prevent-app-suspension. Only the highest precedence type takes effect. In other words, prevent-display-sleep always takes precedence over prevent-app-suspension. For example, an API calling A requests for prevent-app-suspension, and another calling B requests for prevent-display-sleep. prevent-display-sleep will be used until B stops its request. After that, prevent-app-suspension is used.
	Start func(Type PowerSaveBlockerModuleStartType) (Obj int64) `js:"start"`
	// Stops the specified power save blocker.
	Stop      func(Id int64)            `js:"stop"`
	IsStarted func(Id int64) (Obj bool) `js:"isStarted"`
}

func GetPowerSaveBlockerModule() *PowerSaveBlockerModule {
	o := Get("powerSaveBlocker")
	return &PowerSaveBlockerModule{
		Object: o,
	}
}

type PowerSaveBlockerModuleStartType string

// consts
const (
	PowerSaveBlockerModuleStartTypePreventAppSuspension PowerSaveBlockerModuleStartType = "prevent-app-suspension"
	PowerSaveBlockerModuleStartTypePreventDisplaySleep  PowerSaveBlockerModuleStartType = "prevent-display-sleep"
)
