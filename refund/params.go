package refund

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

// CreateRefundLineItem requests a refund allocation against one paid order
// line item. Reason and ReasonDetails are independent from the overall reason.
type CreateLineItem struct {
	OrderLineItemID string             `json:"order_line_item_id"`
	RefundAmount    money.AmountParams `json:"refund_amount"`
	Reason          *Reason            `json:"reason,omitempty"`
	ReasonDetails   string             `json:"reason_details,omitempty"`
}

// CreateRefundRequest starts a refund for 1 to 64 paid order line items.
type CreateParams struct {
	LineItems     []CreateLineItem  `json:"line_items"`
	OrderID       string            `json:"order_id"`
	Reason        Reason            `json:"reason"`
	CustomData    map[string]string `json:"custom_data,omitempty"`
	ReasonDetails string            `json:"reason_details,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	RequestMeta   *request.Meta     `json:"request_meta,omitempty"`
}

// CancelRefundRequest cancels a refund that has not begun processing.
type CancelParams struct {
	RefundID    string        `json:"refund_id"`
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// LookupRefundRequest identifies the refund to retrieve.
type LookupParams struct {
	RefundID string `json:"refund_id"`
}

// PageRefundsRequest selects a one-based refund page. PageNumber is required.
type PageParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size,omitempty"`
}
