// Package purchaseintent provides Buy link purchase-intent resources and operations.
//
// Its names are exact aliases of the v4 root types, so existing and
// package-scoped code interoperate without conversions. The resource package is
// the preferred public API for new code.
package purchaseintent

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service             = inttegro.PurchaseIntentsService
	Status              = inttegro.PurchaseIntentStatus
	ActivityType        = inttegro.PurchaseIntentActivityType
	Quantity            = inttegro.PurchaseIntentQuantity
	ProductSelector     = inttegro.PurchaseIntentProductSelector
	PriceSelector       = inttegro.PurchaseIntentPriceSelector
	OriginalPriceParams = inttegro.PurchaseIntentOriginalPriceParams
	OriginalPrice       = inttegro.PurchaseIntentOriginalPrice
	Price               = inttegro.PurchaseIntentPrice
	Usage               = inttegro.PurchaseIntentUsage
	CreateParams        = inttegro.CreatePurchaseIntentParams
	UpdateParams        = inttegro.UpdatePurchaseIntentParams
	PageParams          = inttegro.PagePurchaseIntentsParams
	ActivityAttribution = inttegro.PurchaseIntentActivityAttribution
	ActivityVisitor     = inttegro.PurchaseIntentActivityVisitor
	Activity            = inttegro.PurchaseIntentActivity
	ActivityLog         = inttegro.PurchaseIntentActivityLog
	Resource            = inttegro.PurchaseIntent
	Page                = inttegro.PurchaseIntentsPage
)

const (
	StatusActive   = inttegro.PurchaseIntentStatusActive
	StatusExpired  = inttegro.PurchaseIntentStatusExpired
	StatusInactive = inttegro.PurchaseIntentStatusInactive
	StatusUsed     = inttegro.PurchaseIntentStatusUsed

	ActivityTypeExpiredViewed  = inttegro.PurchaseIntentActivityTypeExpiredViewed
	ActivityTypeOrderCreated   = inttegro.PurchaseIntentActivityTypeOrderCreated
	ActivityTypePaymentFailed  = inttegro.PurchaseIntentActivityTypePaymentFailed
	ActivityTypePaymentStarted = inttegro.PurchaseIntentActivityTypePaymentStarted
	ActivityTypeViewed         = inttegro.PurchaseIntentActivityTypeViewed
)
