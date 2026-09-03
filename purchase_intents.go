package inttegro

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v4/money"
)

// PurchaseIntentsService manages Buy link purchase intents.
type PurchaseIntentsService struct {
	client *Client
}

type PurchaseIntentQuantity struct {
	Min int `json:"min"`
	Max int `json:"max,omitempty"`
}

type PurchaseIntentProductSelector struct {
	ID           string `json:"id"`
	VariantSetID string `json:"variant_set_id,omitempty"`
}

type PurchaseIntentPriceSelector struct {
	ID         string                             `json:"id,omitempty"`
	Nominal    *PriceParams                       `json:"nominal,omitempty"`
	Original   *PurchaseIntentOriginalPriceParams `json:"original,omitempty"`
	OriginalID string                             `json:"original_id,omitempty"`
}

type PurchaseIntentOriginalPriceParams struct {
	ID      string       `json:"id,omitempty"`
	Nominal *PriceParams `json:"nominal,omitempty"`
}

type PurchaseIntentOriginalPrice struct {
	Active  bool          `json:"active"`
	ID      string        `json:"id,omitempty"`
	Label   string        `json:"label,omitempty"`
	Nominal *money.Amount `json:"nominal,omitempty"`
}

type PurchaseIntentPrice struct {
	Active   bool                         `json:"active"`
	ID       string                       `json:"id,omitempty"`
	Label    string                       `json:"label,omitempty"`
	Nominal  *money.Amount                `json:"nominal,omitempty"`
	Original *PurchaseIntentOriginalPrice `json:"original,omitempty"`
}

type PurchaseIntentUsage struct {
	SingleUse *bool `json:"single_use,omitempty"`
	MultiUse  *bool `json:"multi_use,omitempty"`
}

type CreatePurchaseIntentParams struct {
	Product   *PurchaseIntentProductSelector `json:"product,omitempty"`
	ProductID string                         `json:"product_id,omitempty"`
	Price     *PurchaseIntentPriceSelector   `json:"price,omitempty"`
	PriceID   string                         `json:"price_id,omitempty"`
	Quantity  PurchaseIntentQuantity         `json:"quantity"`
	Usage     *PurchaseIntentUsage           `json:"usage,omitempty"`
	ExpiresAt string                         `json:"expires_at,omitempty"`
}

type UpdatePurchaseIntentParams struct {
	ID         string                  `json:"id"`
	Quantity   *PurchaseIntentQuantity `json:"quantity,omitempty"`
	ExpiresAt  any                     `json:"expires_at,omitempty"`
	Reactivate *bool                   `json:"reactivate,omitempty"`
}

type PagePurchaseIntentsParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type PurchaseIntentActivityAttribution struct {
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

type PurchaseIntentActivityVisitor struct {
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

type PurchaseIntentActivity struct {
	ID               string                             `json:"id,omitempty"`
	PurchaseIntentID string                             `json:"purchase_intent_id,omitempty"`
	Type             PurchaseIntentActivityType         `json:"type,omitempty"`
	Source           string                             `json:"source,omitempty"`
	Attribution      *PurchaseIntentActivityAttribution `json:"attribution,omitempty"`
	Visitor          *PurchaseIntentActivityVisitor     `json:"visitor,omitempty"`
	ProductID        string                             `json:"product_id,omitempty"`
	VariantProductID string                             `json:"variant_product_id,omitempty"`
	Quantity         int                                `json:"quantity,omitempty"`
	Amount           *money.Amount                      `json:"amount,omitempty"`
	OrderID          string                             `json:"order_id,omitempty"`
	PaymentID        string                             `json:"payment_id,omitempty"`
	ErrorCode        string                             `json:"error_code,omitempty"`
	CreatedAt        string                             `json:"created_at,omitempty"`
}

type PurchaseIntentActivityLog struct {
	Recent []PurchaseIntentActivity `json:"recent,omitempty"`
}

type PurchaseIntent struct {
	ID                 string                     `json:"id"`
	ApplicationID      string                     `json:"application_id"`
	ProductID          string                     `json:"product_id"`
	PriceID            string                     `json:"price_id"`
	Quantity           PurchaseIntentQuantity     `json:"quantity"`
	AdjustableQuantity bool                       `json:"adjustable_quantity"`
	AllowVariants      bool                       `json:"allow_variants"`
	Status             PurchaseIntentStatus       `json:"status"`
	CreatedAt          string                     `json:"created_at"`
	UpdatedAt          string                     `json:"updated_at,omitempty"`
	Activity           *PurchaseIntentActivityLog `json:"activity,omitempty"`
	Product            *Product                   `json:"product,omitempty"`
	Price              *PurchaseIntentPrice       `json:"price,omitempty"`
}

type PurchaseIntentsPage struct {
	Number          int              `json:"number,omitempty"`
	Size            int              `json:"size,omitempty"`
	PurchaseIntents []PurchaseIntent `json:"purchase_intents,omitempty"`
}

// Create creates a Buy link purchase intent.
func (s *PurchaseIntentsService) Create(ctx context.Context, params CreatePurchaseIntentParams) (*PurchaseIntent, error) {
	var resp struct {
		PurchaseIntent PurchaseIntent `json:"purchase_intent"`
	}
	if err := s.client.do(ctx, "POST", "/purchase_intents/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Update modifies mutable Buy link purchase intent fields.
func (s *PurchaseIntentsService) Update(ctx context.Context, params UpdatePurchaseIntentParams) (*PurchaseIntent, error) {
	var resp struct {
		PurchaseIntent PurchaseIntent `json:"purchase_intent"`
	}
	if err := s.client.do(ctx, "POST", "/purchase_intents/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Cancel cancels a Buy link purchase intent.
func (s *PurchaseIntentsService) Cancel(ctx context.Context, id string) (*PurchaseIntent, error) {
	var resp struct {
		PurchaseIntent PurchaseIntent `json:"purchase_intent"`
	}
	if err := s.client.do(ctx, "POST", "/purchase_intents/cancel", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Lookup retrieves a Buy link purchase intent by ID.
func (s *PurchaseIntentsService) Lookup(ctx context.Context, id string) (*PurchaseIntent, error) {
	var resp struct {
		PurchaseIntent PurchaseIntent `json:"purchase_intent"`
	}
	if err := s.client.do(ctx, "POST", "/purchase_intents/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Page lists Buy link purchase intents.
func (s *PurchaseIntentsService) Page(ctx context.Context, params PagePurchaseIntentsParams) (*PurchaseIntentsPage, error) {
	var resp struct {
		Page PurchaseIntentsPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/purchase_intents/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}
