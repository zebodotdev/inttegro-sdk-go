package order

import (
	"context"
	"fmt"

	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
)

// Complete marks an order as completed and fulfilled.
//
// Completing an order:
//   - Marks order as fulfilled
//   - Creates balance transaction (makes funds available for payout after aging)
//   - Triggers payout eligibility countdown
//
// Only valid for paid orders. Use PaidOutOfBand if payment happened outside
// Inttegro (cash, bank transfer, etc).
//
// Parameters:
//   - params.OrderID: The order to complete (required)
//   - params.PaidOutOfBand: Mark as paid externally (optional)
//
// Returns the completed order.
//
// Example:
//
//	order, err := client.Orders.Complete(ctx, order.CompleteParams{
//	    OrderID: "or_abc123",
//	})
func (s *Service) Complete(ctx context.Context, params CompleteParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/complete", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Cancel cancels an order.
//
// Canceling an order:
//   - Prevents future payment attempts
//   - Marks order as permanently closed
//
// This operation does not move funds or create a refund. Use Refunds.Create
// when funds must be returned for paid line items.
//
// Parameters:
//   - orderID: The order to cancel (required)
//
// Returns the canceled order.
//
// Example:
//
//	order, err := client.Orders.Cancel(ctx, "or_abc123")
func (s *Service) Cancel(ctx context.Context, orderID string) (*Order, error) {
	return s.CancelWithParams(ctx, CancelParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("cancel", orderID),
	})
}

func (s *Service) CancelWithParams(ctx context.Context, params CancelParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/cancel", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

func stableOrderRequestMeta(action, orderID string) *request.Meta {
	return &request.Meta{IdempotencyKey: fmt.Sprintf("orders_%s_%s", action, orderID)}
}

// OrderCompleteParams marks an order as completed.
//
// Completing an order:
// - Marks order as fulfilled
// - Triggers payout eligibility (after aging period)
// - Creates balance transaction
//
// Only valid for paid orders. Use PaidOutOfBand if payment happened
// outside the Inttegro platform.
type CompleteParams struct {
	// OrderID is the order to complete (required).
	OrderID string `json:"order_id"`

	// PaidOutOfBand indicates payment happened externally (optional).
	// If true, marks order as paid without charging the payment method.
	// Use for cash, bank transfer, or other offline payment methods.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderCancelParams cancels an order.
//
// Canceling an order:
// - Prevents any future payment attempts
// - Marks order as permanently closed
//
// This operation does not move funds or create a refund.
type CancelParams struct {
	// OrderID is the order to cancel (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}
