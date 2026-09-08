package order

import (
	"github.com/zebodotdev/inttegro-sdk-go/v6/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v6/financialaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v6/wallet"
)

// OrderPayoutSettings overrides payout configuration for a single order.
type PayoutSettings struct {
	// Destination specifies where payout funds should be sent (optional).
	Destination *PayoutDestination `json:"destination,omitempty"`

	// EnableFX controls foreign exchange conversion for the payout (optional).
	EnableFX *bool `json:"enable_fx,omitempty"`
}

// OrderPayoutDestination sets the payout destination for the order.
type PayoutDestination struct {
	// FinancialAccountID references an existing financial account (optional).
	FinancialAccountID string `json:"financial_account_id,omitempty"`

	// FinancialAccountData provides inline financial account details (optional).
	FinancialAccountData *PayoutFinancialAccount `json:"financial_account_data,omitempty"`
}

// OrderPayoutFinancialAccount defines inline payout destination details.
type PayoutFinancialAccount struct {
	// Type specifies the account type (wallet, bank_account, dosh_account).
	Type financialaccount.Type `json:"type"`

	// Wallet contains mobile money details when Type is "wallet".
	Wallet *wallet.Config `json:"wallet,omitempty"`

	// BankAccount contains bank account details when Type is "bank_account".
	BankAccount *bankaccount.Config `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details when Type is "dosh_account".
	DoshAccount map[string]any `json:"dosh_account,omitempty"`
}
