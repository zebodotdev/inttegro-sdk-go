package product

type Type string

const (
	TypePhysical Type = "physical"
	TypeDigital  Type = "digital"
	TypeService  Type = "service"
	TypeVoucher  Type = "voucher"
	TypeCustom   Type = "custom"
	TypeCause    Type = "cause"
)

type ShipmentType string

const (
	ShipmentTypeDelivery ShipmentType = "delivery"
	ShipmentTypeDownload ShipmentType = "download"
	ShipmentTypeRender   ShipmentType = "render"
	ShipmentTypeService  ShipmentType = "service"
	ShipmentTypeStream   ShipmentType = "stream"
)

type ShipmentInputType string

const (
	ShipmentInputTypeDelivery ShipmentInputType = "delivery"
	ShipmentInputTypeDownload ShipmentInputType = "download"
	ShipmentInputTypeRender   ShipmentInputType = "render"
	ShipmentInputTypeStream   ShipmentInputType = "stream"
)
