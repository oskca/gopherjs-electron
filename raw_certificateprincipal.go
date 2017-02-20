package electron

import "github.com/gopherjs/gopherjs/js"

// CertificatePrincipal a Structure
type CertificatePrincipal struct {
	*js.Object
	// Common Name
	CommonName string `js:"commonName"`
	// Organization names
	Organizations *js.Object `js:"organizations"`
	// Organization Unit names
	OrganizationUnits *js.Object `js:"organizationUnits"`
	// Locality
	Locality string `js:"locality"`
	// State or province
	State string `js:"state"`
	// Country or region
	Country string `js:"country"`
}
