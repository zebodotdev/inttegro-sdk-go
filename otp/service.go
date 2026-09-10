package otp

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
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
func (s *Service) Initiate(ctx context.Context, params InitiateParams, opts ...request.Option) (*Transaction, error) {
	var resp struct {
		Transaction Transaction `json:"transaction"`
	}
	if err := s.client.DoJSON(ctx, "/otp/initiate", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Transaction, nil
}

// Verify checks whether the provided OTP code is correct.
//
// Returns verification result including success status.
func (s *Service) Verify(ctx context.Context, params VerifyParams, opts ...request.Option) (*Verification, error) {
	var resp Verification
	if err := s.client.DoJSON(ctx, "/otp/verify", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Lookup retrieves an existing OTP transaction by ID.
func (s *Service) Lookup(ctx context.Context, params LookupParams) (*Transaction, error) {
	var resp struct {
		Transaction Transaction `json:"transaction"`
	}
	if err := s.client.Do(ctx, "POST", "/otp/lookup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Transaction, nil
}

// Cancel invalidates an active OTP transaction.
func (s *Service) Cancel(ctx context.Context, params CancelParams, opts ...request.Option) (*Transaction, error) {
	var resp struct {
		Transaction Transaction `json:"transaction"`
	}
	if err := s.client.DoJSON(ctx, "/otp/cancel", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Transaction, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
