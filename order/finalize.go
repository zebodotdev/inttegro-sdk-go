package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

// Finalize seals an order and generates hosted checkout page and invoice.
//
// Finalizing an order:
//   - Locks the order (no more modifications)
//   - Calculates final totals
//   - Generates invoice documents (PDF and web page)
//   - Creates hosted checkout URL
//
// After finalization, the order cannot be edited. Required for hosted checkout flow.
//
// Parameters:
//   - orderID: The order to finalize (required)
//
// Returns the finalized order with invoice links.
//
// Example:
//
//	order, err := client.Orders.Finalize(ctx, "or_abc123")
//	if err != nil {
//	    return err
//	}
//	// Redirect customer to order.Invoice.Format.Web.URL
//	checkoutURL := order.Invoice.Format.Web.URL
func (s *Service) Finalize(ctx context.Context, orderID string) (*Order, error) {
	return s.FinalizeWithParams(ctx, FinalizeParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("finalize", orderID),
	})
}

func (s *Service) FinalizeWithParams(ctx context.Context, params FinalizeParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/finalize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// OrderFinalizeParams seals an order and generates checkout assets.
//
// Finalizing an order:
// - Locks the order (no more modifications)
// - Calculates final totals
// - Generates hosted checkout page
// - Creates invoice documents
//
// After finalization, the order cannot be edited. Use this when you're
// ready for the customer to pay.
type FinalizeParams struct {
	// OrderID is the order to finalize (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}
