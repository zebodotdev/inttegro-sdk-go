// Package checkout provides hosted-checkout settings and lifecycle values.
package checkout

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Settings      = inttegro.CheckoutSettings
	OrderStatus   = inttegro.CheckoutOrderStatus
	PaymentStatus = inttegro.CheckoutPaymentStatus
)

const (
	OrderStatusPreparing       = inttegro.CheckoutOrderStatusPreparing
	OrderStatusRequiresPayment = inttegro.CheckoutOrderStatusRequiresPayment
	OrderStatusCompleted       = inttegro.CheckoutOrderStatusCompleted
	OrderStatusCanceled        = inttegro.CheckoutOrderStatusCanceled
	OrderStatusExpired         = inttegro.CheckoutOrderStatusExpired

	PaymentStatusRequiresAction = inttegro.CheckoutPaymentStatusRequiresAction
	PaymentStatusProcessing     = inttegro.CheckoutPaymentStatusProcessing
	PaymentStatusSucceeded      = inttegro.CheckoutPaymentStatusSucceeded
	PaymentStatusFailed         = inttegro.CheckoutPaymentStatusFailed
	PaymentStatusCancelled      = inttegro.CheckoutPaymentStatusCancelled
)
