package spec

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/internal/transport"
)

// SpecService provides access to Inttegro platform specifications.
//
// Use this service to discover supported features before integrating:
//   - Supported countries and currencies
//   - Available payment methods by country
//   - Payout schedules and aging options
//   - Required documents and account types
//
// Example:
//
//	countries, err := client.Spec.Countries(ctx)
//	if err != nil {
//	    return err
//	}
//	ghana := countries["gh"]
//	fmt.Printf("Ghana currencies: %v\n", ghana.Currencies)
//	fmt.Printf("Ghana payment methods: %v\n", ghana.PaymentMethods)
type Service struct {
	client transport.Client
}

// Countries retrieves Inttegro capabilities for all supported countries.
//
// Returns country specifications keyed by lowercase country code. Use this to
// discover supported currencies, payment methods, and payout options before
// building your integration.
func (s *Service) Countries(ctx context.Context) (CountrySpecifications, error) {
	var resp struct {
		Countries CountrySpecifications `json:"countries"`
	}
	if err := s.client.Do(ctx, "POST", "/spec/countries", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return resp.Countries, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
