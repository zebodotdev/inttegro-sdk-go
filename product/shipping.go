package product

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
