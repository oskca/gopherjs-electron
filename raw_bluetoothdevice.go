package electron

import "github.com/gopherjs/gopherjs/js"

// BluetoothDevice a Structure
type BluetoothDevice struct {
	*js.Object
	DeviceName string `js:"deviceName"`
	DeviceId   string `js:"deviceId"`
}
