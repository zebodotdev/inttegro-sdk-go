// Package balance provides balance resources and operations.
package balance

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

// BalancesService retrieves balance snapshots across currencies.
type Service struct {
	client transport.Client
}

// BalanceAmount represents a balance amount in minor units.
type Amount struct {
	Amount int64 `json:"amount"`
}

// BalanceBreakdown is a per-currency breakdown of balances.
type Breakdown struct {
	Available                  *Amount `json:"available,omitempty"`
	Pending                    *Amount `json:"pending,omitempty"`
	Reserved                   *Amount `json:"reserved,omitempty"`
	Refund                     *Amount `json:"refund,omitempty"`
	IncludesTransactionsBefore string  `json:"includes_transactions_before,omitempty"`
}

// Balance is the current balance breakdown keyed by currency.
type Balance struct {
	Balances map[string]Breakdown `json:"balances"`
}

// Get retrieves the current balances snapshot.
func (s *Service) Get(ctx context.Context) (*Balance, error) {
	var resp Balance
	if err := s.client.Do(ctx, "POST", "/balances", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
