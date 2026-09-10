package payment

import "time"

type NextAction struct {
	Type                NextActionType       `json:"type"`
	ConfirmPayment      *ConfirmPayment      `json:"confirm_payment,omitempty"`
	Redirect            *Redirect            `json:"redirect,omitempty"`
	Authorize           *Authorize           `json:"authorize,omitempty"`
	RequestConfirmation *RequestConfirmation `json:"request_confirmation,omitempty"`
}

type ConfirmPayment struct {
	ExpiresAt time.Time            `json:"expires_at"`
	Scheme    string               `json:"scheme"`
	Request   *ConfirmationRequest `json:"request,omitempty"`
	Attempt   *ConfirmationAttempt `json:"attempt,omitempty"`
	Confirmed bool                 `json:"confirmed"`
	Status    string               `json:"status"`
}

type ConfirmationRequest struct {
	ID        string              `json:"id"`
	Recipient string              `json:"recipient"`
	SentVia   ConfirmationChannel `json:"sent_via"`
	TokenSize int                 `json:"token_size"`
	SenderID  string              `json:"sender_id"`
	Status    string              `json:"status,omitempty"`
}

type ConfirmationAttempt struct {
	Status     string     `json:"status"`
	Confirmed  bool       `json:"confirmed"`
	Reason     string     `json:"reason"`
	ExecutedAt *time.Time `json:"executed_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Redirect struct {
	ValidUntil  time.Time      `json:"valid_until"`
	LatestVisit *RedirectVisit `json:"latest_visit,omitempty"`
	RedirectURL string         `json:"redirect_url"`
}

type RedirectVisit struct {
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address"`
	At        time.Time `json:"at"`
}

type Authorize struct {
	Beneficiary string    `json:"beneficiary"`
	ExpiresAt   time.Time `json:"expires_at"`
	Scheme      string    `json:"scheme"`
}

type RequestConfirmation struct {
	LastRequest *ConfirmationRequest `json:"last_request,omitempty"`
	After       *time.Time           `json:"after,omitempty"`
}
