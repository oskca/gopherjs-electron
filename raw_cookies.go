package electron

import "github.com/oskca/gopherjs-nodejs/events"

import "github.com/gopherjs/gopherjs/js"

// events
const (
	// Emitted when a cookie is changed because it was added, edited, removed, or expired.
	EvtCookiesChanged = "changed"
)

// Cookies version@1.4.15
//
// Query and modify a session's cookies.
type Cookies struct {
	*events.Emitter
	// Sends a request to get all cookies matching details, callback will be called with callback(error, cookies) on complete. cookies is an Array of cookie objects.
	Get func(Filter *CookiesGetFilter, Callback CookiesGetCallback) `js:"get"`
	// Sets a cookie with details, callback will be called with callback(error) on complete.
	Set func(Details *CookiesSetDetails, Callback CookiesSetCallback) `js:"set"`
	// Removes the cookies matching url and name, callback will called with callback() on complete.
	Remove func(URL string, Name string, Callback CookiesRemoveCallback) `js:"remove"`
}

func WrapCookies(o *js.Object) *Cookies {
	return &Cookies{
		Emitter: events.New(o),
	}
}

type CookiesSetDetails struct {
	*js.Object
	// The url to associate the cookie with.
	URL string `js:"url"`
	// The name of the cookie. Empty by default if omitted.
	Name string `js:"name"`
	// The value of the cookie. Empty by default if omitted.
	Value string `js:"value"`
	// The domain of the cookie. Empty by default if omitted.
	Domain string `js:"domain"`
	// The path of the cookie. Empty by default if omitted.
	Path string `js:"path"`
	// Whether the cookie should be marked as Secure. Defaults to false.
	Secure bool `js:"secure"`
	// Whether the cookie should be marked as HTTP only. Defaults to false.
	HttpOnly bool `js:"httpOnly"`
	// The expiration date of the cookie as the number of seconds since the UNIX epoch. If omitted then the cookie becomes a session cookie and will not be retained between sessions.
	ExpirationDate float64 `js:"expirationDate"`
}

type CookiesSetCallback func(Error *js.Object)
type CookiesRemoveCallback func()
type CookiesGetFilter struct {
	*js.Object
	// Retrieves cookies which are associated with . Empty implies retrieving cookies of all urls.
	URL string `js:"url"`
	// Filters cookies by name.
	Name string `js:"name"`
	// Retrieves cookies whose domains match or are subdomains of
	Domain string `js:"domain"`
	// Retrieves cookies whose path matches .
	Path string `js:"path"`
	// Filters cookies by their Secure property.
	Secure bool `js:"secure"`
	// Filters out session or persistent cookies.
	Session bool `js:"session"`
}

type CookiesGetCallback func(Error *js.Object, Cookies *js.Object)
