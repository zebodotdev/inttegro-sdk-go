// Package balancetransaction provides balance-transaction resources and operations.
package balancetransaction

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service    = inttegro.BalanceTransactionsService
	Type       = inttegro.BalanceTransactionType
	Resource   = inttegro.BalanceTransaction
	PageParams = inttegro.BalanceTransactionPageParams
)

const (
	TypePayment = inttegro.BalanceTransactionTypePayment
	TypeRefund  = inttegro.BalanceTransactionTypeRefund
)
