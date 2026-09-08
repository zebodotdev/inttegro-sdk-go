// Package wallet provides wallet variants used by financial accounts.
package wallet

import "github.com/zebodotdev/inttegro-sdk-go/v4/wallets"

type (
	Type        = wallets.Type
	MobileMoney = wallets.MobileMoney
	Config      = wallets.Config
)

const TypeMobileMoney = wallets.TypeMobileMoney
