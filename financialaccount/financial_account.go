// Package financialaccount provides payout-destination account resources and operations.
package financialaccount

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service           = inttegro.FinancialAccountsService
	Type              = inttegro.FinancialAccountType
	PullPushConfig    = inttegro.PullPushConfig
	CreateParams      = inttegro.FinancialAccountCreateParams
	DisablePushParams = inttegro.FinancialAccountDisablePushParams
	DisconnectParams  = inttegro.FinancialAccountDisconnectParams
	Resource          = inttegro.FinancialAccount
	PageParams        = inttegro.PageFinancialAccountsParams
	Page              = inttegro.FinancialAccountsPage
)

const (
	TypeWallet = inttegro.FinancialAccountTypeWallet
	TypeBank   = inttegro.FinancialAccountTypeBank
	TypeDosh   = inttegro.FinancialAccountTypeDosh
)
