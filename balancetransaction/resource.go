package balancetransaction

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
	"github.com/zebodotdev/inttegro-sdk-go/v6/payout"
)

// BalanceTransaction represents a merchant balance entry caused by a payment or
// refund. Exactly one of PaymentID and RefundID is present, matching Type.
type BalanceTransaction struct {
	// ID is the unique balance transaction identifier (read-only).
	// Starts with "bt_". Example: "bt_abc123def456"
	ID string `json:"id"`

	// Type identifies the semantic source, not the accounting direction.
	Type Type `json:"type"`

	// PaymentID is present only when Type is payment.
	PaymentID string `json:"payment_id,omitempty"`

	// RefundID is present only when Type is refund.
	RefundID string `json:"refund_id,omitempty"`

	// PayoutID identifies the payout that claimed this transaction, when present.
	PayoutID string `json:"payout_id,omitempty"`

	// OrderID is the strongly referenced source order ID.
	OrderID string `json:"order_id"`

	// Amount is the transaction amount in the public money shape.
	Amount money.Amount `json:"amount"`

	// AvailableAt is when funds become eligible for payout (ISO 8601, read-only).
	AvailableAt *time.Time `json:"available_at,omitempty"`

	// ClaimedAt is when the transaction was claimed for payout.
	ClaimedAt *time.Time `json:"claimed_at,omitempty"`

	// PaidAt is when the transaction was paid out or otherwise settled.
	PaidAt *time.Time `json:"paid_at,omitempty"`

	// CreatedAt is when the balance transaction was created (ISO 8601).
	CreatedAt time.Time `json:"created_at"`

	// PayoutConfiguration is the routing used when this transaction is paid out.
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`
}

// SourceID returns the matching strong source reference. It returns false for
// incomplete or contradictory transaction values.
func (t BalanceTransaction) SourceID() (string, bool) {
	switch t.Type {
	case TypePayment:
		if t.PaymentID != "" && t.RefundID == "" {
			return t.PaymentID, true
		}
	case TypeRefund:
		if t.RefundID != "" && t.PaymentID == "" {
			return t.RefundID, true
		}
	}
	return "", false
}
