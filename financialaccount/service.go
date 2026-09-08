package financialaccount

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
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

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
