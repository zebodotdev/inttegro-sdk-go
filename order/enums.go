package order

type LineItemType string

const (
	LineItemTypeProduct  LineItemType = "product"
	LineItemTypeFee      LineItemType = "fee"
	LineItemTypeShipping LineItemType = "shipping"
	LineItemTypeDiscount LineItemType = "discount"
)

type Status string

const (
	StatusPreparing       Status = "preparing"
	StatusRequiresPayment Status = "requires_payment"
	StatusPaid            Status = "paid"
	StatusCompleted       Status = "completed"
	StatusCanceled        Status = "canceled"
	StatusExpired         Status = "expired"
	StatusUnknown         Status = "unknown"
)

type CreatedFromResourceType string

const CreatedFromResourceTypePurchaseIntent CreatedFromResourceType = "purchase_intent"
