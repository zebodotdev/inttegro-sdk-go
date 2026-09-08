// Package payment provides payment resources and lifecycle values.
package payment

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Status              = inttegro.PaymentStatus
	AttemptStatus       = inttegro.PaymentAttemptStatus
	CheckoutStatus      = inttegro.CheckoutPaymentStatus
	ResultStatus        = inttegro.PaymentResultStatus
	NextActionType      = inttegro.PaymentNextActionType
	ConfirmationChannel = inttegro.PaymentConfirmationChannel
	Method              = inttegro.PaymentMethod
	Attempt             = inttegro.PaymentAttempt
	NextAction          = inttegro.PaymentNextAction
	Resource            = inttegro.Payment
)

const (
	StatusInitiated      = inttegro.PaymentStatusInitiated
	StatusRequiresAction = inttegro.PaymentStatusRequiresAction
	StatusOverdue        = inttegro.PaymentStatusOverdue
	StatusExecuted       = inttegro.PaymentStatusExecuted
	StatusPaid           = inttegro.PaymentStatusPaid
	StatusCanceled       = inttegro.PaymentStatusCanceled
	StatusExpired        = inttegro.PaymentStatusExpired
	StatusFailed         = inttegro.PaymentStatusFailed
	StatusUnknown        = inttegro.PaymentStatusUnknown

	AttemptStatusInitiated = inttegro.PaymentAttemptStatusInitiated
	AttemptStatusExecuted  = inttegro.PaymentAttemptStatusExecuted
	AttemptStatusSucceeded = inttegro.PaymentAttemptStatusSucceeded
	AttemptStatusCanceled  = inttegro.PaymentAttemptStatusCanceled
	AttemptStatusExpired   = inttegro.PaymentAttemptStatusExpired
	AttemptStatusFailed    = inttegro.PaymentAttemptStatusFailed
	AttemptStatusUnknown   = inttegro.PaymentAttemptStatusUnknown

	CheckoutStatusRequiresAction = inttegro.CheckoutPaymentStatusRequiresAction
	CheckoutStatusProcessing     = inttegro.CheckoutPaymentStatusProcessing
	CheckoutStatusSucceeded      = inttegro.CheckoutPaymentStatusSucceeded
	CheckoutStatusFailed         = inttegro.CheckoutPaymentStatusFailed
	CheckoutStatusCancelled      = inttegro.CheckoutPaymentStatusCancelled

	ResultStatusPending              = inttegro.PaymentResultStatusPending
	ResultStatusRequiresConfirmation = inttegro.PaymentResultStatusRequiresConfirmation
	ResultStatusProcessing           = inttegro.PaymentResultStatusProcessing
	ResultStatusSucceeded            = inttegro.PaymentResultStatusSucceeded
	ResultStatusFailed               = inttegro.PaymentResultStatusFailed

	NextActionTypeConfirmPayment = inttegro.PaymentNextActionTypeConfirmPayment
	NextActionTypeExecute        = inttegro.PaymentNextActionTypeExecute
	NextActionTypeRedirect       = inttegro.PaymentNextActionTypeRedirect
	NextActionTypeAuthorize      = inttegro.PaymentNextActionTypeAuthorize
	NextActionTypeNone           = inttegro.PaymentNextActionTypeNone

	ConfirmationChannelSMS   = inttegro.PaymentConfirmationChannelSMS
	ConfirmationChannelEmail = inttegro.PaymentConfirmationChannelEmail
	ConfirmationChannelPush  = inttegro.PaymentConfirmationChannelPush
)
