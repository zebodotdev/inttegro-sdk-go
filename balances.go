package inttegro

import "context"

// BalancesService retrieves balance snapshots across currencies.
type BalancesService struct {
	client *Client
}

// BalanceAmount represents a balance amount in minor units.
type BalanceAmount struct {
	Amount int64 `json:"amount"`
}

// BalanceBreakdown is a per-currency breakdown of balances.
type BalanceBreakdown struct {
	Available                  *BalanceAmount `json:"available,omitempty"`
	Pending                    *BalanceAmount `json:"pending,omitempty"`
	Reserved                   *BalanceAmount `json:"reserved,omitempty"`
	Refund                     *BalanceAmount `json:"refund,omitempty"`
	IncludesTransactionsBefore string         `json:"includes_transactions_before,omitempty"`
}

// BalancesResponse is returned from /balances.
type BalancesResponse struct {
	Balances map[string]BalanceBreakdown `json:"balances"`
}

// Get retrieves the current balances snapshot.
func (s *BalancesService) Get(ctx context.Context) (*BalancesResponse, error) {
	var resp BalancesResponse
	if err := s.client.do(ctx, "POST", "/balances", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
