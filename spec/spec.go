// Package spec provides spec resources and operations.
package spec

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
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
//	ghana := countries["GH"]
//	fmt.Printf("Ghana currencies: %v\n", ghana.Currencies)
//	fmt.Printf("Ghana payment methods: %v\n", ghana.PaymentMethods)
type Service struct {
	client transport.Client
}

// Countries retrieves Inttegro capabilities for all supported countries.
//
// Returns a map of country code to specification. Use this to discover
// supported currencies, payment methods, and payout options before building
// your integration.
func (s *Service) Countries(ctx context.Context) (map[string]Resource, error) {
	var resp struct {
		Countries map[string]Resource `json:"countries"`
	}
	if err := s.client.Do(ctx, "POST", "/spec/countries", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return resp.Countries, nil
}

// CountryBankBranch describes a bank branch available in a country bank directory.
type BankBranch struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SortCode string `json:"sort_code,omitempty"`
}

// CountryBank describes a banking institution available in a country bank directory.
type Bank struct {
	ID             string       `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	SwiftCode      string       `json:"swift_code,omitempty"`
	SortCodePrefix string       `json:"sort_code_prefix,omitempty"`
	Branches       []BankBranch `json:"branches,omitempty"`
}

// CountryBankDirectory describes bank reference data for country-specific bank accounts.
type BankDirectory struct {
	BankAccountType string `json:"bank_account_type,omitempty"`
	CodeScheme      string `json:"code_scheme,omitempty"`
	Items           []Bank `json:"items,omitempty"`
}

// CountrySpecification describes supported Inttegro features for a country.
//
// Use this to discover supported currencies, payment methods, payout schedules,
// and other country-specific capabilities before integrating.
//
// Query with Spec.Countries() to get all country specifications.
type Resource struct {
	// CountryCode is the two-letter ISO 3166-1 alpha-2 code (read-only).
	// Example: "GH", "KE", "UG", "US"
	CountryCode string `json:"country_code,omitempty"`

	// CountryName is the full country name (read-only).
	// Example: "Ghana", "Kenya", "Uganda"
	CountryName string `json:"country_name,omitempty"`

	// Currencies lists supported currency codes (read-only).
	// Example: ["ghs", "usd"] for Ghana
	Currencies []string `json:"currencies,omitempty"`

	// PaymentMethods lists supported payment method types (read-only).
	// Example: ["mobile_money", "bank_account"]
	PaymentMethods []string `json:"payment_methods,omitempty"`

	// PayoutSchedules lists available payout schedule types (read-only).
	// Example: ["weekly", "manual"]
	PayoutSchedules []string `json:"payout_schedules,omitempty"`

	// BTAgingSpecs lists balance transaction aging options (read-only).
	// Example: ["t+7", "t+14"] for 7-day or 14-day aging
	BTAgingSpecs []string `json:"bt_aging_specs,omitempty"`

	// LegalEntityTypes lists supported business types (read-only).
	// Structure varies by country requirements.
	LegalEntityTypes []map[string]any `json:"legal_entity_types,omitempty"`

	// FinancialAccountTypes lists supported payout destination types (read-only).
	// Details wallet, bank_account, and dosh_account configurations.
	FinancialAccountTypes []map[string]any `json:"financial_account_types,omitempty"`

	// IDDocumentTypes lists accepted identification documents (read-only).
	// Used for KYC/verification requirements.
	IDDocumentTypes []map[string]any `json:"id_document_types,omitempty"`

	// Banks lists country-specific bank reference data, when available.
	// Ghana uses bank_account_type "ghana_bank_account" and sort code branches.
	Banks *BankDirectory `json:"banks,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
