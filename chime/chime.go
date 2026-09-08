// Package chime provides transactional notification resources and operations.
package chime

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service         = inttegro.ChimesService
	RecipientType   = inttegro.ChimeRecipientType
	Transport       = inttegro.ChimeTransport
	EmailSchemaKind = inttegro.ChimeEmailSchemaKind
	Recipient       = inttegro.ChimeRecipient
	SendParams      = inttegro.SendChimeParams
	ScheduleParams  = inttegro.ScheduleChimeParams
	BroadcastParams = inttegro.BroadcastChimeParams
	LookupParams    = inttegro.LookupChimeParams
	Resource        = inttegro.Chime
	PageParams      = inttegro.ChimePageParams
	Page            = inttegro.ChimesPage
)

const (
	RecipientTypePhone = inttegro.ChimeRecipientTypePhone
	RecipientTypeEmail = inttegro.ChimeRecipientTypeEmail

	TransportSMS   = inttegro.ChimeTransportSMS
	TransportEmail = inttegro.ChimeTransportEmail

	EmailSchemaKindGmailViewAction  = inttegro.ChimeEmailSchemaKindGmailViewAction
	EmailSchemaKindSchemaOrgOrder   = inttegro.ChimeEmailSchemaKindSchemaOrgOrder
	EmailSchemaKindSchemaOrgInvoice = inttegro.ChimeEmailSchemaKindSchemaOrgInvoice
)
