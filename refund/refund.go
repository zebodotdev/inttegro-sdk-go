// Package refund provides refund resources and operations.
package refund

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service        = inttegro.RefundsService
	Reason         = inttegro.RefundReason
	Status         = inttegro.RefundStatus
	CreateLineItem = inttegro.CreateRefundLineItem
	CreateParams   = inttegro.CreateRefundRequest
	CancelParams   = inttegro.CancelRefundRequest
	LookupParams   = inttegro.LookupRefundRequest
	PageParams     = inttegro.PageRefundsRequest
	LineItem       = inttegro.RefundLineItem
	Resource       = inttegro.Refund
	Page           = inttegro.RefundPage
)

const (
	ReasonRequestedByCustomer = inttegro.RefundReasonRequestedByCustomer
	ReasonDuplicate           = inttegro.RefundReasonDuplicate
	ReasonFraudulent          = inttegro.RefundReasonFraudulent
	ReasonOrderCanceled       = inttegro.RefundReasonOrderCanceled
	ReasonItemReturned        = inttegro.RefundReasonItemReturned
	ReasonItemDamaged         = inttegro.RefundReasonItemDamaged
	ReasonItemNotReceived     = inttegro.RefundReasonItemNotReceived
	ReasonItemNotAsDescribed  = inttegro.RefundReasonItemNotAsDescribed
	ReasonCustom              = inttegro.RefundReasonCustom

	StatusCanceled   = inttegro.RefundStatusCanceled
	StatusFailed     = inttegro.RefundStatusFailed
	StatusPending    = inttegro.RefundStatusPending
	StatusProcessing = inttegro.RefundStatusProcessing
	StatusSucceeded  = inttegro.RefundStatusSucceeded
)
