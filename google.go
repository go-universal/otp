package otp

import (
	"crypto/md5"
	"encoding/base32"
	"strings"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

// google is a concrete implementation of the OTP interface.
type google struct {
	key *otp.Key
}

// NewGoogleOTP create new Google Authenticator OTP driver.
func NewGoogleOTP(username, issuer string, period Period, ids ...string) (OTP, error) {
	// Generate a unique secret using MD5 hash of the concatenated identifiers, username, and issuer.
	hash := md5.Sum([]byte(strings.Join(append(ids, username, issuer), ":")))
	secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash[:])

	// Generate the OTP key with the provided options.
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: username,
		Period:      uint(period),
		Algorithm:   otp.AlgorithmSHA1,
		SecretSize:  16,
		Secret:      []byte(secret),
	})
	if err != nil {
		return nil, err
	}

	return &google{key: key}, nil
}

func (g *google) Validate(code string) bool {
	return totp.Validate(code, g.key.Secret())
}

func (g *google) URI() string {
	return g.key.URL()
}

func (g *google) QR() ([]byte, error) {
	return qrcode.Encode(g.URI(), qrcode.Medium, 512)
}
