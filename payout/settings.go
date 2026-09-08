package payout

// PayoutSettings contains your complete payout configuration.
//
// Controls when payouts happen, where funds go, and whether currency
// conversion is enabled.
type Settings struct {
	// ID is the settings identifier (read-only).
	ID string `json:"id,omitempty"`

	// FxEnabled indicates whether currency conversion is enabled (read-only).
	// When true, can receive payouts in different currency than source funds.
	// Requires FX-enabled destination accounts.
	FxEnabled bool `json:"fx_enabled,omitempty"`

	// Destinations maps currencies to financial account IDs.
	// Key: currency code (e.g., "ghs", "usd")
	// Value: financial account ID (e.g., "fa_abc123")
	// Example: {"ghs": "fa_abc123", "usd": "fa_def456"}
	Destinations map[string]string `json:"destinations,omitempty"`

	// Schedule describes payout timing and frequency.
	Schedule *Schedule `json:"schedule,omitempty"`
}

// PayoutConfiguration describes payout routing and FX settings for a payment or balance transaction.
type Configuration struct {
	// EnableFX indicates whether FX conversion is enabled for this payout.
	EnableFX *bool `json:"enable_fx,omitempty"`

	// Destination specifies the financial account receiving the payout.
	Destination *Destination `json:"destination,omitempty"`
}

// PayoutDestination identifies the payout financial account.
type Destination struct {
	// FinancialAccountID is the ID of the destination financial account.
	FinancialAccountID string `json:"financial_account_id,omitempty"`
}
