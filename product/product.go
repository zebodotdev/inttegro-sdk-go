// Package product provides catalog product resources and operations.
//
// Its names are exact aliases of the v4 root types, so existing and
// package-scoped code interoperate without conversions. The resource package is
// the preferred public API for new code.
package product

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service                   = inttegro.ProductsService
	Type                      = inttegro.ProductType
	ShipmentType              = inttegro.ProductShipmentType
	ShipmentInputType         = inttegro.ProductShipmentInputType
	Category                  = inttegro.ProductCategory
	DefaultUnitPrice          = inttegro.ProductDefaultUnitPrice
	PriceSummary              = inttegro.ProductPriceSummary
	ShipmentDimensions        = inttegro.ProductShipmentDimensions
	Shipment                  = inttegro.ProductShipment
	ShipmentInput             = inttegro.ProductShipmentInput
	MediaItem                 = inttegro.ProductMediaItem
	CreateParams              = inttegro.CreateProductParams
	LookupParams              = inttegro.LookupProductParams
	UpdateParams              = inttegro.UpdateProductParams
	ActionParams              = inttegro.ProductActionParams
	AddPriceParams            = inttegro.AddProductPriceParams
	SetDefaultUnitPriceParams = inttegro.SetDefaultUnitPriceParams
	PageParams                = inttegro.PageProductsParams
	Resource                  = inttegro.Product
	Page                      = inttegro.ProductsPage
)

const (
	TypePhysical = inttegro.ProductTypePhysical
	TypeDigital  = inttegro.ProductTypeDigital
	TypeService  = inttegro.ProductTypeService
	TypeVoucher  = inttegro.ProductTypeVoucher
	TypeCustom   = inttegro.ProductTypeCustom
	TypeCause    = inttegro.ProductTypeCause

	ShipmentTypeDelivery = inttegro.ProductShipmentTypeDelivery
	ShipmentTypeDownload = inttegro.ProductShipmentTypeDownload
	ShipmentTypeRender   = inttegro.ProductShipmentTypeRender
	ShipmentTypeService  = inttegro.ProductShipmentTypeService
	ShipmentTypeStream   = inttegro.ProductShipmentTypeStream

	ShipmentInputTypeDelivery = inttegro.ProductShipmentInputTypeDelivery
	ShipmentInputTypeDownload = inttegro.ProductShipmentInputTypeDownload
	ShipmentInputTypeRender   = inttegro.ProductShipmentInputTypeRender
	ShipmentInputTypeStream   = inttegro.ProductShipmentInputTypeStream
)
