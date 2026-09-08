package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

// OrdersService provides access to order creation, payment, and lifecycle management.
//
// Orders represent purchase transactions with line items, customer information,
// and payment details. Use this service to:
//
//   - Create draft orders or charge immediately
//   - Initiate and confirm payments
//   - Finalize orders for hosted checkout
//   - Complete, cancel, or refund orders
//   - List recent orders
//
// Example:
//
//	// Create and charge an order
//	order, err := client.Orders.Create(ctx, order.CreateParams{
//	    CustomerData: &customer.Data{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    LineItems: []order.LineItemParams{
//	        {
//	            Type: order.LineItemTypeProduct,
//	            Product: &order.ProductLineItemParams{
//	                Type:     "digital",
//	                Name:     "Premium Plan",
//	                Quantity: 1,
//	                Price: price.InlineParams{AmountParams: money.AmountParams{
//	                    Currency: money.USD,
//	                    Value:    999,
//	                }},
//	            },
//	        },
//	    },
//	    BillingDetails: order.BillingDetails{...},
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &request.Meta{IdempotencyKey: "order_20231215_jane_001"},
//	})
type Service struct {
	client transport.Client
}

// Exec wraps do for shorter tests.
func (s *Service) Exec(ctx context.Context, method, path string, body any, out any) error {
	return s.client.Do(ctx, method, path, body, out)
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
