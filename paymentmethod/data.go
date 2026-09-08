package paymentmethod

// MobileMoneyParams describes a mobile money wallet.
//
// Used when creating orders or tokenizing payment methods with inline
// mobile money data instead of referencing a saved payment method.
type MobileMoneyParams struct {
	// Network is the mobile money network code.
	// Examples: "mtn", "vodafone", "airteltigo", "airtel", "telecel".
	Network MobileMoneyNetwork `json:"network"`

	// AccountNumber is the mobile money account phone number.
	// Must include country code. Example: "+233244123456"
	// Used as the payment source and for sending OTP verification codes.
	AccountNumber string `json:"account_number"`
}

// PaymentMethodData represents inline payment method data for one-time use.
//
// Use this to charge a payment method without saving it for future use.
// For repeat customers, tokenize the payment method first using
// PaymentMethods.Tokenize, then reference it by ID.
//
// Example (mobile money):
//
//	paymentData := &paymentmethod.Data{
//	    Type: paymentmethod.TypeMobileMoney,
//	    MobileMoney: &paymentmethod.MobileMoneyParams{
//	        Network: "mtn",
//	        AccountNumber: "+233244123456",
//	    },
//	}
type Data struct {
	// Type specifies the payment method category.
	Type Type `json:"type"`

	// MobileMoney provides mobile money wallet details.
	// Required when Type is PaymentMethodTypeMobileMoney.
	MobileMoney *MobileMoneyParams `json:"mobile_money,omitempty"`
}
