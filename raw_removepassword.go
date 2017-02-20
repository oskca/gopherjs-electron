package electron

import "github.com/gopherjs/gopherjs/js"

// RemovePassword a Structure
type RemovePassword struct {
	*js.Object
	// .
	Type string `js:"type"`
	// When provided, the authentication info related to the origin will only be removed otherwise the entire cache will be cleared.
	Origin string `js:"origin"`
	// Scheme of the authentication. Can be , , , . Must be provided if removing by .
	Scheme RemovePasswordRemovePasswordScheme `js:"scheme"`
	// Realm of the authentication. Must be provided if removing by .
	Realm string `js:"realm"`
	// Credentials of the authentication. Must be provided if removing by .
	Username string `js:"username"`
	// Credentials of the authentication. Must be provided if removing by .
	Password string `js:"password"`
}

type RemovePasswordRemovePasswordScheme string

// consts
const (
	RemovePasswordRemovePasswordSchemeBasic     RemovePasswordRemovePasswordScheme = "basic"
	RemovePasswordRemovePasswordSchemeDigest    RemovePasswordRemovePasswordScheme = "digest"
	RemovePasswordRemovePasswordSchemeNtlm      RemovePasswordRemovePasswordScheme = "ntlm"
	RemovePasswordRemovePasswordSchemeNegotiate RemovePasswordRemovePasswordScheme = "negotiate"
)
