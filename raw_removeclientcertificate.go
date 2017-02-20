package electron

import "github.com/gopherjs/gopherjs/js"

// RemoveClientCertificate a Structure
type RemoveClientCertificate struct {
	*js.Object
	// .
	Type string `js:"type"`
	// Origin of the server whose associated client certificate must be removed from the cache.
	Origin string `js:"origin"`
}
