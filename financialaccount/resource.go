package financialaccount

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v6/wallet"
)

// FinancialAccount represents a connected payout destination account.
//
// Financial accounts must be verified before use. Some account types
// require document verification or test deposits.
type FinancialAccount struct {
	// ID is the unique financial account identifier (read-only).
	// Starts with "fa_". Example: "fa_abc123def456"
	ID string `json:"id,omitempty"`

	// Label is the account's descriptive name.
	Label string `json:"label,omitempty"`

	// Type is the account category.
	// Values: "wallet", "bank_account", "dosh_account"
	Type Type `json:"type,omitempty"`

	// Reference is your internal identifier.
	Reference string `json:"reference,omitempty"`

	// Currency is the account's currency (ISO 4217, lowercase).
	Currency string `json:"currency,omitempty"`

	// Description provides additional context.
	Description string `json:"description,omitempty"`

	// PullConfiguration indicates whether Inttegro can debit this account.
	PullConfiguration *PullPushConfig `json:"pull_configuration,omitempty"`

	// PushConfiguration indicates whether Inttegro can credit this account.
	PushConfiguration *PullPushConfig `json:"push_configuration,omitempty"`

	// Wallet contains mobile money wallet details (when Type is "wallet").
	Wallet *wallet.Config `json:"wallet,omitempty"`

	// BankAccount contains bank account details (when Type is "bank_account").
	BankAccount *bankaccount.Config `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details (when Type is "dosh_account").
	DoshAccount map[string]any `json:"dosh_account,omitempty"`

	// CustomData contains optional key-value metadata for tracking.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Owner contains financial account owner information.
	Owner *bankaccount.Owner `json:"owner,omitempty"`

	// Verification contains verification status and requirements.
	// Nil if verification not started or not required.
	Verification any `json:"verification,omitempty"`

	// ArchivedAt is when the account was archived (ISO 8601, read-only).
	// Archived accounts cannot receive new payouts.
	// Nil if not archived.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`

	// DisconnectedAt is when the account was disconnected (ISO 8601, read-only).
	// Nil if the account is still active.
	DisconnectedAt *time.Time `json:"disconnected_at,omitempty"`

	// CreatedAt is when the account was created (ISO 8601, read-only).
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// FinancialAccountsPage holds paginated account data.
type Page struct {
	Number   int                `json:"number,omitempty"`
	Size     int                `json:"size,omitempty"`
	Accounts []FinancialAccount `json:"accounts,omitempty"`
}
