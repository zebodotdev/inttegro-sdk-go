package product

// Shipment describes the product's fulfillment mechanism.
type Shipment struct {
	Type     ShipmentType    `json:"type"`
	Delivery *Delivery       `json:"delivery,omitempty"`
	Download *Download       `json:"download,omitempty"`
	Render   *Render         `json:"render,omitempty"`
	Service  *ServiceDetails `json:"service,omitempty"`
	Stream   *Stream         `json:"stream,omitempty"`
}

// The fulfillment variants are explicit marker objects. The API currently
// returns no fields inside them.
type Delivery struct{}
type Download struct{}
type Render struct{}
type ServiceDetails struct{}
type Stream struct{}

// ShipmentInput describes fulfillment accepted by create and update requests.
type ShipmentInput struct {
	Type ShipmentInputType `json:"type"`
}
