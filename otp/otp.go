// Package otp provides one-time-password values and operations.
package otp

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service             = inttegro.OtpService
	AlphabetType        = inttegro.OTPAlphabetType
	Status              = inttegro.OTPStatus
	TransmissionStatus  = inttegro.OTPTransmissionStatus
	VerificationVerdict = inttegro.OTPVerificationVerdict
)

const (
	AlphabetTypeNumeric      = inttegro.OTPAlphabetTypeNumeric
	AlphabetTypeAlpha        = inttegro.OTPAlphabetTypeAlpha
	AlphabetTypeAlphanumeric = inttegro.OTPAlphabetTypeAlphanumeric

	StatusCanceled            = inttegro.OTPStatusCanceled
	StatusExpired             = inttegro.OTPStatusExpired
	StatusPending             = inttegro.OTPStatusPending
	StatusPendingDelivery     = inttegro.OTPStatusPendingDelivery
	StatusPendingVerification = inttegro.OTPStatusPendingVerification
	StatusVerified            = inttegro.OTPStatusVerified

	TransmissionStatusDelivered = inttegro.OTPTransmissionStatusDelivered
	TransmissionStatusFailed    = inttegro.OTPTransmissionStatusFailed
	TransmissionStatusSubmitted = inttegro.OTPTransmissionStatusSubmitted

	VerificationVerdictFail = inttegro.OTPVerificationVerdictFail
	VerificationVerdictPass = inttegro.OTPVerificationVerdictPass
)
