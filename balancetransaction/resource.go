package balancetransaction

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payout"
)

// BalanceTransaction represents a merchant balance entry caused by a payment or
// refund. Exactly one of PaymentID and RefundID is present, matching Type.
type Resource struct {
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

	// Deprecated: the reviewed API does not return amount_expected. Use Amount.
	AmountExpected *money.Amount `json:"amount_expected,omitempty"`

	// Deprecated: the reviewed API does not return amount_available. Use Amount.
	AmountAvailable *money.Amount `json:"amount_available,omitempty"`

	// AvailableAt is when funds become eligible for payout (ISO 8601, read-only).
	AvailableAt *string `json:"available_at,omitempty"`

	// ClaimedAt is when the transaction was claimed for payout.
	ClaimedAt *string `json:"claimed_at,omitempty"`

	// PaidAt is when the transaction was paid out or otherwise settled.
	PaidAt *string `json:"paid_at,omitempty"`

	// CreatedAt is when the balance transaction was created (ISO 8601).
	CreatedAt string `json:"created_at"`

	// Deprecated: the reviewed API does not return payout_configuration on balance transactions.
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`
}

// SourceID returns the matching strong source reference. It returns false for
// incomplete or contradictory transaction values.
func (t Resource) SourceID() (string, bool) {
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
