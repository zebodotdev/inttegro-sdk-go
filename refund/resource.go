package refund

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v7/money"
)

// RefundLineItem is one immutable order-line allocation in a refund.
type LineItem struct {
	ID                 string       `json:"id"`
	OrderLineItemID    string       `json:"order_line_item_id"`
	OriginalAmountPaid money.Amount `json:"original_amount_paid"`
	RefundAmount       money.Amount `json:"refund_amount"`
	Reason             *Reason      `json:"reason,omitempty"`
	ReasonDetails      string       `json:"reason_details,omitempty"`
}

// Refund is the canonical refund object embedded in order responses.
type Refund struct {
	ID            string            `json:"id"`
	OrderID       string            `json:"order_id"`
	Status        Status            `json:"status"`
	Total         money.Amount      `json:"total"`
	LineItems     []LineItem        `json:"line_items"`
	Reason        Reason            `json:"reason"`
	ReasonDetails string            `json:"reason_details,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	CustomData    map[string]string `json:"custom_data,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	ProcessingAt  *time.Time        `json:"processing_at,omitempty"`
	SucceededAt   *time.Time        `json:"succeeded_at,omitempty"`
	FailedAt      *time.Time        `json:"failed_at,omitempty"`
	CanceledAt    *time.Time        `json:"canceled_at,omitempty"`
}

// RefundPage contains one page of refunds.
type Page struct {
	Number  int      `json:"number"`
	Refunds []Refund `json:"refunds"`
	Size    int      `json:"size"`
}
