// Package wallet provides wallet values.
package wallet

import "github.com/zebodotdev/inttegro-sdk-go/v7/paymentmethod"

// Type identifies a wallet implementation.
type Type string

const TypeMobileMoney Type = "mobile_money"

// MobileMoney contains the mobile money details for a wallet.
type MobileMoney struct {
	ID            string                           `json:"id,omitempty"`
	AccountNumber string                           `json:"account_number"`
	Network       paymentmethod.MobileMoneyNetwork `json:"network"`
}

// Config describes a wallet attached to a financial account.
type Config struct {
	Type        Type         `json:"type"`
	MobileMoney *MobileMoney `json:"mobile_money,omitempty"`
}
