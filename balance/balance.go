// Package balance provides balance resources and operations.
package balance

import (
	"context"
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

// Service retrieves the current application balance.
type Service struct {
	client transport.Client
}

// BalanceAmount represents a balance amount in minor units.
type Amount struct {
	Amount int64 `json:"amount"`
}

// BalanceBreakdown is a per-currency breakdown of balances.
type Breakdown struct {
	Available                  Amount    `json:"available"`
	Pending                    Amount    `json:"pending"`
	Reserved                   Amount    `json:"reserved"`
	Refund                     Amount    `json:"refund"`
	IncludesTransactionsBefore time.Time `json:"includes_transactions_before"`
}

// Balance is the current application balance.
type Balance struct {
	GHS Breakdown `json:"ghs"`
}

// Get retrieves the current balances snapshot.
func (s *Service) Get(ctx context.Context) (*Balance, error) {
	var resp struct {
		Balances Balance `json:"balances"`
	}
	if err := s.client.Do(ctx, "POST", "/balances", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Balances, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
