package product

import (
	"bytes"
	"encoding/json"
)

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
