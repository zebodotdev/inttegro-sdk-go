// Package invoice provides invoice resources and delivery values.
package invoice

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Format          = inttegro.InvoiceFormat
	Resource        = inttegro.Invoice
	DocumentKind    = inttegro.OrderDocumentKind
	DeliveryChannel = inttegro.DeliveryChannel
	DeliveryResult  = inttegro.OrderDocumentDeliveryResult
	Delivery        = inttegro.OrderDocumentDelivery
	DeliveryAttempt = inttegro.OrderDocumentDeliveryAttempt
)

const (
	DocumentKindInvoice = inttegro.OrderDocumentKindInvoice
	DocumentKindReceipt = inttegro.OrderDocumentKindReceipt

	DeliveryChannelEmail = inttegro.DeliveryChannelEmail
	DeliveryChannelSMS   = inttegro.DeliveryChannelSMS
)
