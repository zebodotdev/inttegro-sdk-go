package purchaseintent

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/product"
)

type Product struct {
	ID           string                 `json:"id"`
	About        string                 `json:"about,omitempty"`
	Active       bool                   `json:"active"`
	ArchivedAt   *time.Time             `json:"archived_at,omitempty"`
	Attributes   []ProductAttribute     `json:"attributes,omitempty"`
	Category     string                 `json:"category,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	CustomData   map[string]string      `json:"custom_data,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Dimensions   *ProductDimensions     `json:"dimensions,omitempty"`
	Media        *ProductMedia          `json:"media,omitempty"`
	Name         string                 `json:"name"`
	Prices       []product.PriceSummary `json:"prices,omitempty"`
	PublishedAt  *time.Time             `json:"published_at,omitempty"`
	Reference    string                 `json:"reference,omitempty"`
	Shipment     *ProductShipment       `json:"shipment,omitempty"`
	TaxCode      string                 `json:"tax_code,omitempty"`
	Type         product.Type           `json:"type"`
	UnitDim      string                 `json:"unit_dim,omitempty"`
	UpdatedAt    *time.Time             `json:"updated_at,omitempty"`
	VariantSetID string                 `json:"variant_set_id,omitempty"`
}

type ProductAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductDimensions struct {
	Physical *PhysicalDimensions `json:"physical,omitempty"`
	Digital  *DigitalDimensions  `json:"digital,omitempty"`
	Custom   *CustomDimensions   `json:"custom,omitempty"`
}

type PhysicalDimensions struct {
	WeightUnit string  `json:"weight_unit,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	Size       float64 `json:"size,omitempty"`
	VolumeUnit string  `json:"volume_unit,omitempty"`
	Volume     float64 `json:"volume,omitempty"`
	Length     float64 `json:"length,omitempty"`
	Height     float64 `json:"height,omitempty"`
	Width      float64 `json:"width,omitempty"`
}

type DigitalDimensions struct {
	Bytes    float64 `json:"bytes,omitempty"`
	SizeUnit string  `json:"size_unit,omitempty"`
	Size     float64 `json:"size,omitempty"`
}

type CustomDimensions struct {
	SizeUnit string            `json:"size_unit,omitempty"`
	Size     float64           `json:"size,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

type ProductMedia struct {
	HeroImage   string   `json:"hero_image,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	WebPageURL  string   `json:"web_page_url,omitempty"`
	BrandLogo   string   `json:"brand_logo,omitempty"`
	Infographic string   `json:"infographic,omitempty"`
	PromoVideo  string   `json:"promo_video,omitempty"`
	DemoVideo   string   `json:"demo_video,omitempty"`
	Gallery     []string `json:"gallery,omitempty"`
	Downloads   []string `json:"downloads,omitempty"`
}

type ProductShipment struct {
	Type     product.ShipmentType `json:"type"`
	Delivery *ProductDelivery     `json:"delivery,omitempty"`
	Download *ProductDownload     `json:"download,omitempty"`
	Render   *ProductRender       `json:"render,omitempty"`
	Service  *ProductService      `json:"service,omitempty"`
	Stream   *ProductStream       `json:"stream,omitempty"`
}

type ProductDelivery struct{}
type ProductDownload struct{}
type ProductRender struct{}
type ProductService struct{}
type ProductStream struct{}
