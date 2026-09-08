// Package schedule provides scheduled-chime resources and operations.
package schedule

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service        = inttegro.SchedulesService
	Detail         = inttegro.ScheduleDetail
	Error          = inttegro.ScheduleError
	Resource       = inttegro.ScheduleDetail
	ScheduledChime = inttegro.ScheduledChime
	LookupParams   = inttegro.LookupScheduleParams
	CancelParams   = inttegro.CancelScheduleParams
)
