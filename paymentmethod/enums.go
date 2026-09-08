package paymentmethod

type Type string

const (
	TypeMobileMoney Type = "mobile_money"
	TypeBankAccount Type = "bank_account"
	TypeCard        Type = "card"
	TypeMotito      Type = "motito"
)
