package commerce

import "context"

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
//	account, err := client.FinancialAccounts.Create(ctx, commerce.FinancialAccountCreateParams{
//	    Label:    "Primary Payout Account",
//	    Type:     commerce.FinancialAccountTypeWallet,
//	    Reference: "main_wallet",
//	    Currency: "ghs",
//	    PushConfiguration: &commerce.PullPushConfig{Enabled: commerce.Bool(true)},
//	    Wallet: &commerce.WalletConfig{...},
//	})
//
// Learn more: https://studio.inttegro.com/set-up-financial-account
type FinancialAccountsService struct {
	client *Client
}

// Create connects a new financial account for receiving payouts.
//
// Returns the created account with verification requirements (if any).
func (s *FinancialAccountsService) Create(ctx context.Context, params FinancialAccountCreateParams) (*FinancialAccount, error) {
	var resp struct {
		Account FinancialAccount `json:"account"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Lookup retrieves financial account details and verification status by ID.
func (s *FinancialAccountsService) Lookup(ctx context.Context, accountID string) (*FinancialAccount, error) {
	var resp struct {
		Account FinancialAccount `json:"account"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/lookup", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Connect is an alias for Create. Both methods do the same thing.
func (s *FinancialAccountsService) Connect(ctx context.Context, params FinancialAccountCreateParams) (*FinancialAccount, error) {
	var resp struct {
		Account FinancialAccount `json:"account"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/connect", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Update modifies a financial account (PATCH semantics).
func (s *FinancialAccountsService) Update(ctx context.Context, payload map[string]any) (*FinancialAccount, error) {
	var resp struct {
		Account FinancialAccount `json:"account"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// EnablePush enables push configuration for payouts.
func (s *FinancialAccountsService) EnablePush(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/enable_push", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePush disables push configuration for payouts.
func (s *FinancialAccountsService) DisablePush(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/disable_push", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePushWithOptions disables push configuration with optional payout-destination handling.
func (s *FinancialAccountsService) DisablePushWithOptions(ctx context.Context, params FinancialAccountDisablePushParams) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/disable_push", params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EnablePull enables pull configuration for charges.
func (s *FinancialAccountsService) EnablePull(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/enable_pull", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DisablePull disables pull configuration for charges.
func (s *FinancialAccountsService) DisablePull(ctx context.Context, accountID string) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/disable_pull", map[string]string{"account_id": accountID}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Disconnect permanently disconnects a financial account.
func (s *FinancialAccountsService) Disconnect(ctx context.Context, params FinancialAccountDisconnectParams) (*FinancialAccount, error) {
	var resp struct {
		Account FinancialAccount `json:"account"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/disconnect", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Account, nil
}

// Archive is currently not implemented by the API (returns 501) but exposed for completeness.
func (s *FinancialAccountsService) Archive(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/archive", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Page retrieves a paginated list of financial accounts.
func (s *FinancialAccountsService) Page(ctx context.Context, params PageFinancialAccountsParams) (*FinancialAccountsPage, error) {
	var resp struct {
		Page FinancialAccountsPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/financial_accounts/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Verify is currently not implemented by the API (returns 501) but exposed for completeness.
func (s *FinancialAccountsService) Verify(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/financial_accounts/verify", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
