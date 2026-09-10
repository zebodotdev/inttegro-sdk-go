package product

// Dimensions contains one dimension variant appropriate for the product.
type Dimensions struct {
	Physical *PhysicalDimensions `json:"physical,omitempty"`
	Digital  *DigitalDimensions  `json:"digital,omitempty"`
	Custom   *CustomDimensions   `json:"custom,omitempty"`
}

type PhysicalDimensions struct {
	WeightUnit string  `json:"weight_unit,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	VolumeUnit string  `json:"volume_unit,omitempty"`
	Volume     float64 `json:"volume,omitempty"`
	Length     float64 `json:"length,omitempty"`
	Height     float64 `json:"height,omitempty"`
	Width      float64 `json:"width,omitempty"`
}

type DigitalDimensions struct {
	Bytes    float64 `json:"bytes,omitempty"`
	SizeUnit string  `json:"size_unit,omitempty"`
	Size     float64 `json:"size,omitempty"`
}

type CustomDimensions struct {
	SizeUnit string            `json:"size_unit,omitempty"`
	Size     float64           `json:"size,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}
