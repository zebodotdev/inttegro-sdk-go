package refund

type Reason string

const (
	ReasonRequestedByCustomer Reason = "requested_by_customer"
	ReasonDuplicate           Reason = "duplicate"
	ReasonFraudulent          Reason = "fraudulent"
	ReasonOrderCanceled       Reason = "order_canceled"
	ReasonItemReturned        Reason = "item_returned"
	ReasonItemDamaged         Reason = "item_damaged"
	ReasonItemNotReceived     Reason = "item_not_received"
	ReasonItemNotAsDescribed  Reason = "item_not_as_described"
	ReasonCustom              Reason = "custom"
)

type Status string

const (
	StatusCanceled   Status = "canceled"
	StatusFailed     Status = "failed"
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusSucceeded  Status = "succeeded"
)
