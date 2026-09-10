package payout

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
)

// Payout represents a settlement transfer to your bank or mobile money account.
//
// Payouts move funds from your Inttegro balance to your financial accounts.
// Each payout contains one or more balance transactions that have aged
// past the dispute window.
type Payout struct {
	// ID is the unique payout identifier (read-only).
	// Starts with "po_". Example: "po_abc123def456"
	ID string `json:"id,omitempty"`

	// ApplicationID is your application's ID (read-only).
	ApplicationID string `json:"application_id,omitempty"`

	// DestinationID is the receiving financial account's ID (read-only).
	// Corresponds to a financial account you've connected.
	DestinationID string `json:"destination_id,omitempty"`

	// Amount is the payout total (read-only).
	Amount *money.Amount `json:"amount,omitempty"`

	// Status is the payout's current state (read-only).
	// Values include "scheduled", "initiated", "processing", "succeeded", "failed", "canceled"
	Status Status `json:"status,omitempty"`

	// InitiatedBy indicates who triggered the payout (read-only).
	// Values: "schedule" (automatic), "manual" (you initiated)
	InitiatedBy string `json:"initiated_by,omitempty"`

	// LatestAttemptID is the most recent execution attempt's ID (read-only).
	LatestAttemptID string `json:"latest_attempt_id,omitempty"`

	// LatestError contains error details if payout failed (read-only).
	// Nil if payout succeeded or is still processing.
	LatestError any `json:"latest_error,omitempty"`

	// InitiatedAt is when the payout was created (ISO 8601, read-only).
	InitiatedAt *time.Time `json:"initiated_at,omitempty"`

	// ExecuteAfter is the scheduled execution timestamp for queued payouts (ISO 8601, read-only).
	// Nil for immediate/manual payouts that are not scheduled.
	ExecuteAfter *time.Time `json:"execute_after,omitempty"`

	// ScheduledAt is when the payout was queued for execution (ISO 8601, read-only).
	// Nil when not scheduled.
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`

	// CanceledAt is when a scheduled payout was canceled (ISO 8601, read-only).
	// Nil unless the payout has status "canceled".
	CanceledAt *time.Time `json:"canceled_at,omitempty"`

	// MaxAmount is the maximum amount authorized for scheduled payouts (read-only).
	// This may differ from Amount when payout execution has not started.
	MaxAmount *money.Amount `json:"max_amount,omitempty"`

	// ExecutedAt is when the payout was submitted to the network (ISO 8601, read-only).
	// Nil if not yet executed.
	ExecutedAt *time.Time `json:"executed_at,omitempty"`

	// ExpectedAt is when the payout should arrive (ISO 8601, read-only).
	// Estimate based on network speed. Actual arrival may vary.
	ExpectedAt *time.Time `json:"expected_at,omitempty"`

	// SucceededAt is when the payout was confirmed (ISO 8601, read-only).
	// Nil if not yet succeeded.
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`

	// BalanceTransactionIDs lists the included balance transactions (read-only).
	// These are the source funds being paid out.
	BalanceTransactionIDs []string `json:"balance_transaction_ids,omitempty"`
}
