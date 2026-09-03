package inttegro

import (
	"github.com/zebodotdev/inttegro-sdk-go/v4/bankaccounts"
	"github.com/zebodotdev/inttegro-sdk-go/v4/paymentmethods"
	"github.com/zebodotdev/inttegro-sdk-go/v4/wallets"
)

// Backward-compatible aliases keep existing root imports working while the
// focused packages provide the canonical public homes for these types.
type MobileMoneyNetwork = paymentmethods.MobileMoneyNetwork

const (
	MobileMoneyNetworkAirtel   = paymentmethods.MobileMoneyNetworkAirtel
	MobileMoneyNetworkMTN      = paymentmethods.MobileMoneyNetworkMTN
	MobileMoneyNetworkTelecel  = paymentmethods.MobileMoneyNetworkTelecel
	MobileMoneyNetworkVodafone = paymentmethods.MobileMoneyNetworkVodafone
)

type WalletType = wallets.Type

const WalletTypeMobileMoney = wallets.TypeMobileMoney

type WalletConfig = wallets.Config
type WalletMobileMoney = wallets.MobileMoney

type BankAccountType = bankaccounts.Type

const BankAccountTypeGhanaBankAccount = bankaccounts.TypeGhanaBankAccount

type BankAccountOwnerAddress = bankaccounts.OwnerAddress
type BankAccountOwner = bankaccounts.Owner
type GhanaBankAccount = bankaccounts.GhanaBankAccount
type BankAccountConfig = bankaccounts.Config
