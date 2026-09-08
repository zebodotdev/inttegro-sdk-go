package otp

type AlphabetType string

const (
	AlphabetTypeNumeric      AlphabetType = "numeric"
	AlphabetTypeAlpha        AlphabetType = "alpha"
	AlphabetTypeAlphanumeric AlphabetType = "alphanumeric"
)

type Status string

const (
	StatusCanceled            Status = "canceled"
	StatusExpired             Status = "expired"
	StatusPending             Status = "pending"
	StatusPendingDelivery     Status = "pending_delivery"
	StatusPendingVerification Status = "pending_verification"
	StatusVerified            Status = "verified"
)

type TransmissionStatus string

const (
	TransmissionStatusDelivered TransmissionStatus = "delivered"
	TransmissionStatusFailed    TransmissionStatus = "failed"
	TransmissionStatusSubmitted TransmissionStatus = "submitted"
)

type VerificationVerdict string

const (
	VerificationVerdictFail VerificationVerdict = "fail"
	VerificationVerdictPass VerificationVerdict = "pass"
)
