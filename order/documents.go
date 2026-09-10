package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

// SendInvoice sends the hosted invoice link for an existing order.
//
// Inttegro delivers the invoice link to every contact method available on the
// order customer.
func (s *Service) SendInvoice(ctx context.Context, params SendInvoiceParams) (*DocumentDeliveryResult, error) {
	var resp DocumentDeliveryResult
	if err := s.client.Do(ctx, "POST", "/orders/send_invoice", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendReceipt sends the hosted receipt link for a paid order.
//
// Receipt delivery is only valid after the order has been paid.
func (s *Service) SendReceipt(ctx context.Context, params SendReceiptParams) (*DocumentDeliveryResult, error) {
	var resp DocumentDeliveryResult
	if err := s.client.Do(ctx, "POST", "/orders/send_receipt", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderSendInvoiceParams sends an invoice link for an order.
type SendInvoiceParams struct {
	// OrderID is the order whose invoice should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderSendReceiptParams sends a receipt link for a paid order.
type SendReceiptParams struct {
	// OrderID is the paid order whose receipt should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}
