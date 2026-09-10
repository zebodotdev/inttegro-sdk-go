package purchaseintent

import (
	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
)

type OriginalPrice struct {
	Active  bool         `json:"active"`
	ID      string       `json:"id,omitempty"`
	Label   string       `json:"label,omitempty"`
	Nominal money.Amount `json:"nominal"`
}

type Price struct {
	Active   bool           `json:"active"`
	ID       string         `json:"id,omitempty"`
	Label    string         `json:"label,omitempty"`
	Nominal  money.Amount   `json:"nominal"`
	Original *OriginalPrice `json:"original,omitempty"`
}
