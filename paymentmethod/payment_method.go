// Package paymentmethod provides payment-method resources, values, and operations.
package paymentmethod

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service                   = inttegro.PaymentMethodsService
	Type                      = inttegro.PaymentMethodType
	MobileMoneyNetwork        = inttegro.MobileMoneyNetwork
	MobileMoneyParams         = inttegro.MobileMoneyParams
	Data                      = inttegro.PaymentMethodData
	PageParams                = inttegro.PaymentMethodPageParams
	Page                      = inttegro.PaymentMethodPage
	ActionParams              = inttegro.PaymentMethodActionParams
	Resource                  = inttegro.PaymentMethod
	TokenizeParams            = inttegro.TokenizePaymentMethodParams
	VerifyParams              = inttegro.VerifyPaymentMethodParams
	VerificationSession       = inttegro.PaymentMethodVerificationSession
	VerificationDelivery      = inttegro.PaymentMethodVerificationDelivery
	ConfirmVerificationParams = inttegro.ConfirmPaymentMethodVerificationParams
	LookupParams              = inttegro.LookupPaymentMethodParams
	DeleteParams              = inttegro.DeletePaymentMethodParams
	Deletion                  = inttegro.PaymentMethodDeletion
)

const (
	TypeMobileMoney = inttegro.PaymentMethodTypeMobileMoney
	TypeBankAccount = inttegro.PaymentMethodTypeBankAccount
	TypeCard        = inttegro.PaymentMethodTypeCard
	TypeMotito      = inttegro.PaymentMethodTypeMotito

	MobileMoneyNetworkAirtel   = inttegro.MobileMoneyNetworkAirtel
	MobileMoneyNetworkMTN      = inttegro.MobileMoneyNetworkMTN
	MobileMoneyNetworkTelecel  = inttegro.MobileMoneyNetworkTelecel
	MobileMoneyNetworkVodafone = inttegro.MobileMoneyNetworkVodafone
)
