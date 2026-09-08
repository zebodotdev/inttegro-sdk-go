// Package price provides catalog and inline price resources and operations.
//
// Its names are exact aliases of the v4 root types, so existing and
// package-scoped code interoperate without conversions.
package price

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service      = inttegro.PricesService
	InlineParams = inttegro.PriceParams
	Inline       = inttegro.Price
	CreateParams = inttegro.CatalogPriceParams
	LookupParams = inttegro.LookupPriceParams
	UpdateParams = inttegro.UpdatePriceParams
	ActionParams = inttegro.PriceActionParams
	PageParams   = inttegro.PricePageParams
	Resource     = inttegro.CatalogPrice
	Page         = inttegro.PricesPage
)
