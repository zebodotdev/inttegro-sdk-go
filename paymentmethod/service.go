package paymentmethod

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
)

// PaymentMethodsService manages payment method tokenization, verification, and deletion.
//
// Payment methods are saved customer payment instruments that can be charged
// repeatedly. Use this service to:
//
//   - Tokenize payment methods for repeat customers
//   - Verify ownership via OTP
//   - Look up payment method details
//   - Delete payment methods
//   - View payment method acceptance settings
//
// Example:
//
//	// Tokenize and verify a mobile money wallet
//	pm, err := client.PaymentMethods.Tokenize(ctx, paymentmethod.TokenizeParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: inttegro.Bool(true),
//	})
//
// Learn more: https://studio.inttegro.com/save-payment-methods
type Service struct {
	client transport.Client
}

// Tokenize saves a payment method for future use. Optionally verifies immediately.
//
// Tokenized payment methods can be charged repeatedly without re-entering details.
// The customer owns the payment method—only they can delete it.
//
// Learn more: https://studio.inttegro.com/tokenize-payment-methods
func (s *Service) Tokenize(ctx context.Context, params TokenizeParams) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/tokenize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Verify sends an OTP to confirm the customer owns the payment method.
//
// Required before first use (unless confirms_use is false). Returns verification
// status and OTP delivery details.
func (s *Service) Verify(ctx context.Context, paymentMethodID string) (*VerificationSession, error) {
	return s.VerifyWithParams(ctx, VerifyParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("verify", paymentMethodID),
	})
}

func (s *Service) VerifyWithParams(ctx context.Context, params VerifyParams) (*VerificationSession, error) {
	var resp struct {
		Verification VerificationSession `json:"verification"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Verification, nil
}

// ConfirmVerification submits the OTP to complete verification.
//
// Call this after Verify() once the customer provides their OTP.
// Returns the verified payment method.
func (s *Service) ConfirmVerification(ctx context.Context, params ConfirmVerificationParams) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/confirm_verification", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Lookup retrieves payment method details by ID.
//
// Returns masked payment details, verification status, and enabled state.
func (s *Service) Lookup(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/lookup", LookupParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Page retrieves a paginated list of payment methods.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update modifies mutable payment method fields.
func (s *Service) Update(ctx context.Context, payload any) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Activate marks a payment method active.
func (s *Service) Activate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/activate", paymentMethodID)
}

// Disactivate marks a payment method inactive.
func (s *Service) Disactivate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/disactivate", paymentMethodID)
}

// Deactivate is an alias for Disactivate.
func (s *Service) Deactivate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.Disactivate(ctx, paymentMethodID)
}

// Archive archives a payment method.
func (s *Service) Archive(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/archive", paymentMethodID)
}

// Unarchive unarchives a payment method.
func (s *Service) Unarchive(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/unarchive", paymentMethodID)
}

func (s *Service) paymentMethodAction(ctx context.Context, path, paymentMethodID string) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", path, ActionParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Delete permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete their
// own payment methods—merchants cannot prevent this or restore them.
func (s *Service) Delete(ctx context.Context, paymentMethodID string) (*Deletion, error) {
	return s.DeleteWithParams(ctx, DeleteParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("delete", paymentMethodID),
	})
}

func (s *Service) DeleteWithParams(ctx context.Context, params DeleteParams) (*Deletion, error) {
	var resp Deletion
	if err := s.client.Do(ctx, "POST", "/payment_methods/delete", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func stablePaymentMethodRequestMeta(action, paymentMethodID string) *request.Meta {
	return &request.Meta{IdempotencyKey: "payment_methods_" + action + "_" + paymentMethodID}
}

// Settings retrieves payment method acceptance configuration.
//
// Shows which payment types are enabled and whether OTP confirmation is required.
func (s *Service) Settings(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/settings", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
