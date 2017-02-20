package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

const (
	EvtSystemPreferencesAccentColorChanged         = "accent-color-changed"
	EvtSystemPreferencesColorChanged               = "color-changed"
	EvtSystemPreferencesInvertedColorSchemeChanged = "inverted-color-scheme-changed"
)

// SystemPreferencesModule version@1.4.15
//
// Get system preferences.
type SystemPreferencesModule struct {
	*events.Emitter
	IsDarkMode                             func() (Obj bool) `js:"isDarkMode"`
	IsSwipeTrackingFromScrollEventsEnabled func() (Obj bool) `js:"isSwipeTrackingFromScrollEventsEnabled"`
	// Posts event as native notifications of macOS. The userInfo is an Object that contains the user information dictionary sent along with the notification.
	PostNotification func(Event string, UserInfo *SystemPreferencesModulePostNotificationUserInfo) `js:"postNotification"`
	// Posts event as native notifications of macOS. The userInfo is an Object that contains the user information dictionary sent along with the notification.
	PostLocalNotification func(Event string, UserInfo *SystemPreferencesModulePostLocalNotificationUserInfo) `js:"postLocalNotification"`
	// Subscribes to native notifications of macOS, callback will be called with callback(event, userInfo) when the corresponding event happens. The userInfo is an Object that contains the user information dictionary sent along with the notification. The id of the subscriber is returned, which can be used to unsubscribe the event. Under the hood this API subscribes to NSDistributedNotificationCenter, example values of event are:
	SubscribeNotification func(Event string, Callback SystemPreferencesModuleSubscribeNotificationCallback) `js:"subscribeNotification"`
	// Removes the subscriber with id.
	UnsubscribeNotification func(Id int64) `js:"unsubscribeNotification"`
	// Same as subscribeNotification, but uses NSNotificationCenter for local defaults. This is necessary for events such as NSUserDefaultsDidChangeNotification
	SubscribeLocalNotification func(Event string, Callback SystemPreferencesModuleSubscribeLocalNotificationCallback) `js:"subscribeLocalNotification"`
	// Same as unsubscribeNotification, but removes the subscriber from NSNotificationCenter.
	UnsubscribeLocalNotification func(Id int64) `js:"unsubscribeLocalNotification"`
	// Get the value of key in system preferences. This API uses NSUserDefaults on macOS. Some popular key and types are:
	GetUserDefault func(Key string, Type SystemPreferencesModuleGetUserDefaultType) `js:"getUserDefault"`
	// Set the value of key in system preferences. Note that type should match actual type of value. An exception is thrown if they don't. This API uses NSUserDefaults on macOS. Some popular key and types are:
	SetUserDefault func(Key string, Type string, Value string) `js:"setUserDefault"`
	// This method returns true if DWM composition (Aero Glass) is enabled, and false otherwise. An example of using it to determine if you should create a transparent window or not (transparent windows won't work correctly when DWM composition is disabled):
	IsAeroGlassEnabled    func()                                                        `js:"isAeroGlassEnabled"`
	GetAccentColor        func() (Obj string)                                           `js:"getAccentColor"`
	GetColor              func(Color SystemPreferencesModuleGetColorColor) (Obj string) `js:"getColor"`
	IsInvertedColorScheme func() (Obj bool)                                             `js:"isInvertedColorScheme"`
}

func GetSystemPreferencesModule() *SystemPreferencesModule {
	o := Get("systemPreferences")
	return &SystemPreferencesModule{
		Emitter: events.New(o),
	}
}

type SystemPreferencesModuleSubscribeLocalNotificationCallback func(Event string, UserInfo *SystemPreferencesModuleCallbackUserInfo)
type SystemPreferencesModuleCallbackUserInfo struct {
	*js.Object
}

type SystemPreferencesModulePostNotificationUserInfo struct {
	*js.Object
}

type SystemPreferencesModulePostLocalNotificationUserInfo struct {
	*js.Object
}

type SystemPreferencesModuleSubscribeNotificationCallback func(Event string, UserInfo *SystemPreferencesModuleCallbackUserInfo2)
type SystemPreferencesModuleCallbackUserInfo2 struct {
	*js.Object
}

type SystemPreferencesModuleGetColorColor string

// consts
const (
	SystemPreferencesModuleGetColorColor3dDarkShadow            SystemPreferencesModuleGetColorColor = "3d-dark-shadow"
	SystemPreferencesModuleGetColorColor3dFace                  SystemPreferencesModuleGetColorColor = "3d-face"
	SystemPreferencesModuleGetColorColor3dHighlight             SystemPreferencesModuleGetColorColor = "3d-highlight"
	SystemPreferencesModuleGetColorColor3dLight                 SystemPreferencesModuleGetColorColor = "3d-light"
	SystemPreferencesModuleGetColorColor3dShadow                SystemPreferencesModuleGetColorColor = "3d-shadow"
	SystemPreferencesModuleGetColorColorActiveBorder            SystemPreferencesModuleGetColorColor = "active-border"
	SystemPreferencesModuleGetColorColorActiveCaption           SystemPreferencesModuleGetColorColor = "active-caption"
	SystemPreferencesModuleGetColorColorActiveCaptionGradient   SystemPreferencesModuleGetColorColor = "active-caption-gradient"
	SystemPreferencesModuleGetColorColorAppWorkspace            SystemPreferencesModuleGetColorColor = "app-workspace"
	SystemPreferencesModuleGetColorColorButtonText              SystemPreferencesModuleGetColorColor = "button-text"
	SystemPreferencesModuleGetColorColorCaptionText             SystemPreferencesModuleGetColorColor = "caption-text"
	SystemPreferencesModuleGetColorColorDesktop                 SystemPreferencesModuleGetColorColor = "desktop"
	SystemPreferencesModuleGetColorColorDisabledText            SystemPreferencesModuleGetColorColor = "disabled-text"
	SystemPreferencesModuleGetColorColorHighlight               SystemPreferencesModuleGetColorColor = "highlight"
	SystemPreferencesModuleGetColorColorHighlightText           SystemPreferencesModuleGetColorColor = "highlight-text"
	SystemPreferencesModuleGetColorColorHotlight                SystemPreferencesModuleGetColorColor = "hotlight"
	SystemPreferencesModuleGetColorColorInactiveBorder          SystemPreferencesModuleGetColorColor = "inactive-border"
	SystemPreferencesModuleGetColorColorInactiveCaption         SystemPreferencesModuleGetColorColor = "inactive-caption"
	SystemPreferencesModuleGetColorColorInactiveCaptionGradient SystemPreferencesModuleGetColorColor = "inactive-caption-gradient"
	SystemPreferencesModuleGetColorColorInactiveCaptionText     SystemPreferencesModuleGetColorColor = "inactive-caption-text"
	SystemPreferencesModuleGetColorColorInfoBackground          SystemPreferencesModuleGetColorColor = "info-background"
	SystemPreferencesModuleGetColorColorInfoText                SystemPreferencesModuleGetColorColor = "info-text"
	SystemPreferencesModuleGetColorColorMenu                    SystemPreferencesModuleGetColorColor = "menu"
	SystemPreferencesModuleGetColorColorMenuHighlight           SystemPreferencesModuleGetColorColor = "menu-highlight"
	SystemPreferencesModuleGetColorColorMenubar                 SystemPreferencesModuleGetColorColor = "menubar"
	SystemPreferencesModuleGetColorColorMenuText                SystemPreferencesModuleGetColorColor = "menu-text"
	SystemPreferencesModuleGetColorColorScrollbar               SystemPreferencesModuleGetColorColor = "scrollbar"
	SystemPreferencesModuleGetColorColorWindow                  SystemPreferencesModuleGetColorColor = "window"
	SystemPreferencesModuleGetColorColorWindowFrame             SystemPreferencesModuleGetColorColor = "window-frame"
	SystemPreferencesModuleGetColorColorWindowText              SystemPreferencesModuleGetColorColor = "window-text"
)

type SystemPreferencesModuleGetUserDefaultType string

// consts
const (
	SystemPreferencesModuleGetUserDefaultTypeString     SystemPreferencesModuleGetUserDefaultType = "string"
	SystemPreferencesModuleGetUserDefaultTypeBoolean    SystemPreferencesModuleGetUserDefaultType = "boolean"
	SystemPreferencesModuleGetUserDefaultTypeInteger    SystemPreferencesModuleGetUserDefaultType = "integer"
	SystemPreferencesModuleGetUserDefaultTypeFloat      SystemPreferencesModuleGetUserDefaultType = "float"
	SystemPreferencesModuleGetUserDefaultTypeDouble     SystemPreferencesModuleGetUserDefaultType = "double"
	SystemPreferencesModuleGetUserDefaultTypeURL        SystemPreferencesModuleGetUserDefaultType = "url"
	SystemPreferencesModuleGetUserDefaultTypeArray      SystemPreferencesModuleGetUserDefaultType = "array"
	SystemPreferencesModuleGetUserDefaultTypeDictionary SystemPreferencesModuleGetUserDefaultType = "dictionary"
)
