package financialaccount

import (
	"github.com/zebodotdev/inttegro-sdk-go/v6/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v6/wallet"
)

// FinancialAccountCreateParams creates a payout destination account.
//
// Financial accounts are where your payouts are sent. Connect mobile money
// wallets, bank accounts, or Dosh wallets to receive settlement funds.
//
// Example (mobile money):
//
//	params := financialaccount.CreateParams{
//	    Label:       "Primary Payout Account",
//	    Type:        financialaccount.TypeWallet,
//	    Reference:   "main_wallet",
//	    Currency:    "ghs",
//	    Description: "Main MTN wallet for receiving payouts",
//	    PushConfiguration: &financialaccount.PullPushConfig{
//	        Enabled: inttegro.Bool(true),
//	    },
//	    Wallet: &wallet.Config{
//	        Type: wallet.TypeMobileMoney,
//	        MobileMoney: &wallet.MobileMoney{
//	            AccountNumber: "+233244123456",
//	            Network: paymentmethod.MobileMoneyNetworkMTN,
//	        },
//	    },
//	}
type CreateParams struct {
	// Label is a descriptive name for this account (required).
	// Displayed in dashboard and payout reports.
	// Example: "Primary Payout Account", "GHS Mobile Money"
	Label string `json:"label"`

	// Type specifies the account category (required).
	// Values: "wallet", "bank_account", "dosh_account"
	Type Type `json:"type"`

	// Reference is your internal identifier for this account (required).
	// Must be unique across your financial accounts.
	// Example: "primary_wallet", "backup_bank_01"
	Reference string `json:"reference"`

	// Currency is the account's currency (required).
	// Three-letter ISO 4217 code (lowercase).
	// Must match the currency you want to receive payouts in.
	// Example: "ghs", "usd", "kes"
	Currency string `json:"currency"`

	// Description provides additional context (optional).
	// Not displayed to customers, for internal use only.
	Description string `json:"description,omitempty"`

	// PullConfiguration controls whether Inttegro can debit this account (optional).
	// Typically false for payout destinations.
	PullConfiguration *PullPushConfig `json:"pull_configuration,omitempty"`

	// PushConfiguration controls whether Inttegro can credit this account (required).
	// Must be enabled for payout destinations.
	PushConfiguration *PullPushConfig `json:"push_configuration,omitempty"`

	// Wallet contains mobile money wallet details (required when Type is "wallet").
	Wallet *wallet.Config `json:"wallet,omitempty"`

	// BankAccount contains bank account details (required when Type is "bank_account").
	BankAccount *bankaccount.Config `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details (required when Type is "dosh_account").
	DoshAccount map[string]any `json:"dosh_account,omitempty"`

	// CustomData contains optional key-value metadata for tracking.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Owner contains financial account owner information (required for payouts).
	Owner *bankaccount.Owner `json:"owner,omitempty"`
}

// FinancialAccountDisablePushParams disables push configuration for a financial account.
type DisablePushParams struct {
	// AccountID is the financial account identifier (required).
	AccountID string `json:"account_id"`

	// UnsetAsPayoutDestination removes the account from payout destinations first.
	// Defaults to false when omitted.
	UnsetAsPayoutDestination *bool `json:"unset_as_payout_destination,omitempty"`
}

// FinancialAccountDisconnectParams disconnects a financial account.
type DisconnectParams struct {
	// AccountID is the financial account identifier (required).
	AccountID string `json:"account_id"`

	// UnsetAsPayoutDestination removes the account from payout destinations first.
	// Defaults to false when omitted.
	UnsetAsPayoutDestination *bool `json:"unset_as_payout_destination,omitempty"`
}

// PageFinancialAccountsParams paginates financial accounts.
type PageParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}
