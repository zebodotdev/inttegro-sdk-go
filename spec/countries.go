package spec

// CountrySpecification describes supported Inttegro features for a country.
//
// Use this to discover supported currencies, payment methods, payout schedules,
// and other country-specific capabilities before integrating.
//
// Query with Client.Spec.Countries() to get all country specifications.
type CountrySpecification struct {
	// CountryCode is the two-letter ISO 3166-1 alpha-2 code (read-only).
	// Example: "gh", "ke", "ug", "us"
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
	// Example: ["individual", "company"]
	LegalEntityTypes []string `json:"legal_entity_types,omitempty"`

	// FinancialAccountTypes lists supported payout destination types (read-only).
	// Example: ["bank_account", "mobile_money"]
	FinancialAccountTypes []string `json:"financial_account_types,omitempty"`

	// IDDocumentTypes lists accepted identification documents (read-only).
	// Example: ["passport", "national_id"]
	IDDocumentTypes []string `json:"id_document_types,omitempty"`

	// Banks lists country-specific bank reference data, when available.
	// Ghana uses bank_account_type "ghana_bank_account" and sort code branches.
	Banks *BankDirectory `json:"banks,omitempty"`
}

// CountrySpecifications contains country specifications keyed by lowercase
// ISO 3166-1 alpha-2 country code.
type CountrySpecifications map[string]CountrySpecification

// Spec is retained as a compatibility alias for CountrySpecification.
// Deprecated: use CountrySpecification.
type Spec = CountrySpecification
