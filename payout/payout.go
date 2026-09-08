// Package payout provides payout resources, settings, and operations.
package payout

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service        = inttegro.PayoutsService
	Status         = inttegro.PayoutStatus
	ScheduleParams = inttegro.SchedulePayoutParams
	ScheduleSpec   = inttegro.PayoutScheduleSpec
	Schedule       = inttegro.PayoutSchedule
	Settings       = inttegro.PayoutSettings
	Resource       = inttegro.Payout
	Configuration  = inttegro.PayoutConfiguration
	Destination    = inttegro.PayoutDestination
	PageParams     = inttegro.PayoutPageParams
)

const (
	StatusInitialized = inttegro.PayoutStatusInitialized
	StatusScheduled   = inttegro.PayoutStatusScheduled
	StatusProcessing  = inttegro.PayoutStatusProcessing
	StatusExecuting   = inttegro.PayoutStatusExecuting
	StatusSucceeded   = inttegro.PayoutStatusSucceeded
	StatusInvalid     = inttegro.PayoutStatusInvalid
	StatusCanceled    = inttegro.PayoutStatusCanceled
)
