// Package payout provides payout resources and operations.
package payout

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
)

type Status string

const (
	StatusInitialized Status = "initialized"
)

const (
	StatusScheduled Status = "scheduled"
)

const (
	StatusProcessing Status = "processing"
)

const (
	StatusExecuting Status = "executing"
)

const (
	StatusSucceeded Status = "succeeded"
)

const (
	StatusInvalid Status = "invalid"
)

const (
	StatusCanceled Status = "canceled"
)

// PayoutsService manages payout configuration, scheduling, and history.
//
// Payouts move funds from your Inttegro balance to your connected financial
// accounts (mobile money, bank, Dosh). Use this service to:
//
//   - Configure payout destination accounts
//   - View payout settings and schedule
//   - Switch between automatic and manual payout modes
//   - Enable/disable currency conversion
//   - List payout history
//
// Example:
//
//	// Configure payout destinations
//	settings, err := client.Payouts.SetDestinations(ctx, map[string]string{
//	    "ghs": "fa_abc123",  // GHS payouts go to this mobile money account
//	    "usd": "fa_def456",  // USD payouts go to this bank account
//	})
//
// Learn more: https://studio.inttegro.com/set-up-payouts
type Service struct {
	client transport.Client
}

type ScheduleParams struct {
	DestinationID string `json:"destination_id"`
	ExecuteAfter  string `json:"execute_after,omitempty"`
	MaxAmount     int64  `json:"max_amount"`
	Reference     string `json:"reference"`
}

// SetDestinations configures which financial accounts receive payouts by currency.
//
// Map each currency you accept to a financial account ID. When balance
// transactions in that currency become eligible, they're paid out to the
// corresponding account.
//
// Parameters:
//   - destinations: Map of currency code to financial account ID
//     Example: {"ghs": "fa_abc123", "usd": "fa_def456"}
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.SetDestinations(ctx, map[string]string{
//	    "ghs": "fa_abc123",
//	})
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("GHS destination: %s\n", settings.Destinations["ghs"])
//
// Learn more: https://studio.inttegro.com/set-payout-destinations
func (s *Service) SetDestinations(ctx context.Context, destinations map[string]string) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	payload := struct {
		Destinations map[string]string `json:"destinations"`
	}{Destinations: destinations}
	if err := s.client.Do(ctx, "POST", "/payouts/set_destinations", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Settings retrieves your current payout configuration.
//
// Returns payout schedule, destination accounts, and FX settings.
//
// Example:
//
//	settings, err := client.Payouts.Settings(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Schedule: %s\n", settings.Schedule.Type)
//	fmt.Printf("Destinations: %v\n", settings.Destinations)
func (s *Service) Settings(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/settings", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Schedule creates a payout to a connected financial account.
func (s *Service) Schedule(ctx context.Context, params ScheduleParams) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/schedule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// Lookup retrieves a payout by ID.
func (s *Service) Lookup(ctx context.Context, payoutID string) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/lookup", map[string]string{"payout_id": payoutID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// DisableAutomatic switches to manual payout mode.
//
// In manual mode, payouts only happen when you explicitly request them.
// Use this for marketplace platforms or when you need precise control
// over payout timing.
//
// Returns the updated payout settings with manual schedule.
//
// Example:
//
//	settings, err := client.Payouts.DisableAutomatic(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Schedule type: %s\n", settings.Schedule.Type) // "manual"
//
// Learn more: https://studio.inttegro.com/disable-automatic-payouts
func (s *Service) DisableAutomatic(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/disable", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// EnableAutomatic switches payout scheduling back to automatic mode.
func (s *Service) EnableAutomatic(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/enable", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// EnableFX enables currency conversion for payouts.
//
// When enabled, you can receive payouts in a different currency than
// your source funds. Requires FX-enabled destination accounts.
//
// Example: Accept USD payments, receive GHS payouts.
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.EnableFX(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("FX enabled: %v\n", settings.FxEnabled) // true
func (s *Service) EnableFX(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/enable_fx", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// DisableFX disables currency conversion for payouts.
//
// When disabled, you can only receive payouts in the same currency as
// your source funds.
//
// Returns the updated payout settings.
//
// Example:
//
//	settings, err := client.Payouts.DisableFX(ctx)
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("FX enabled: %v\n", settings.FxEnabled) // false
func (s *Service) DisableFX(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/disable_fx", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// Cancel cancels a scheduled payout before execution.
//
// Only payouts in "scheduled" status with a future execution time can be
// canceled. Once canceled, included balance transactions remain available for
// future payouts.
//
// Parameters:
//   - payoutID: Scheduled payout ID to cancel
//
// Returns the canceled payout object.
func (s *Service) Cancel(ctx context.Context, payoutID string) (*Resource, error) {
	var resp struct {
		Payout Resource `json:"payout"`
	}
	payload := struct {
		PayoutID string `json:"payout_id"`
	}{PayoutID: payoutID}
	if err := s.client.Do(ctx, "POST", "/payouts/cancel", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Payout, nil
}

// Page returns a paginated list of recent payouts.
//
// View payout history including amounts, statuses, and timing. Results
// are sorted by initiation date (newest first).
//
// Parameters:
//   - params.PageNumber: Page to retrieve (optional, default: 1)
//   - params.PageSize: Payouts per page (optional, default: 20, max: 100)
//
// Returns a slice of payouts for the requested page.
//
// Example:
//
//	payouts, err := client.Payouts.Page(ctx, payout.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	})
//	for _, payout := range payouts {
//	    fmt.Printf("Payout %s: %s %d %s\n",
//	        payout.ID, payout.Amount.Currency,
//	        payout.Amount.Value, payout.Status)
//	}
func (s *Service) Page(ctx context.Context, params PageParams) ([]Resource, error) {
	var resp struct {
		Page struct {
			Payouts []Resource `json:"payouts"`
		} `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/payouts/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Payouts, nil
}

// PayoutScheduleSpec describes the balance transaction aging period.
//
// Balance transactions must age before becoming eligible for payout.
// This protects against late-arriving disputes and chargebacks.
type ScheduleSpec struct {
	// ID is the spec identifier.
	ID string `json:"id,omitempty"`

	// TPlus indicates the aging period in days.
	// Example: "t+7" means 7 days after transaction.
	TPlus string `json:"t_plus,omitempty"`

	// Label is a human-readable description.
	Label string `json:"label,omitempty"`

	// Abide is the formal aging rule specification.
	Abide string `json:"abide,omitempty"`
}

// PayoutSchedule describes your payout timing configuration.
//
// Payouts can be automatic (weekly, daily) or manual (on-demand).
// Automatic schedules trigger payouts at regular intervals for eligible
// balance transactions. Manual mode requires explicit payout initiation.
type Schedule struct {
	// ID is the schedule identifier (read-only).
	ID string `json:"id,omitempty"`

	// Name is the schedule's display name (read-only).
	// Example: "Weekly Automatic", "Manual"
	Name string `json:"name,omitempty"`

	// Type indicates automatic or manual mode (read-only).
	// Values: "automatic", "manual"
	Type string `json:"type,omitempty"`

	// Interval is the payout frequency for automatic schedules (read-only).
	// Values: "weekly", "daily", nil (for manual)
	Interval string `json:"interval,omitempty"`

	// ScheduleOn specifies when automatic payouts run (read-only).
	// For weekly: day of week (e.g., "monday")
	// For daily: time of day
	ScheduleOn string `json:"schedule_on,omitempty"`

	// Description explains the schedule behavior (read-only).
	Description string `json:"description,omitempty"`

	// Spec contains the balance transaction aging rules (read-only).
	Spec *ScheduleSpec `json:"spec,omitempty"`
}

// PayoutSettings contains your complete payout configuration.
//
// Controls when payouts happen, where funds go, and whether currency
// conversion is enabled.
type Settings struct {
	// ID is the settings identifier (read-only).
	ID string `json:"id,omitempty"`

	// FxEnabled indicates whether currency conversion is enabled (read-only).
	// When true, can receive payouts in different currency than source funds.
	// Requires FX-enabled destination accounts.
	FxEnabled bool `json:"fx_enabled,omitempty"`

	// Destinations maps currencies to financial account IDs.
	// Key: currency code (e.g., "ghs", "usd")
	// Value: financial account ID (e.g., "fa_abc123")
	// Example: {"ghs": "fa_abc123", "usd": "fa_def456"}
	Destinations map[string]string `json:"destinations,omitempty"`

	// Schedule describes payout timing and frequency.
	Schedule *Schedule `json:"schedule,omitempty"`
}

// Payout represents a settlement transfer to your bank or mobile money account.
//
// Payouts move funds from your Inttegro balance to your financial accounts.
// Each payout contains one or more balance transactions that have aged
// past the dispute window.
type Resource struct {
	// ID is the unique payout identifier (read-only).
	// Starts with "po_". Example: "po_abc123def456"
	ID string `json:"id,omitempty"`

	// ApplicationID is your application's ID (read-only).
	ApplicationID string `json:"application_id,omitempty"`

	// DestinationID is the receiving financial account's ID (read-only).
	// Corresponds to a financial account you've connected.
	DestinationID string `json:"destination_id,omitempty"`

	// Amount is the payout total (read-only).
	Amount *money.Amount `json:"amount,omitempty"`

	// Status is the payout's current state (read-only).
	// Values include "scheduled", "initiated", "processing", "succeeded", "failed", "canceled"
	Status Status `json:"status,omitempty"`

	// InitiatedBy indicates who triggered the payout (read-only).
	// Values: "schedule" (automatic), "manual" (you initiated)
	InitiatedBy string `json:"initiated_by,omitempty"`

	// LatestAttemptID is the most recent execution attempt's ID (read-only).
	LatestAttemptID string `json:"latest_attempt_id,omitempty"`

	// LatestError contains error details if payout failed (read-only).
	// Nil if payout succeeded or is still processing.
	LatestError any `json:"latest_error,omitempty"`

	// InitiatedAt is when the payout was created (ISO 8601, read-only).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// ExecuteAfter is the scheduled execution timestamp for queued payouts (ISO 8601, read-only).
	// Nil for immediate/manual payouts that are not scheduled.
	ExecuteAfter *string `json:"execute_after,omitempty"`

	// ScheduledAt is when the payout was queued for execution (ISO 8601, read-only).
	// Nil when not scheduled.
	ScheduledAt *string `json:"scheduled_at,omitempty"`

	// CanceledAt is when a scheduled payout was canceled (ISO 8601, read-only).
	// Nil unless the payout has status "canceled".
	CanceledAt *string `json:"canceled_at,omitempty"`

	// MaxAmount is the maximum amount authorized for scheduled payouts (read-only).
	// This may differ from Amount when payout execution has not started.
	MaxAmount *money.Amount `json:"max_amount,omitempty"`

	// ExecutedAt is when the payout was submitted to the network (ISO 8601, read-only).
	// Nil if not yet executed.
	ExecutedAt *string `json:"executed_at,omitempty"`

	// ExpectedAt is when the payout should arrive (ISO 8601, read-only).
	// Estimate based on network speed. Actual arrival may vary.
	ExpectedAt *string `json:"expected_at,omitempty"`

	// SucceededAt is when the payout was confirmed (ISO 8601, read-only).
	// Nil if not yet succeeded.
	SucceededAt *string `json:"succeeded_at,omitempty"`

	// BalanceTransactionIDs lists the included balance transactions (read-only).
	// These are the source funds being paid out.
	BalanceTransactionIDs []string `json:"balance_transaction_ids,omitempty"`
}

// PayoutConfiguration describes payout routing and FX settings for a payment or balance transaction.
type Configuration struct {
	// EnableFX indicates whether FX conversion is enabled for this payout.
	EnableFX *bool `json:"enable_fx,omitempty"`

	// Destination specifies the financial account receiving the payout.
	Destination *Destination `json:"destination,omitempty"`
}

// PayoutDestination identifies the payout financial account.
type Destination struct {
	// FinancialAccountID is the ID of the destination financial account.
	FinancialAccountID string `json:"financial_account_id,omitempty"`
}

// PayoutPageParams specifies pagination for listing payouts.
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of payouts per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
