// Package payment provides payment resources and operations.
package payment

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/balancetransaction"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payout"
)

type NextActionType string

const (
	NextActionTypeConfirmPayment NextActionType = "confirm_payment"
	NextActionTypeExecute        NextActionType = "execute"
	NextActionTypeRedirect       NextActionType = "redirect"
	NextActionTypeAuthorize      NextActionType = "authorize"
	NextActionTypeNone           NextActionType = "none"
)

type ConfirmationChannel string

const (
	ConfirmationChannelSMS   ConfirmationChannel = "sms"
	ConfirmationChannelEmail ConfirmationChannel = "email"
	ConfirmationChannelPush  ConfirmationChannel = "push"
)

type Status string

const (
	StatusInitiated      Status = "initiated"
	StatusRequiresAction Status = "requires_action"
	StatusOverdue        Status = "overdue"
	StatusExecuted       Status = "executed"
	StatusPaid           Status = "paid"
	StatusCanceled       Status = "canceled"
	StatusExpired        Status = "expired"
	StatusFailed         Status = "failed"
	StatusUnknown        Status = "unknown"
)

type AttemptStatus string

const (
	AttemptStatusInitiated AttemptStatus = "initiated"
	AttemptStatusExecuted  AttemptStatus = "executed"
	AttemptStatusSucceeded AttemptStatus = "succeeded"
	AttemptStatusCanceled  AttemptStatus = "canceled"
	AttemptStatusExpired   AttemptStatus = "expired"
	AttemptStatusFailed    AttemptStatus = "failed"
	AttemptStatusUnknown   AttemptStatus = "unknown"
)

type ResultStatus string

const (
	ResultStatusPending              ResultStatus = "pending"
	ResultStatusRequiresConfirmation ResultStatus = "requires_confirmation"
	ResultStatusProcessing           ResultStatus = "processing"
	ResultStatusSucceeded            ResultStatus = "succeeded"
	ResultStatusFailed               ResultStatus = "failed"
)

// PaymentAttempt captures details of a single payment attempt.
//
// Each payment may have multiple attempts if initial attempts fail.
// This tracks the most recent attempt's status and timing.
type Attempt struct {
	// PaymentMethodType is the type of payment method used.
	PaymentMethodType string `json:"payment_method_type,omitempty"`

	// PaymentMethodID is the ID of the payment method charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// Reference is the external transaction reference from the payment provider.
	// Use this when contacting support or investigating payment issues.
	Reference string `json:"reference,omitempty"`

	// Status is the attempt's current state.
	// Values: "initiated", "succeeded", "failed"
	Status AttemptStatus `json:"status,omitempty"`

	// InitiatedAt is when the attempt started (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SucceededAt is when the attempt succeeded (ISO 8601).
	// Nil if not yet succeeded.
	SucceededAt *string `json:"succeeded_at,omitempty"`

	// FailedAt is when the attempt failed (ISO 8601).
	// Nil if not yet failed.
	FailedAt *string `json:"failed_at,omitempty"`
}

// PaymentNextAction describes the next step required to complete payment.
//
// When a payment requires customer action (OTP confirmation, redirect to
// bank page, etc), this field describes what needs to happen.
//
// Check Type to determine the required action:
//   - "confirm_payment": Customer must provide OTP
//   - "redirect": Customer must visit a URL
//   - "execute": Internal processing (no action needed)
type NextAction struct {
	// Type specifies the action category.
	// Values: "confirm_payment", "redirect", "execute"
	Type NextActionType `json:"type"`

	// ConfirmPayment contains OTP confirmation details.
	// Only present when Type is "confirm_payment".
	ConfirmPayment *struct {
		// ExpiresAt is when the OTP expires (ISO 8601).
		// Customer must confirm before this time.
		ExpiresAt string `json:"expires_at"`

		// Scheme describes the confirmation method.
		// Example: "otp"
		Scheme string `json:"scheme,omitempty"`

		// Request contains OTP delivery details.
		Request *struct {
			// ID is the OTP request identifier.
			ID string `json:"id"`

			// Recipient is where the OTP was sent.
			// For SMS: the phone number.
			Recipient string `json:"recipient"`

			// SentVia is the delivery channel.
			// Values: "sms", "email"
			SentVia ConfirmationChannel `json:"sent_via"`

			// TokenSize is the number of OTP digits.
			// Typically 4 or 6.
			TokenSize int `json:"token_size"`

			// SenderID is the SMS sender name.
			SenderID string `json:"sender_id"`
		} `json:"request,omitempty"`
	} `json:"confirm_payment,omitempty"`

	// Execute is present when Type is "execute" (internal processing).
	Execute any `json:"execute,omitempty"`

	// Redirect contains redirect details when Type is "redirect".
	Redirect *struct {
		// URL is where the customer should be redirected.
		// Open this URL in a browser for the customer to complete payment.
		URL string `json:"url"`
	} `json:"redirect,omitempty"`
}

// Payment represents payment details and status for an order.
//
// Every paid order has an associated payment object tracking the charge
// lifecycle, attempts, and any required customer actions.
type Resource struct {
	// ID is the unique payment identifier.
	// Starts with "py_". Example: "py_abc123def456"
	ID string `json:"id,omitempty"`

	// Status is the payment's current state.
	// Values: "initiated", "requires_action", "processing", "paid", "failed"
	Status Status `json:"status,omitempty"`

	// StatementDescriptor is what appears on the customer's statement.
	// Maximum 22 characters.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// Amount is the charged amount.
	Amount *money.Amount `json:"amount,omitempty"`

	// PaymentMethod is the charged payment method details.
	PaymentMethod *paymentmethod.Resource `json:"payment_method,omitempty"`

	// LatestAttempt is the most recent payment attempt.
	// Nil if no attempts yet.
	LatestAttempt *Attempt `json:"latest_attempt,omitempty"`

	// NextAction describes any required customer action.
	// Nil if no action required (payment is processing or complete).
	NextAction *NextAction `json:"next_action,omitempty"`

	// BalanceTransaction is the resulting balance entry when payment succeeds.
	// Used for tracking payouts and available balance.
	BalanceTransaction *balancetransaction.Resource `json:"balance_transaction,omitempty"`

	// PayoutConfiguration is the payout setup used for this payment (if applicable).
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`

	// InitiatedAt is when payment was first attempted (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// ExecutedAt is when payment was submitted to the network (ISO 8601).
	// Nil if not yet executed.
	ExecutedAt *string `json:"executed_at,omitempty"`

	// PaidAt is when payment was confirmed successful (ISO 8601).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// FailedAt is when payment was marked failed (ISO 8601).
	// Nil if not failed.
	FailedAt *string `json:"failed_at,omitempty"`
}
