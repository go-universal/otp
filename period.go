package otp

// Period represents the time duration (in seconds) for OTP validity.
type Period uint

const (
	// Period30 defines a 30-second OTP validity period.
	Period30 Period = 30

	// Period60 defines a 60-second OTP validity period.
	Period60 Period = 60
)
