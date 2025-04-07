package otp

// OTP validating codes, retrieving the setup URI,
// and generating QR code images for One-Time Password (OTP).
type OTP interface {
	// Validate checks if the provided OTP code is valid.
	Validate(code string) bool

	// URI retrieves the otpauth:// URI used by authenticator apps.
	URI() string

	// QR generates a QR code as a PNG byte slice based on the otpauth URI.
	QR() ([]byte, error)
}
