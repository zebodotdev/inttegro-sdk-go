package paymentmethod

import "time"

// PaymentMethod is a tokenized payment instrument tied to a customer.
type PaymentMethod struct {
	ID           string            `json:"id"`
	Active       bool              `json:"active"`
	ArchivedAt   *time.Time        `json:"archived_at,omitempty"`
	BankAccount  *BankAccount      `json:"bank_account,omitempty"`
	Card         *Card             `json:"card,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	CustomerID   string            `json:"customer_id"`
	CustomData   map[string]string `json:"custom_data,omitempty"`
	Ephemeral    bool              `json:"ephemeral,omitempty"`
	ExpiresOn    *time.Time        `json:"expires_on,omitempty"`
	MobileMoney  *MobileMoney      `json:"mobile_money,omitempty"`
	Owner        *Owner            `json:"owner,omitempty"`
	Supplied     *Supplied         `json:"supplied,omitempty"`
	Type         Type              `json:"type"`
	Verification *Verification     `json:"verification,omitempty"`
	VerifiedAt   *time.Time        `json:"verified_at,omitempty"`
}

type Page struct {
	Number         int             `json:"number"`
	Size           int             `json:"size"`
	PaymentMethods []PaymentMethod `json:"payment_methods"`
}
