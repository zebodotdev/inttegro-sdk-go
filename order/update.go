package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

// Update modifies mutable fields on an existing order.
func (s *Service) Update(ctx context.Context, params UpdateParams, opts ...request.Option) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.DoJSON(ctx, "/orders/update", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// UpdateParams contains the mutable fields accepted by /orders/update.
// LineItems replaces the order's complete line item set when supplied.
type UpdateParams struct {
	OrderID                   string              `json:"order_id"`
	ClearPaymentMethod        *bool               `json:"clear_payment_method,omitempty"`
	CustomData                map[string]string   `json:"custom_data,omitempty"`
	InvoiceSettings           *InvoiceSettings    `json:"invoice_settings,omitempty"`
	Finalize                  *bool               `json:"finalize,omitempty"`
	LineItems                 []LineItemParams    `json:"line_items,omitempty"`
	Number                    string              `json:"number,omitempty"`
	ReceiptNumber             string              `json:"receipt_number,omitempty"`
	PaymentMethodData         *paymentmethod.Data `json:"payment_method_data,omitempty"`
	PaymentMethodID           string              `json:"payment_method_id,omitempty"`
	StatementDescriptor       string              `json:"statement_descriptor,omitempty"`
	StatementDescriptorPrefix string              `json:"statement_descriptor_prefix,omitempty"`
}
