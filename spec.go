package commerce

import "context"

// SpecService provides access to Commerce platform specifications.
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
//	ghana := countries["GH"]
//	fmt.Printf("Ghana currencies: %v\n", ghana.Currencies)
//	fmt.Printf("Ghana payment methods: %v\n", ghana.PaymentMethods)
type SpecService struct {
	client *Client
}

// Countries retrieves Commerce capabilities for all supported countries.
//
// Returns a map of country code to specification. Use this to discover
// supported currencies, payment methods, and payout options before building
// your integration.
func (s *SpecService) Countries(ctx context.Context) (map[string]CountrySpecification, error) {
	var resp struct {
		Countries map[string]CountrySpecification `json:"countries"`
	}
	if err := s.client.do(ctx, "POST", "/spec/countries", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return resp.Countries, nil
}
