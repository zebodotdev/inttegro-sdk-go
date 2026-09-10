package otp

import "time"

// Transaction describes an OTP delivery and verification session.
type Transaction struct {
	CancelReason string        `json:"cancel_reason,omitempty"`
	CanceledAt   *time.Time    `json:"canceled_at,omitempty"`
	ExpiresAt    time.Time     `json:"expires_at"`
	FullMessage  string        `json:"full_message"`
	ID           string        `json:"id"`
	InitiatedAt  time.Time     `json:"initiated_at"`
	Status       Status        `json:"status"`
	Transmission *Transmission `json:"transmission,omitempty"`
}

// Transmission describes delivery of an OTP message.
type Transmission struct {
	Recipient string             `json:"recipient"`
	SenderID  string             `json:"sender_id"`
	SentAt    *time.Time         `json:"sent_at,omitempty"`
	SentVia   string             `json:"sent_via,omitempty"`
	Status    TransmissionStatus `json:"status,omitempty"`
}

// Verification records the updated transaction and the submitted attempt.
type Verification struct {
	Transaction         Transaction         `json:"transaction"`
	VerificationAttempt VerificationAttempt `json:"verification_attempt"`
}

// VerificationAttempt describes one submitted OTP token.
type VerificationAttempt struct {
	AttemptedAt    time.Time          `json:"attempted_at"`
	ID             string             `json:"id"`
	PresentedToken string             `json:"presented_token"`
	Recipient      string             `json:"recipient"`
	Result         VerificationResult `json:"result"`
}

// VerificationResult says whether a submitted OTP token matched.
type VerificationResult struct {
	Detail  string              `json:"detail,omitempty"`
	Verdict VerificationVerdict `json:"verdict"`
}
