package inttegro

import "context"

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
type OtpService struct {
	client *Client
}

// Initiate starts an OTP session and sends a verification code.
//
// Returns OTP details including expiration and delivery info.
func (s *OtpService) Initiate(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/otp/initiate", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Verify checks whether the provided OTP code is correct.
//
// Returns verification result including success status.
func (s *OtpService) Verify(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/otp/verify", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Lookup retrieves an existing OTP transaction by ID.
func (s *OtpService) Lookup(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/otp/lookup", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Cancel invalidates an active OTP transaction.
func (s *OtpService) Cancel(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/otp/cancel", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Initialize is a backwards-compatible alias for Initiate.
func (s *OtpService) Initialize(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return s.Initiate(ctx, payload)
}
