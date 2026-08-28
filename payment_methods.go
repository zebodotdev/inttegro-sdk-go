package commerce

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
//	pm, err := client.PaymentMethods.Tokenize(ctx, commerce.TokenizePaymentMethodParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: commerce.PaymentMethodData{
//	        Type: commerce.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &commerce.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: commerce.Bool(true),
//	})
//
// Learn more: https://studio.zebo.dev/save-payment-methods
type PaymentMethodsService struct {
	client *Client
}

// Tokenize saves a payment method for future use. Optionally verifies immediately.
//
// Tokenized payment methods can be charged repeatedly without re-entering details.
// The customer owns the payment method—only they can delete it.
//
// Learn more: https://studio.zebo.dev/tokenize-payment-methods
func (s *PaymentMethodsService) Tokenize(ctx context.Context, params TokenizePaymentMethodParams) (*PaymentMethodObject, error) {
	var resp struct {
		PaymentMethod PaymentMethodObject `json:"payment_method"`
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
func (s *PaymentMethodsService) Verify(ctx context.Context, paymentMethodID string) (*VerificationStatusResponse, error) {
	return s.VerifyWithParams(ctx, VerifyPaymentMethodParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("verify", paymentMethodID),
	})
}

func (s *PaymentMethodsService) VerifyWithParams(ctx context.Context, params VerifyPaymentMethodParams) (*VerificationStatusResponse, error) {
	var resp VerificationStatusResponse
	if err := s.client.do(ctx, "POST", "/payment_methods/verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConfirmVerification submits the OTP to complete verification.
//
// Call this after Verify() once the customer provides their OTP.
// Returns the verified payment method.
func (s *PaymentMethodsService) ConfirmVerification(ctx context.Context, params ConfirmPaymentMethodVerificationParams) (*PaymentMethodObject, error) {
	var resp struct {
		PaymentMethod PaymentMethodObject `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/confirm_verification", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Lookup retrieves payment method details by ID.
//
// Returns masked payment details, verification status, and enabled state.
func (s *PaymentMethodsService) Lookup(ctx context.Context, paymentMethodID string) (*PaymentMethodObject, error) {
	var resp struct {
		PaymentMethod PaymentMethodObject `json:"payment_method"`
	}
	if err := s.client.do(ctx, "POST", "/payment_methods/lookup", LookupPaymentMethodParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Delete permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete their
// own payment methods—merchants cannot prevent this or restore them.
func (s *PaymentMethodsService) Delete(ctx context.Context, paymentMethodID string) (*DeletePaymentMethodResponse, error) {
	return s.DeleteWithParams(ctx, DeletePaymentMethodParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("delete", paymentMethodID),
	})
}

func (s *PaymentMethodsService) DeleteWithParams(ctx context.Context, params DeletePaymentMethodParams) (*DeletePaymentMethodResponse, error) {
	var resp DeletePaymentMethodResponse
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
