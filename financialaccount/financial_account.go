// Package financialaccount provides financialaccount resources and operations.
package financialaccount

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/wallet"
)

type Type string

const (
	TypeWallet Type = "wallet"
)

const (
	TypeBank Type = "bank_account"
)

const (
	TypeDosh Type = "dosh_account"
)

// FinancialAccountsService manages payout destination accounts.
//
// Financial accounts are where your payouts are sent. Connect mobile money
// wallets, bank accounts, or Dosh wallets to receive settlement funds.
//
// Use this service to:
//   - Connect payout destination accounts
//   - Retrieve account details and verification status
//   - Manage account configuration
//
// Example:
//
//	account, err := client.FinancialAccounts.Create(ctx, financialaccount.CreateParams{
//	    Label:    "Primary Payout Account",
//	    Type:     financialaccount.TypeWallet,
//	    Reference: "main_wallet",
//	    Currency: "ghs",
//	    PushConfiguration: &financialaccount.PullPushConfig{Enabled: inttegro.Bool(true)},
//	    Wallet: &wallet.Config{...},
//	})
//
// Learn more: https://studio.inttegro.com/set-up-financial-account
type Service struct {
	client transport.Client
}

// Create connects a new financial account for receiving payouts.
//
// Returns the created account with verification requirements (if any).
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Lookup retrieves financial account details and verification status by ID.
func (s *Service) Lookup(ctx context.Context, accountID string) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/lookup", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Connect is an alias for Create. Both methods do the same thing.
func (s *Service) Connect(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/connect", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Update modifies a financial account (PATCH semantics).
func (s *Service) Update(ctx context.Context, payload map[string]any) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// EnablePush enables push configuration for payouts.
func (s *Service) EnablePush(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/enable_push", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePush disables push configuration for payouts.
func (s *Service) DisablePush(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/disable_push", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePushWithOptions disables push configuration with optional payout-destination handling.
func (s *Service) DisablePushWithOptions(ctx context.Context, params DisablePushParams) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/disable_push", params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EnablePull enables pull configuration for charges.
func (s *Service) EnablePull(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/enable_pull", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePull disables pull configuration for charges.
func (s *Service) DisablePull(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/disable_pull", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Disconnect permanently disconnects a financial account.
func (s *Service) Disconnect(ctx context.Context, params DisconnectParams) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/disconnect", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Reconnect clears local disconnected state for a previously disconnected financial account.
func (s *Service) Reconnect(ctx context.Context, accountID string) (*Resource, error) {
	var resp struct {
		Account Resource `json:"account"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/reconnect", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Archive is currently not implemented by the API (returns 501) but exposed for completeness.
func (s *Service) Archive(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/archive", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Page retrieves a paginated list of financial accounts.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/financial_accounts/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Verify is currently not implemented by the API (returns 501) but exposed for completeness.
func (s *Service) Verify(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.Do(ctx, "POST", "/financial_accounts/verify", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

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
	EnabledAt string `json:"enabled_at,omitempty"`

	// Mandate contains mandate details for pull authorization (optional).
	Mandate map[string]any `json:"mandate,omitempty"`
}

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

// FinancialAccount represents a connected payout destination account.
//
// Financial accounts must be verified before use. Some account types
// require document verification or test deposits.
type Resource struct {
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
	ArchivedAt *string `json:"archived_at,omitempty"`

	// DisconnectedAt is when the account was disconnected (ISO 8601, read-only).
	// Nil if the account is still active.
	DisconnectedAt *string `json:"disconnected_at,omitempty"`

	// CreatedAt is when the account was created (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`
}

// PageFinancialAccountsParams paginates financial accounts.
type PageParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

// FinancialAccountsPage holds paginated account data.
type Page struct {
	Number   int        `json:"number,omitempty"`
	Size     int        `json:"size,omitempty"`
	Accounts []Resource `json:"accounts,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
