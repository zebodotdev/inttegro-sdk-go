// Package otp provides otp resources and operations.
package otp

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

type AlphabetType string

const (
	AlphabetTypeNumeric AlphabetType = "numeric"
)

const (
	AlphabetTypeAlpha AlphabetType = "alpha"
)

const (
	AlphabetTypeAlphanumeric AlphabetType = "alphanumeric"
)

type Status string

const (
	StatusCanceled Status = "canceled"
)

const (
	StatusExpired Status = "expired"
)

const (
	StatusPending Status = "pending"
)

const (
	StatusPendingDelivery Status = "pending_delivery"
)

const (
	StatusPendingVerification Status = "pending_verification"
)

const (
	StatusVerified Status = "verified"
)

type TransmissionStatus string

const (
	TransmissionStatusDelivered TransmissionStatus = "delivered"
)

const (
	TransmissionStatusFailed TransmissionStatus = "failed"
)

const (
	TransmissionStatusSubmitted TransmissionStatus = "submitted"
)

type VerificationVerdict string

const (
	VerificationVerdictFail VerificationVerdict = "fail"
)

const (
	VerificationVerdictPass VerificationVerdict = "pass"
)

// OtpService manages one-time password generation and verification.
//
// This service provides low-level OTP functionality for custom authentication
// flows. Most integrations should use payment confirmation or payment method
// verification instead—those handle OTP automatically.
//
// Use this service when:
//   - Building custom 2FA flows
//   - Implementing passwordless authentication
//   - Verifying phone numbers outside payment context
type Service struct {
	client transport.Client
}

// Initiate starts an OTP session and sends a verification code.
//
// Returns OTP details including expiration and delivery info.
func (s *Service) Initiate(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/otp/initiate", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Verify checks whether the provided OTP code is correct.
//
// Returns verification result including success status.
func (s *Service) Verify(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/otp/verify", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Lookup retrieves an existing OTP transaction by ID.
func (s *Service) Lookup(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/otp/lookup", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Cancel invalidates an active OTP transaction.
func (s *Service) Cancel(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/otp/cancel", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
