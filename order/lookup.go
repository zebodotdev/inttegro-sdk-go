package order

import (
	"context"
)

// Lookup retrieves an order by ID.
//
// Use this to fetch current order state, payment status, and any required
// customer actions (like OTP confirmation).
//
// Parameters:
//   - orderID: The order's unique identifier (starts with "or_")
//
// Returns the complete order object including:
//   - Current status and timestamps
//   - Payment details and status
//   - Customer and billing information
//   - Line items and totals
//   - Invoice links (if finalized)
//
// Example:
//
//	order, err := client.Orders.Lookup(ctx, "or_abc123def456")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Order status: %s\n", order.Status)
//	if order.Payment != nil {
//	    fmt.Printf("Payment status: %s\n", order.Payment.Status)
//	}
//
// Common use cases:
//   - Polling for payment status after initiating charge
//   - Checking if OTP confirmation is still required
//   - Displaying order details to customer
//   - Syncing order state with your system
//
// Learn more: https://studio.inttegro.com/retrieve-order
func (s *Service) Lookup(ctx context.Context, orderID string) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/lookup", LookupParams{OrderID: orderID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// OrderLookupParams specifies which order to retrieve.
type LookupParams struct {
	// OrderID is the unique order identifier (required).
	// Starts with "or_". Example: "or_abc123def456"
	OrderID string `json:"order_id"`
}
