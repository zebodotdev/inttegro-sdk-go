// Package bankaccount provides bank-account variants used by financial accounts.
package bankaccount

import "github.com/zebodotdev/inttegro-sdk-go/v4/bankaccounts"

type (
	Type             = bankaccounts.Type
	OwnerAddress     = bankaccounts.OwnerAddress
	Owner            = bankaccounts.Owner
	GhanaBankAccount = bankaccounts.GhanaBankAccount
	Config           = bankaccounts.Config
)

const TypeGhanaBankAccount = bankaccounts.TypeGhanaBankAccount
