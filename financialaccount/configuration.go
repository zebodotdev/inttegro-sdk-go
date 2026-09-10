package financialaccount

import "time"

// PullPushConfig configures whether an account can send or receive funds.
//
// Pull configuration controls whether Inttegro can debit the account.
// Push configuration controls whether Inttegro can credit the account.
// Most payout destinations only need push enabled.
type PullPushConfig struct {
	// Enabled indicates whether the operation is allowed.
	// For payout destinations, PushConfiguration.Enabled should be true.
	// For payment sources, PullConfiguration.Enabled should be true.
	Enabled *bool `json:"enabled,omitempty"`

	// EnabledAt indicates when this configuration was enabled (read-only).
	EnabledAt *time.Time `json:"enabled_at,omitempty"`

	// Mandate contains mandate details for pull authorization (optional).
	Mandate map[string]any `json:"mandate,omitempty"`
}
