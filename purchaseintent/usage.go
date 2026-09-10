package purchaseintent

import "time"

type Usage struct {
	SingleUse *bool       `json:"single_use,omitempty"`
	MultiUse  *bool       `json:"multi_use,omitempty"`
	Order     *UsageOrder `json:"order,omitempty"`
}

type UsageOrder struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
}
