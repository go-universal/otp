package otp_test

import (
	"testing"

	"github.com/go-universal/otp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createOTP() (otp.OTP, error) {
	issuer := "Test Issuer"
	username := "test_user"
	ids := []string{"key1", "key2"}

	return otp.NewGoogleOTP(username, issuer, otp.Period30, ids...)
}

func TestGoogle(t *testing.T) {
	otp, err := createOTP()
	require.NoError(t, err, "Failed to create OTP")

	code := "081966"
	valid := otp.Validate(code)
	assert.True(t, valid, "Expected code to be valid")
}

func TestQR(t *testing.T) {
	otp, err := createOTP()
	require.NoError(t, err, "Failed to create OTP")

	_, err = otp.QR()
	assert.NoError(t, err, "Failed to generate QR code")
}
