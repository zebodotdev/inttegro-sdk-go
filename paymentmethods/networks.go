// Package paymentmethods defines payment-method-specific value types.
//
// Deprecated: use the singular paymentmethod package for new code.
package paymentmethods

// MobileMoneyNetwork identifies a supported mobile money network.
type MobileMoneyNetwork string

const (
	MobileMoneyNetworkAirtel   MobileMoneyNetwork = "airtel"
	MobileMoneyNetworkMTN      MobileMoneyNetwork = "mtn"
	MobileMoneyNetworkTelecel  MobileMoneyNetwork = "telecel"
	MobileMoneyNetworkVodafone MobileMoneyNetwork = "vodafone"
)
