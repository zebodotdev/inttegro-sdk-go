package order

import "github.com/zebodotdev/inttegro-sdk-go/v5/invoice"

// DocumentDeliveryResult contains the updated order and its delivery result.
type DocumentDeliveryResult struct {
	Order    Resource         `json:"order"`
	Delivery invoice.Delivery `json:"delivery"`
}
