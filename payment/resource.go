package payment

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/balancetransaction"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payout"
)

// Payment represents payment details and status for an order.
//
// Every paid order has an associated payment object tracking the charge
// lifecycle, attempts, and any required customer actions.
type Resource struct {
	// ID is the unique payment identifier.
	// Starts with "py_". Example: "py_abc123def456"
	ID string `json:"id,omitempty"`

	// Status is the payment's current state.
	// Values: "initiated", "requires_action", "processing", "paid", "failed"
	Status Status `json:"status,omitempty"`

	// StatementDescriptor is what appears on the customer's statement.
	// Maximum 22 characters.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// Amount is the charged amount.
	Amount *money.Amount `json:"amount,omitempty"`

	// PaymentMethod is the charged payment method details.
	PaymentMethod *paymentmethod.Resource `json:"payment_method,omitempty"`

	// LatestAttempt is the most recent payment attempt.
	// Nil if no attempts yet.
	LatestAttempt *Attempt `json:"latest_attempt,omitempty"`

	// NextAction describes any required customer action.
	// Nil if no action required (payment is processing or complete).
	NextAction *NextAction `json:"next_action,omitempty"`

	// BalanceTransaction is the resulting balance entry when payment succeeds.
	// Used for tracking payouts and available balance.
	BalanceTransaction *balancetransaction.Resource `json:"balance_transaction,omitempty"`

	// PayoutConfiguration is the payout setup used for this payment (if applicable).
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`

	// InitiatedAt is when payment was first attempted (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// ExecutedAt is when payment was submitted to the network (ISO 8601).
	// Nil if not yet executed.
	ExecutedAt *string `json:"executed_at,omitempty"`

	// PaidAt is when payment was confirmed successful (ISO 8601).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// FailedAt is when payment was marked failed (ISO 8601).
	// Nil if not failed.
	FailedAt *string `json:"failed_at,omitempty"`
}
