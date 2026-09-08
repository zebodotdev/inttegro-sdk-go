package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

// Pay initiates payment for an existing order.
//
// Use this to charge an order after creation, or to retry a failed payment.
// Provide either PaymentMethodID (for saved methods) or PaymentMethodData
// (for one-time use).
//
// Parameters:
//   - OrderID: The order to charge (required)
//   - PaymentMethodID: Saved payment method ID (optional, mutually exclusive with PaymentMethodData)
//   - PaymentMethodData: Inline payment details (optional, mutually exclusive with PaymentMethodID)
//   - PaidOutOfBand: Mark as paid externally (optional)
//
// Returns the updated order with its payment and next-action state.
//
// Example (charging with saved payment method):
//
//	order, err := client.Orders.Pay(ctx, order.PayParams{
//	    OrderID:         "or_abc123",
//	    PaymentMethodID: "pm_def456",
//	})
//	if err != nil {
//	    return err
//	}
//	if order.Payment != nil && order.Payment.NextAction != nil {
//	    // Prompt for OTP and call ConfirmPayment()
//	}
//
// Example (charging with inline payment method):
//
//	order, err := client.Orders.Pay(ctx, order.PayParams{
//	    OrderID: "or_abc123",
//	    PaymentMethodData: &paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	})
//
// Learn more: https://studio.inttegro.com/orders#pay-for-an-order
func (s *Service) Pay(ctx context.Context, params PayParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/pay", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// ConfirmPayment confirms a payment using an OTP token.
//
// Call this after initiating payment when confirms_use is enabled.
// The customer receives an OTP via SMS—submit it here to complete payment.
//
// Parameters:
//   - OrderID: The order awaiting confirmation (required)
//   - Token: The OTP from the customer (required, typically 4-6 digits)
//
// Returns the updated order with payment status.
//
// Example:
//
//	// Customer enters OTP: "123456"
//	order, err := client.Orders.ConfirmPayment(ctx, order.ConfirmParams{
//	    OrderID: "or_abc123",
//	    Token:   "123456",
//	})
//	if err != nil {
//	    // Invalid or expired token
//	    return err
//	}
//	// Payment confirmed, check order.Payment.Status
//
// Learn more: https://studio.inttegro.com/confirm-payment
func (s *Service) ConfirmPayment(ctx context.Context, params ConfirmParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/confirm_payment", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// RequestConfirmation requests a new OTP for payment confirmation.
//
// Use this if the customer didn't receive the original OTP or if it expired.
// Triggers a new OTP to be sent to the customer's phone number.
//
// Parameters:
//   - orderID: The order needing a new confirmation token (required)
//
// Returns the updated order with the new confirmation request.
//
// Example:
//
//	// Customer says they didn't receive OTP
//	order, err := client.Orders.RequestConfirmation(ctx, "or_abc123")
//	if err != nil {
//	    return err
//	}
//	// New OTP sent, prompt customer again
func (s *Service) RequestConfirmation(ctx context.Context, orderID string) (*Resource, error) {
	return s.RequestConfirmationWithParams(ctx, RequestConfirmationParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("request_confirmation", orderID),
	})
}

func (s *Service) RequestConfirmationWithParams(ctx context.Context, params RequestConfirmationParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/request_confirmation", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// OrderPayParams initiates payment for an existing order.
//
// Use this to charge an order after creation, or to retry payment
// on an order with a failed payment attempt.
//
// Provide either PaymentMethodID (for saved payment methods) or
// PaymentMethodData (for one-time payment methods).
//
// Example (with saved payment method):
//
//	params := order.PayParams{
//	    OrderID:         createdOrder.ID,
//	    PaymentMethodID: "pm_abc123",
//	}
//	response, err := client.Orders.Pay(ctx, params)
type PayParams struct {
	// OrderID is the order to charge (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID references a saved payment method (optional).
	// Mutually exclusive with PaymentMethodData.
	// Use this for repeat customers with tokenized payment methods.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// PaymentMethodData provides inline payment details (optional).
	// Mutually exclusive with PaymentMethodID.
	// Use this for one-time payments without saving the method.
	PaymentMethodData *paymentmethod.Data `json:"payment_method_data,omitempty"`

	// PaidOutOfBand marks payment as completed outside Inttegro (optional).
	// Set to true if customer paid via cash, bank transfer, or other method.
	// The order is marked paid without actually charging the payment method.
	// Use carefully—this bypasses payment processing entirely.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderConfirmParams confirms a payment using a verification token.
//
// After initiating payment on an order with confirms_use enabled,
// the customer receives an OTP. Use this to submit their OTP and
// complete the payment.
//
// Example:
//
//	params := order.ConfirmParams{
//	    OrderID: createdOrder.ID,
//	    Token:   "123456", // OTP from customer
//	}
//	order, err := client.Orders.ConfirmPayment(ctx, params)
type ConfirmParams struct {
	// OrderID is the order awaiting confirmation (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// Token is the OTP or verification code (required).
	// Typically 4-6 digits sent to the customer's phone via SMS.
	Token string `json:"token"`
}

// OrderRequestConfirmationParams requests a new OTP for payment confirmation.
//
// Use this if the customer didn't receive the original OTP or if it expired.
// Triggers a new OTP to be sent to the customer's phone number.
type RequestConfirmationParams struct {
	// OrderID is the order needing a new confirmation token (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}
