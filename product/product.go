// Package product provides product resources and operations.
package product

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
)

type Type string

const (
	TypePhysical Type = "physical"
)

const (
	TypeDigital Type = "digital"
)

const (
	TypeService Type = "service"
)

const (
	TypeVoucher Type = "voucher"
)

const (
	TypeCustom Type = "custom"
)

const (
	TypeCause Type = "cause"
)

type ShipmentType string

const (
	ShipmentTypeDelivery ShipmentType = "delivery"
)

const (
	ShipmentTypeDownload ShipmentType = "download"
)

const (
	ShipmentTypeRender ShipmentType = "render"
)

const (
	ShipmentTypeService ShipmentType = "service"
)

const (
	ShipmentTypeStream ShipmentType = "stream"
)

type ShipmentInputType string

const (
	ShipmentInputTypeDelivery ShipmentInputType = "delivery"
)

const (
	ShipmentInputTypeDownload ShipmentInputType = "download"
)

const (
	ShipmentInputTypeRender ShipmentInputType = "render"
)

const (
	ShipmentInputTypeStream ShipmentInputType = "stream"
)

// ProductsService manages catalog products.
type Service struct {
	client transport.Client
}

// Create creates a product.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// SetDefaultUnitPrice sets an existing product price as the product's default unit price.
func (s *Service) SetDefaultUnitPrice(ctx context.Context, params SetDefaultUnitPriceParams) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/set_default_unit_price", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Lookup retrieves a product by ID.
func (s *Service) Lookup(ctx context.Context, productID string) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/lookup", LookupParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Page retrieves a page of products.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/products/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update updates a product.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Publish publishes a product.
func (s *Service) Publish(ctx context.Context, productID string) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/publish", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Unpublish unpublishes a product.
func (s *Service) Unpublish(ctx context.Context, productID string) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/unpublish", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Archive archives a product.
func (s *Service) Archive(ctx context.Context, productID string) (*Resource, error) {
	var resp struct {
		Product Resource `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/archive", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// ProductCategory describes a product category.
type Category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// UnmarshalJSON accepts both the canonical category string and the legacy
// expanded category object so existing callers keep their field accessors.
func (c *Category) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if bytes.Equal(trimmed, []byte("null")) {
		*c = Category{}
		return nil
	}
	if len(trimmed) > 0 && trimmed[0] == '"' {
		return json.Unmarshal(trimmed, &c.Name)
	}
	type categoryAlias Category
	return json.Unmarshal(trimmed, (*categoryAlias)(c))
}

// ProductDefaultUnitPrice represents a product's loaded default unit price.
type DefaultUnitPrice struct {
	ID         string        `json:"id,omitempty"`
	ProductID  string        `json:"product_id,omitempty"`
	Label      string        `json:"label,omitempty"`
	About      string        `json:"about,omitempty"`
	Nominal    *money.Amount `json:"nominal,omitempty"`
	CreatedAt  string        `json:"created_at,omitempty"`
	UpdatedAt  string        `json:"updated_at,omitempty"`
	ArchivedAt string        `json:"archived_at,omitempty"`
}

// ProductPriceSummary represents a product price listed with a product response.
type PriceSummary struct {
	ID      string        `json:"id,omitempty"`
	Label   string        `json:"label,omitempty"`
	Nominal *money.Amount `json:"nominal,omitempty"`
	Active  bool          `json:"active,omitempty"`
}

// ProductShipmentDimensions describes physical dimensions.
type ShipmentDimensions struct {
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Weight float64 `json:"weight,omitempty"`
}

// ProductShipment describes fulfillment details.
type Shipment struct {
	Type       ShipmentType        `json:"type,omitempty"`
	Carrier    string              `json:"carrier,omitempty"`
	Dimensions *ShipmentDimensions `json:"dimensions,omitempty"`
}

// ProductShipmentInput describes fulfillment accepted by create and update requests.
type ShipmentInput struct {
	Type       ShipmentInputType   `json:"type,omitempty"`
	Carrier    string              `json:"carrier,omitempty"`
	Dimensions *ShipmentDimensions `json:"dimensions,omitempty"`
}

// ProductMediaItem represents product media.
type MediaItem struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

// CreateProductParams creates a catalog product.
type CreateParams struct {
	Type        Type              `json:"type"`
	Reference   string            `json:"reference,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	About       string            `json:"about,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	Category    *Category         `json:"category,omitempty"`
	Shipment    *ShipmentInput    `json:"shipment,omitempty"`
	Media       []MediaItem       `json:"media,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
}

// LookupProductParams looks up a product by ID.
type LookupParams struct {
	ProductID string `json:"product_id"`
}

// UpdateProductParams updates a product.
type UpdateParams struct {
	ProductID   string            `json:"product_id"`
	Reference   string            `json:"reference,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	About       string            `json:"about,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	Category    *Category         `json:"category,omitempty"`
	Shipment    *ShipmentInput    `json:"shipment,omitempty"`
	Media       []MediaItem       `json:"media,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
}

// ProductActionParams performs an action on a product.
type ActionParams struct {
	ProductID string `json:"product_id"`
}

// SetDefaultUnitPriceParams changes the price used as a product's default.
type SetDefaultUnitPriceParams struct {
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
}

// PageProductsParams pages through products.
type PageParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

// Product represents a catalog product.
type Resource struct {
	ID               string            `json:"id,omitempty"`
	ApplicationID    string            `json:"application_id,omitempty"`
	Type             Type              `json:"type,omitempty"`
	Reference        string            `json:"reference,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	About            string            `json:"about,omitempty"`
	TaxCode          string            `json:"tax_code,omitempty"`
	Category         *Category         `json:"category,omitempty"`
	DefaultUnitPrice *DefaultUnitPrice `json:"default_unit_price,omitempty"`
	Prices           []PriceSummary    `json:"prices,omitempty"`
	Shipment         *Shipment         `json:"shipment,omitempty"`
	Media            []MediaItem       `json:"media,omitempty"`
	Attributes       map[string]string `json:"attributes,omitempty"`
	CustomData       map[string]string `json:"custom_data,omitempty"`
	Active           bool              `json:"active,omitempty"`
	Archived         bool              `json:"archived,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	UpdatedAt        string            `json:"updated_at,omitempty"`
	ArchivedAt       string            `json:"archived_at,omitempty"`
}

// UnmarshalJSON accepts the current Product response while preserving the
// legacy exported category, media, and attribute field types used by v4
// callers. A future major version can expose the canonical shapes directly.
func (p *Resource) UnmarshalJSON(data []byte) error {
	type productAlias Resource
	decoded := struct {
		Media      json.RawMessage `json:"media"`
		Attributes json.RawMessage `json:"attributes"`
		*productAlias
	}{productAlias: (*productAlias)(p)}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	p.Media = nil
	p.Attributes = nil
	if err := decodeProductMedia(decoded.Media, &p.Media); err != nil {
		return err
	}
	return decodeProductAttributes(decoded.Attributes, &p.Attributes)
}

func decodeProductMedia(data json.RawMessage, target *[]MediaItem) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	if trimmed[0] == '[' {
		return json.Unmarshal(trimmed, target)
	}
	var media struct {
		HeroImage   string   `json:"hero_image"`
		Thumbnail   string   `json:"thumbnail"`
		WebPageURL  string   `json:"web_page_url"`
		BrandLogo   string   `json:"brand_logo"`
		Infographic string   `json:"infographic"`
		PromoVideo  string   `json:"promo_video"`
		DemoVideo   string   `json:"demo_video"`
		Gallery     []string `json:"gallery"`
		Downloads   []string `json:"downloads"`
	}
	if err := json.Unmarshal(trimmed, &media); err != nil {
		return err
	}
	appendItem := func(kind, value string) {
		if value != "" {
			*target = append(*target, MediaItem{Type: kind, URL: value})
		}
	}
	appendItem("hero_image", media.HeroImage)
	appendItem("thumbnail", media.Thumbnail)
	appendItem("web_page_url", media.WebPageURL)
	appendItem("brand_logo", media.BrandLogo)
	appendItem("infographic", media.Infographic)
	appendItem("promo_video", media.PromoVideo)
	appendItem("demo_video", media.DemoVideo)
	for _, value := range media.Gallery {
		appendItem("gallery", value)
	}
	for _, value := range media.Downloads {
		appendItem("download", value)
	}
	return nil
}

func decodeProductAttributes(data json.RawMessage, target *map[string]string) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	if trimmed[0] == '{' {
		return json.Unmarshal(trimmed, target)
	}
	var attributes []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(trimmed, &attributes); err != nil {
		return err
	}
	values := make(map[string]string, len(attributes))
	for _, attribute := range attributes {
		if attribute.Name != "" {
			values[attribute.Name] = attribute.Value
		}
	}
	*target = values
	return nil
}

// ProductsPage holds a page of products.
type Page struct {
	Number   int        `json:"number,omitempty"`
	Size     int        `json:"size,omitempty"`
	Products []Resource `json:"products,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
