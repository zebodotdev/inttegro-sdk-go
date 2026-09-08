package financialaccount

type Type string

const (
	TypeWallet Type = "wallet"
	TypeBank   Type = "bank_account"
	TypeDosh   Type = "dosh_account"
)
