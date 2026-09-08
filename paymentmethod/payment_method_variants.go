// Package paymentmethod provides paymentmethod values.
package paymentmethod

// MobileMoneyNetwork identifies a supported mobile money network.
type MobileMoneyNetwork string

const (
	MobileMoneyNetworkAirtel   MobileMoneyNetwork = "airtel"
	MobileMoneyNetworkMTN      MobileMoneyNetwork = "mtn"
	MobileMoneyNetworkTelecel  MobileMoneyNetwork = "telecel"
	MobileMoneyNetworkVodafone MobileMoneyNetwork = "vodafone"
)
