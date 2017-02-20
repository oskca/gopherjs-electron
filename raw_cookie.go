package electron

import "github.com/gopherjs/gopherjs/js"

// Cookie a Structure
type Cookie struct {
	*js.Object
	// The name of the cookie.
	Name string `js:"name"`
	// The value of the cookie.
	Value string `js:"value"`
	// The domain of the cookie.
	Domain string `js:"domain"`
	// Whether the cookie is a host-only cookie.
	HostOnly bool `js:"hostOnly"`
	// The path of the cookie.
	Path string `js:"path"`
	// Whether the cookie is marked as secure.
	Secure bool `js:"secure"`
	// Whether the cookie is marked as HTTP only.
	HttpOnly bool `js:"httpOnly"`
	// Whether the cookie is a session cookie or a persistent cookie with an expiration date.
	Session bool `js:"session"`
	// The expiration date of the cookie as the number of seconds since the UNIX epoch. Not provided for session cookies.
	ExpirationDate float64 `js:"expirationDate"`
}
