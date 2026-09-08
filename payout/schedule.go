package payout

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
