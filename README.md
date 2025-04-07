# OTP Library

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/otp?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/otp.svg)](https://pkg.go.dev/github.com/go-universal/otp)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/otp/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/otp)](https://goreportcard.com/report/github.com/go-universal/otp)
![Contributors](https://img.shields.io/github/contributors/go-universal/otp)
![Issues](https://img.shields.io/github/issues/go-universal/otp)

This library provides an implementation of One-Time Password (OTP) generation and validation, including support for Google Authenticator. It allows you to generate OTP codes, retrieve the setup URI, and generate QR code images for easy integration with authenticator apps.

## Installation

To use this library, add it to your Go project by running:

```bash
go get github.com/go-universal/otp
```

## Features

- Generate OTP codes compatible with Google Authenticator.
- Validate OTP codes.
- Generate QR codes for easy setup in authenticator apps.
- Support for configurable OTP validity periods.

## Usage

### Creating an OTP Instance

To create an OTP instance, use the `NewGoogleOTP` function. This function generates a new OTP key based on the provided username, issuer, and optional identifiers.

```go
import (
    "fmt"

    "github.com/go-universal/otp"
)

func main() {
    issuer := "MyApp"
    username := "user@example.com"
    ids := []string{"device1", "device2"}

    otpInstance, err := otp.NewGoogleOTP(username, issuer, otp.Period30, ids...)
    if err != nil {
        fmt.Println("Error creating OTP:", err)
        return
    }

    fmt.Println("OTP instance created successfully!")
}
```

### Validating an OTP Code

To validate an OTP code, use the `Validate` method. This checks if the provided code is valid for the current OTP instance.

```go
code := "123456" // Replace with the code to validate
isValid := otpInstance.Validate(code)
if isValid {
    fmt.Println("The OTP code is valid!")
} else {
    fmt.Println("The OTP code is invalid.")
}
```

### Retrieving the OTP URI

The `URI` method retrieves the `otpauth://` URI, which can be used to configure an authenticator app manually.

```go
uri := otpInstance.URI()
fmt.Println("OTP URI:", uri)
```

### Generating a QR Code

The `QR` method generates a QR code as a PNG byte slice. This QR code can be displayed to the user for easy setup in an authenticator app.

```go
qrCode, err := otpInstance.QR()
if err != nil {
    fmt.Println("Error generating QR code:", err)
    return
}

// Save the QR code to a file or display it
fmt.Println("QR code generated successfully!")
```

## Constants

### `const Period30 Period = 30`

Defines a 30-second OTP validity period.

### `const Period60 Period = 60`

Defines a 60-second OTP validity period.

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
