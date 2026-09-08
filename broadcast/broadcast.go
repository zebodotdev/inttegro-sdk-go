// Package broadcast provides broadcast-chime resources and operations.
package broadcast

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service      = inttegro.BroadcastsService
	CreateParams = inttegro.BroadcastChimeParams
	Creation     = inttegro.BroadcastCreation
	Detail       = inttegro.BroadcastDetail
	Resource     = inttegro.BroadcastDetail
	Error        = inttegro.BroadcastError
	LookupParams = inttegro.LookupBroadcastParams
	CancelParams = inttegro.CancelBroadcastParams
)
