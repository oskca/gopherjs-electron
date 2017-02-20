package electron

import "github.com/oskca/gopherjs-nodejs/events"

const (
	// Emitted when the system is suspending.
	EvtPowerMonitorSuspend = "suspend"
	// Emitted when system is resuming.
	EvtPowerMonitorResume = "resume"
	// Emitted when the system changes to AC power.
	EvtPowerMonitorOnAc = "on-ac"
	// Emitted when system changes to battery power.
	EvtPowerMonitorOnBattery = "on-battery"
)

// PowerMonitorModule version@1.4.15
//
// Monitor power state changes.
type PowerMonitorModule struct {
	*events.Emitter
}

func GetPowerMonitorModule() *PowerMonitorModule {
	o := Get("powerMonitor")
	return &PowerMonitorModule{
		Emitter: events.New(o),
	}
}
