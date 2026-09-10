package payment

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v7/paymentmethod"
)

type PaymentMethod struct {
	ID          string                    `json:"id"`
	BankAccount *PaymentMethodBankAccount `json:"bank_account,omitempty"`
	Card        *PaymentMethodCard        `json:"card,omitempty"`
	CreatedAt   time.Time                 `json:"created_at"`
	CustomerID  string                    `json:"customer_id"`
	MobileMoney *PaymentMethodMobileMoney `json:"mobile_money,omitempty"`
	Owner       *PaymentMethodOwner       `json:"owner,omitempty"`
	Type        paymentmethod.Type        `json:"type"`
	Verified    bool                      `json:"verified"`
	VerifiedAt  *time.Time                `json:"verified_at,omitempty"`
}

type PaymentMethodMobileMoney struct {
	Network       paymentmethod.MobileMoneyNetwork `json:"network"`
	AccountNumber string                           `json:"account_number"`
	Last4         string                           `json:"last4"`
}

type PaymentMethodBankAccount struct {
	Type             string                         `json:"type"`
	GhanaBankAccount *PaymentMethodGhanaBankAccount `json:"ghana_bank_account,omitempty"`
}

type PaymentMethodGhanaBankAccount struct {
	AccountNumber string `json:"account_number"`
	Branch        string `json:"branch,omitempty"`
	Name          string `json:"name,omitempty"`
	SortCode      string `json:"sort_code,omitempty"`
	SwiftCode     string `json:"swift_code,omitempty"`
}

type PaymentMethodCard struct{}

type PaymentMethodOwner struct {
	Name    string   `json:"name"`
	Address *Address `json:"address,omitempty"`
}
