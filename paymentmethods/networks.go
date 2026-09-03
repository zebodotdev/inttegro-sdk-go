// Package paymentmethods defines payment-method-specific value types.
package paymentmethods

// MobileMoneyNetwork identifies a supported mobile money network.
type MobileMoneyNetwork string

const (
	MobileMoneyNetworkAirtel   MobileMoneyNetwork = "airtel"
	MobileMoneyNetworkMTN      MobileMoneyNetwork = "mtn"
	MobileMoneyNetworkTelecel  MobileMoneyNetwork = "telecel"
	MobileMoneyNetworkVodafone MobileMoneyNetwork = "vodafone"
)
