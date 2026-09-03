// Package bankaccounts defines bank-account variants used by financial accounts.
package bankaccounts

// Type identifies a bank-account implementation.
type Type string

const TypeGhanaBankAccount Type = "ghana_bank_account"

// OwnerAddress captures a bank-account holder's address.
type OwnerAddress struct {
	ID            string `json:"id,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
	Name          string `json:"name"`
	Phone         string `json:"phone,omitempty"`
	Line1         string `json:"line_1"`
	Line2         string `json:"line_2,omitempty"`
	City          string `json:"city"`
	Region        string `json:"region"`
	PostCode      string `json:"post_code,omitempty"`
	Country       string `json:"country"`
}

// Owner captures bank-account holder information.
type Owner struct {
	Name    string       `json:"name"`
	Address OwnerAddress `json:"address"`
}

// GhanaBankAccount contains Ghana bank-account details.
type GhanaBankAccount struct {
	BankName  string `json:"bank_name,omitempty"`
	Branch    string `json:"branch,omitempty"`
	Number    string `json:"number"`
	SortCode  string `json:"sort_code,omitempty"`
	SwiftCode string `json:"swift_code,omitempty"`
	Holder    Owner  `json:"holder"`
}

// Config describes a bank account attached to a financial account.
type Config struct {
	ID               string            `json:"id,omitempty"`
	Type             Type              `json:"type"`
	GhanaBankAccount *GhanaBankAccount `json:"ghana_bank_account,omitempty"`
}
