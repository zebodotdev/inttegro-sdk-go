package paymentmethod

import (
	"github.com/zebodotdev/inttegro-sdk-go/v6/bankaccount"
)

type Page struct {
	Number         int             `json:"number,omitempty"`
	Size           int             `json:"size,omitempty"`
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"`
}

// PaymentMethod represents a tokenized payment method.
//
// Payment methods are saved customer payment instruments that can be
// charged repeatedly without re-entering payment details.
//
// Payment methods must be verified before use (except when confirms_use
// is false). Verification sends an OTP to confirm the customer owns the
// payment instrument.
type PaymentMethod struct {
	// ID is the unique payment method identifier.
	// Starts with "pm_". Example: "pm_abc123def456"
	ID string `json:"id"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id"`

	// Type is the payment method category.
	// Values: "mobile_money", "bank_account", "card", "motito"
	Type Type `json:"type"`

	// MobileMoney contains wallet details when Type is "mobile_money".
	MobileMoney *struct {
		AccountNumber string             `json:"account_number,omitempty"`
		Network       MobileMoneyNetwork `json:"network,omitempty"`
	} `json:"mobile_money,omitempty"`

	// BankAccount contains bank details when Type is "bank_account".
	BankAccount *struct {
		GhanaBankAccount *struct {
			Branch        string `json:"branch,omitempty"`
			Name          string `json:"name,omitempty"`
			AccountNumber string `json:"account_number,omitempty"`
			SortCode      string `json:"sort_code,omitempty"`
			SwiftCode     string `json:"swift_code,omitempty"`
		} `json:"ghana_bank_account,omitempty"`
		Type bankaccount.Type `json:"type,omitempty"`
	} `json:"bank_account,omitempty"`

	// Card contains card details when Type is "card".
	Card *struct {
		Brand     string  `json:"brand,omitempty"`
		ExpiresOn *string `json:"expires_on,omitempty"`
		Issuer    *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"issuer,omitempty"`
		Owner *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
		} `json:"owner,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"card,omitempty"`

	// Verification contains verification metadata when available.
	Verification *struct {
		CompletedAt *string `json:"completed_at,omitempty"`
		InitiatedAt *string `json:"initiated_at,omitempty"`
		Mechanism   string  `json:"mechanism,omitempty"`
		RequestID   string  `json:"request_id,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"verification,omitempty"`

	// CustomData holds arbitrary metadata attached to the payment method.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// ExpiresOn is when this payment method expires (ISO 8601), if applicable.
	ExpiresOn *string `json:"expires_on,omitempty"`

	// CreatedAt is the tokenization timestamp (ISO 8601).
	CreatedAt string `json:"created_at"`

	// Verified indicates whether the payment method has been verified.
	// Unverified payment methods cannot be charged (unless confirms_use is false).
	Verified bool `json:"verified"`

	// VerifiedAt is the verification completion timestamp (ISO 8601).
	// Nil if not yet verified.
	VerifiedAt *string `json:"verified_at,omitempty"`
}
