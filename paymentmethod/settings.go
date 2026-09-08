package paymentmethod

// Settings contains acceptance configuration for every payment method type.
type Settings struct {
	MobileMoney *TypeSetting `json:"mobile_money,omitempty"`
	BankAccount *TypeSetting `json:"bank_account,omitempty"`
	Card        *TypeSetting `json:"card,omitempty"`
	Motito      *TypeSetting `json:"motito,omitempty"`
}

// TypeSetting controls availability and confirmation for one payment method type.
type TypeSetting struct {
	Type        Type   `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	ConfirmsUse bool   `json:"confirms_use,omitempty"`
}
