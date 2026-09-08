package purchaseintent

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
)

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
