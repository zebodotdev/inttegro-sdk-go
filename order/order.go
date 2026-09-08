// Package order provides order, checkout, payment, and invoice resources and operations.
package order

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service                    = inttegro.OrdersService
	LineItemType               = inttegro.LineItemType
	DocumentKind               = inttegro.OrderDocumentKind
	DeliveryChannel            = inttegro.DeliveryChannel
	CheckoutStatus             = inttegro.CheckoutOrderStatus
	Status                     = inttegro.OrderStatus
	PaymentStatus              = inttegro.PaymentStatus
	PaymentAttemptStatus       = inttegro.PaymentAttemptStatus
	PaymentNextActionType      = inttegro.PaymentNextActionType
	PaymentConfirmationChannel = inttegro.PaymentConfirmationChannel
	CheckoutPaymentStatus      = inttegro.CheckoutPaymentStatus
	PaymentResultStatus        = inttegro.PaymentResultStatus
	CreatedFromResourceType    = inttegro.OrderCreatedFromResourceType
	Address                    = inttegro.Address
	BillingDetails             = inttegro.BillingDetails
	Shipping                   = inttegro.Shipping
	ProductLineItemParams      = inttegro.ProductLineItemParams
	ProductLineItem            = inttegro.ProductLineItem
	FeeLineItemParams          = inttegro.FeeLineItemParams
	FeeLineItem                = inttegro.FeeLineItem
	ShippingLineItemParams     = inttegro.ShippingLineItemParams
	ShippingLineItem           = inttegro.ShippingLineItem
	LineItemParams             = inttegro.OrderLineItemParams
	LineItem                   = inttegro.OrderLineItem
	CheckoutSettings           = inttegro.CheckoutSettings
	PayoutSettings             = inttegro.OrderPayoutSettings
	PayoutDestination          = inttegro.OrderPayoutDestination
	PayoutFinancialAccount     = inttegro.OrderPayoutFinancialAccount
	CreateParams               = inttegro.OrderCreateParams
	LookupParams               = inttegro.OrderLookupParams
	PayParams                  = inttegro.OrderPayParams
	ConfirmParams              = inttegro.OrderConfirmParams
	RequestConfirmationParams  = inttegro.OrderRequestConfirmationParams
	FinalizeParams             = inttegro.OrderFinalizeParams
	SendInvoiceParams          = inttegro.OrderSendInvoiceParams
	SendReceiptParams          = inttegro.OrderSendReceiptParams
	CompleteParams             = inttegro.OrderCompleteParams
	CancelParams               = inttegro.OrderCancelParams
	PageParams                 = inttegro.OrderPageParams
	DocumentDeliveryResult     = inttegro.OrderDocumentDeliveryResult
	DocumentDelivery           = inttegro.OrderDocumentDelivery
	DocumentDeliveryAttempt    = inttegro.OrderDocumentDeliveryAttempt
	PaymentMethod              = inttegro.PaymentMethod
	PaymentAttempt             = inttegro.PaymentAttempt
	PaymentNextAction          = inttegro.PaymentNextAction
	Payment                    = inttegro.Payment
	InvoiceFormat              = inttegro.InvoiceFormat
	Invoice                    = inttegro.Invoice
	LineItemGroup              = inttegro.LineItemGroup
	Resource                   = inttegro.Order
)

const (
	LineItemTypeProduct  = inttegro.LineItemTypeProduct
	LineItemTypeFee      = inttegro.LineItemTypeFee
	LineItemTypeShipping = inttegro.LineItemTypeShipping

	DocumentKindInvoice = inttegro.OrderDocumentKindInvoice
	DocumentKindReceipt = inttegro.OrderDocumentKindReceipt

	DeliveryChannelEmail = inttegro.DeliveryChannelEmail
	DeliveryChannelSMS   = inttegro.DeliveryChannelSMS

	CheckoutStatusPreparing       = inttegro.CheckoutOrderStatusPreparing
	CheckoutStatusRequiresPayment = inttegro.CheckoutOrderStatusRequiresPayment
	CheckoutStatusCompleted       = inttegro.CheckoutOrderStatusCompleted
	CheckoutStatusCanceled        = inttegro.CheckoutOrderStatusCanceled
	CheckoutStatusExpired         = inttegro.CheckoutOrderStatusExpired

	StatusPreparing       = inttegro.OrderStatusPreparing
	StatusRequiresPayment = inttegro.OrderStatusRequiresPayment
	StatusPaid            = inttegro.OrderStatusPaid
	StatusCompleted       = inttegro.OrderStatusCompleted
	StatusCanceled        = inttegro.OrderStatusCanceled
	StatusExpired         = inttegro.OrderStatusExpired
	StatusUnknown         = inttegro.OrderStatusUnknown

	PaymentStatusInitiated      = inttegro.PaymentStatusInitiated
	PaymentStatusRequiresAction = inttegro.PaymentStatusRequiresAction
	PaymentStatusOverdue        = inttegro.PaymentStatusOverdue
	PaymentStatusExecuted       = inttegro.PaymentStatusExecuted
	PaymentStatusPaid           = inttegro.PaymentStatusPaid
	PaymentStatusCanceled       = inttegro.PaymentStatusCanceled
	PaymentStatusExpired        = inttegro.PaymentStatusExpired
	PaymentStatusFailed         = inttegro.PaymentStatusFailed
	PaymentStatusUnknown        = inttegro.PaymentStatusUnknown

	PaymentNextActionTypeConfirmPayment = inttegro.PaymentNextActionTypeConfirmPayment
	PaymentNextActionTypeExecute        = inttegro.PaymentNextActionTypeExecute
	PaymentNextActionTypeRedirect       = inttegro.PaymentNextActionTypeRedirect
	PaymentNextActionTypeAuthorize      = inttegro.PaymentNextActionTypeAuthorize
	PaymentNextActionTypeNone           = inttegro.PaymentNextActionTypeNone

	PaymentConfirmationChannelSMS   = inttegro.PaymentConfirmationChannelSMS
	PaymentConfirmationChannelEmail = inttegro.PaymentConfirmationChannelEmail
	PaymentConfirmationChannelPush  = inttegro.PaymentConfirmationChannelPush

	PaymentAttemptStatusInitiated = inttegro.PaymentAttemptStatusInitiated
	PaymentAttemptStatusExecuted  = inttegro.PaymentAttemptStatusExecuted
	PaymentAttemptStatusSucceeded = inttegro.PaymentAttemptStatusSucceeded
	PaymentAttemptStatusCanceled  = inttegro.PaymentAttemptStatusCanceled
	PaymentAttemptStatusExpired   = inttegro.PaymentAttemptStatusExpired
	PaymentAttemptStatusFailed    = inttegro.PaymentAttemptStatusFailed
	PaymentAttemptStatusUnknown   = inttegro.PaymentAttemptStatusUnknown

	CheckoutPaymentStatusRequiresAction = inttegro.CheckoutPaymentStatusRequiresAction
	CheckoutPaymentStatusProcessing     = inttegro.CheckoutPaymentStatusProcessing
	CheckoutPaymentStatusSucceeded      = inttegro.CheckoutPaymentStatusSucceeded
	CheckoutPaymentStatusFailed         = inttegro.CheckoutPaymentStatusFailed
	CheckoutPaymentStatusCancelled      = inttegro.CheckoutPaymentStatusCancelled

	PaymentResultStatusPending              = inttegro.PaymentResultStatusPending
	PaymentResultStatusRequiresConfirmation = inttegro.PaymentResultStatusRequiresConfirmation
	PaymentResultStatusProcessing           = inttegro.PaymentResultStatusProcessing
	PaymentResultStatusSucceeded            = inttegro.PaymentResultStatusSucceeded
	PaymentResultStatusFailed               = inttegro.PaymentResultStatusFailed

	CreatedFromResourceTypePurchaseIntent = inttegro.OrderCreatedFromResourceTypePurchaseIntent
)
