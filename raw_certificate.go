package electron

import "github.com/gopherjs/gopherjs/js"

// Certificate a Structure
type Certificate struct {
	*js.Object
	// PEM encoded data
	Data string `js:"data"`
	// Issuer principal
	Issuer *js.Object `js:"issuer"`
	// Issuer's Common Name
	IssuerName string `js:"issuerName"`
	// Issuer certificate (if not self-signed)
	IssuerCert *js.Object `js:"issuerCert"`
	// Subject principal
	Subject *js.Object `js:"subject"`
	// Subject's Common Name
	SubjectName string `js:"subjectName"`
	// Hex value represented string
	SerialNumber string `js:"serialNumber"`
	// Start date of the certificate being valid in seconds
	ValidStart float64 `js:"validStart"`
	// End date of the certificate being valid in seconds
	ValidExpiry float64 `js:"validExpiry"`
	// Fingerprint of the certificate
	Fingerprint string `js:"fingerprint"`
}
