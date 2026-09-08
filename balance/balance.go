// Package balance provides balance snapshot resources and operations.
package balance

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service   = inttegro.BalancesService
	Amount    = inttegro.BalanceAmount
	Breakdown = inttegro.BalanceBreakdown
	Snapshot  = inttegro.BalanceSnapshot
	Resource  = inttegro.BalanceSnapshot
)
