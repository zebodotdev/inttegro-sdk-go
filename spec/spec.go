// Package spec provides country capability specifications and operations.
package spec

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service               = inttegro.SpecService
	PaymentMethodSettings = inttegro.PaymentMethodSettings
	PaymentMethodType     = inttegro.PaymentMethodTypeSetting
	PayoutSchedule        = inttegro.PayoutScheduleSpec
	BankBranch            = inttegro.CountryBankBranch
	Bank                  = inttegro.CountryBank
	BankDirectory         = inttegro.CountryBankDirectory
	Country               = inttegro.CountrySpecification
	Resource              = inttegro.CountrySpecification
)
