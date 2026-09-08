package purchaseintent

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
)

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
