package inttegro

import "context"

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
//	pm, err := client.PaymentMethods.Tokenize(ctx, inttegro.TokenizePaymentMethodParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: inttegro.PaymentMethodData{
//	        Type: inttegro.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &inttegro.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: inttegro.Bool(true),
//	})
//
// Learn more: https://studio.inttegro.com/save-payment-methods
type PaymentMethodsService struct {
	client *Client
}

type PaymentMethodPageParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

type PaymentMethodPage struct {
	Number         int             `json:"number,omitempty"`
	Size           int             `json:"size,omitempty"`
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"`
}

type PaymentMethodActionParams struct {
	PaymentMethodID string `json:"payment_method_id"`
}

// Tokenize saves a payment method for future use. Optionally verifies immediately.
//
// Tokenized payment methods can be charged repeatedly without re-entering details.
// The customer owns the payment method—only they can delete it.
//
// Learn more: https://studio.inttegro.com/tokenize-payment-methods
func (s *PaymentMethodsService) Tokenize(ctx context.Context, params TokenizePaymentMethodParams) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/tokenize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Verify sends an OTP to confirm the customer owns the payment method.
//
// Required before first use (unless confirms_use is false). Returns verification
// status and OTP delivery details.
func (s *PaymentMethodsService) Verify(ctx context.Context, paymentMethodID string) (*PaymentMethodVerificationSession, error) {
	return s.VerifyWithParams(ctx, VerifyPaymentMethodParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("verify", paymentMethodID),
	})
}

func (s *PaymentMethodsService) VerifyWithParams(ctx context.Context, params VerifyPaymentMethodParams) (*PaymentMethodVerificationSession, error) {
	var resp struct {
		Verification PaymentMethodVerificationSession `json:"verification"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Verification, nil
}

// ConfirmVerification submits the OTP to complete verification.
//
// Call this after Verify() once the customer provides their OTP.
// Returns the verified payment method.
func (s *PaymentMethodsService) ConfirmVerification(ctx context.Context, params ConfirmPaymentMethodVerificationParams) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/confirm_verification", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Lookup retrieves payment method details by ID.
//
// Returns masked payment details, verification status, and enabled state.
func (s *PaymentMethodsService) Lookup(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/lookup", LookupPaymentMethodParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Page retrieves a paginated list of payment methods.
func (s *PaymentMethodsService) Page(ctx context.Context, params PaymentMethodPageParams) (*PaymentMethodPage, error) {
	var resp struct {
		Page PaymentMethodPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update modifies mutable payment method fields.
func (s *PaymentMethodsService) Update(ctx context.Context, payload any) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Activate marks a payment method active.
func (s *PaymentMethodsService) Activate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/activate", paymentMethodID)
}

// Disactivate marks a payment method inactive.
func (s *PaymentMethodsService) Disactivate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/disactivate", paymentMethodID)
}

// Deactivate is an alias for Disactivate.
func (s *PaymentMethodsService) Deactivate(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.Disactivate(ctx, paymentMethodID)
}

// Archive archives a payment method.
func (s *PaymentMethodsService) Archive(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/archive", paymentMethodID)
}

// Unarchive unarchives a payment method.
func (s *PaymentMethodsService) Unarchive(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/unarchive", paymentMethodID)
}

func (s *PaymentMethodsService) paymentMethodAction(ctx context.Context, path, paymentMethodID string) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", path, PaymentMethodActionParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Delete permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete their
// own payment methods—merchants cannot prevent this or restore them.
func (s *PaymentMethodsService) Delete(ctx context.Context, paymentMethodID string) (*PaymentMethodDeletion, error) {
	return s.DeleteWithParams(ctx, DeletePaymentMethodParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("delete", paymentMethodID),
	})
}

func (s *PaymentMethodsService) DeleteWithParams(ctx context.Context, params DeletePaymentMethodParams) (*PaymentMethodDeletion, error) {
	var resp PaymentMethodDeletion
	if err := s.client.do(ctx, "POST", "/payment_methods/delete", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func stablePaymentMethodRequestMeta(action, paymentMethodID string) *RequestMeta {
	return &RequestMeta{IdempotencyKey: "payment_methods_" + action + "_" + paymentMethodID}
}

// Settings retrieves payment method acceptance configuration.
//
// Shows which payment types are enabled and whether OTP confirmation is required.
func (s *PaymentMethodsService) Settings(ctx context.Context) (*PaymentMethodSettings, error) {
	var resp struct {
		Settings PaymentMethodSettings `json:"settings"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/settings", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}
