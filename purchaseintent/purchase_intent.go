// Package purchaseintent provides purchaseintent resources and operations.
package purchaseintent

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/price"
	"github.com/zebodotdev/inttegro-sdk-go/v5/product"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusExpired  Status = "expired"
	StatusInactive Status = "inactive"
	StatusUsed     Status = "used"
)

type ActivityType string

const (
	ActivityTypeExpiredViewed  ActivityType = "expired_viewed"
	ActivityTypeOrderCreated   ActivityType = "order_created"
	ActivityTypePaymentFailed  ActivityType = "payment_failed"
	ActivityTypePaymentStarted ActivityType = "payment_started"
	ActivityTypeViewed         ActivityType = "viewed"
)

// PurchaseIntentsService manages Buy link purchase intents.
type Service struct {
	client transport.Client
}

type Quantity struct {
	Min int `json:"min"`
	Max int `json:"max,omitempty"`
}

type ProductSelector struct {
	ID           string `json:"id"`
	VariantSetID string `json:"variant_set_id,omitempty"`
}

type PriceSelector struct {
	ID         string               `json:"id,omitempty"`
	Nominal    *price.InlineParams  `json:"nominal,omitempty"`
	Original   *OriginalPriceParams `json:"original,omitempty"`
	OriginalID string               `json:"original_id,omitempty"`
}

type OriginalPriceParams struct {
	ID      string              `json:"id,omitempty"`
	Nominal *price.InlineParams `json:"nominal,omitempty"`
}

type OriginalPrice struct {
	Active  bool          `json:"active"`
	ID      string        `json:"id,omitempty"`
	Label   string        `json:"label,omitempty"`
	Nominal *money.Amount `json:"nominal,omitempty"`
}

type Price struct {
	Active   bool           `json:"active"`
	ID       string         `json:"id,omitempty"`
	Label    string         `json:"label,omitempty"`
	Nominal  *money.Amount  `json:"nominal,omitempty"`
	Original *OriginalPrice `json:"original,omitempty"`
}

type Usage struct {
	SingleUse *bool `json:"single_use,omitempty"`
	MultiUse  *bool `json:"multi_use,omitempty"`
}

type CreateParams struct {
	Product   *ProductSelector `json:"product,omitempty"`
	ProductID string           `json:"product_id,omitempty"`
	Price     *PriceSelector   `json:"price,omitempty"`
	PriceID   string           `json:"price_id,omitempty"`
	Quantity  Quantity         `json:"quantity"`
	Usage     *Usage           `json:"usage,omitempty"`
	ExpiresAt string           `json:"expires_at,omitempty"`
}

type UpdateParams struct {
	ID         string    `json:"id"`
	Quantity   *Quantity `json:"quantity,omitempty"`
	ExpiresAt  any       `json:"expires_at,omitempty"`
	Reactivate *bool     `json:"reactivate,omitempty"`
}

type PageParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type ActivityAttribution struct {
	LandingURL   string `json:"landing_url,omitempty"`
	Referrer     string `json:"referrer,omitempty"`
	ReferrerHost string `json:"referrer_host,omitempty"`
	Source       string `json:"source,omitempty"`
	Medium       string `json:"medium,omitempty"`
	Campaign     string `json:"campaign,omitempty"`
	Term         string `json:"term,omitempty"`
	Content      string `json:"content,omitempty"`
	Channel      string `json:"channel,omitempty"`
}

type ActivityVisitor struct {
	SessionID string `json:"session_id,omitempty"`
	VisitorID string `json:"visitor_id,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	Device    string `json:"device,omitempty"`
	Browser   string `json:"browser,omitempty"`
	OS        string `json:"os,omitempty"`
	Country   string `json:"country,omitempty"`
	Region    string `json:"region,omitempty"`
	City      string `json:"city,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}

type Activity struct {
	ID               string               `json:"id,omitempty"`
	PurchaseIntentID string               `json:"purchase_intent_id,omitempty"`
	Type             ActivityType         `json:"type,omitempty"`
	Source           string               `json:"source,omitempty"`
	Attribution      *ActivityAttribution `json:"attribution,omitempty"`
	Visitor          *ActivityVisitor     `json:"visitor,omitempty"`
	ProductID        string               `json:"product_id,omitempty"`
	VariantProductID string               `json:"variant_product_id,omitempty"`
	Quantity         int                  `json:"quantity,omitempty"`
	Amount           *money.Amount        `json:"amount,omitempty"`
	OrderID          string               `json:"order_id,omitempty"`
	PaymentID        string               `json:"payment_id,omitempty"`
	ErrorCode        string               `json:"error_code,omitempty"`
	CreatedAt        string               `json:"created_at,omitempty"`
}

type ActivityLog struct {
	Recent []Activity `json:"recent,omitempty"`
}

type Resource struct {
	ID                 string            `json:"id"`
	ProductID          string            `json:"product_id"`
	PriceID            string            `json:"price_id"`
	Quantity           Quantity          `json:"quantity"`
	AdjustableQuantity bool              `json:"adjustable_quantity"`
	AllowVariants      bool              `json:"allow_variants"`
	Status             Status            `json:"status"`
	CreatedAt          string            `json:"created_at"`
	UpdatedAt          string            `json:"updated_at,omitempty"`
	Activity           *ActivityLog      `json:"activity,omitempty"`
	Product            *product.Resource `json:"product,omitempty"`
	Price              *Price            `json:"price,omitempty"`
}

type Page struct {
	Number          int        `json:"number,omitempty"`
	Size            int        `json:"size,omitempty"`
	PurchaseIntents []Resource `json:"purchase_intents,omitempty"`
}

// Create creates a Buy link purchase intent.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Update modifies mutable Buy link purchase intent fields.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Cancel cancels a Buy link purchase intent.
func (s *Service) Cancel(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/cancel", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Lookup retrieves a Buy link purchase intent by ID.
func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Page lists Buy link purchase intents.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
