package spec

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
