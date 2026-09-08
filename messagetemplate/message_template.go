// Package messagetemplate provides reusable message-template resources and operations.
package messagetemplate

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service             = inttegro.MessageTemplatesService
	Channel             = inttegro.MessageTemplateChannel
	Status              = inttegro.MessageTemplateStatus
	VariableType        = inttegro.MessageTemplateVariableType
	VariableItemType    = inttegro.MessageTemplateVariableItemType
	ContentSafetyStatus = inttegro.ContentSafetyStatus
	Variable            = inttegro.MessageTemplateVariable
	VariableItem        = inttegro.MessageTemplateVariableItem
	SMSContent          = inttegro.MessageTemplateSMSContent
	Mailbox             = inttegro.MessageTemplateMailbox
	EmailContent        = inttegro.MessageTemplateEmailContent
	Resource            = inttegro.MessageTemplate
	CreateParams        = inttegro.MessageTemplateCreateParams
	UpdateParams        = inttegro.MessageTemplateUpdateParams
	PageParams          = inttegro.MessageTemplatePageParams
	Page                = inttegro.MessageTemplatePage
	Reference           = inttegro.MessageTemplateReference
	RenderPreviewParams = inttegro.MessageTemplateRenderPreviewParams
	RenderedContent     = inttegro.MessageTemplateRenderedContent
	RenderPreviewOutput = inttegro.MessageTemplateRenderPreviewOutput
)

const (
	ChannelSMS   = inttegro.MessageTemplateChannelSMS
	ChannelEmail = inttegro.MessageTemplateChannelEmail

	StatusDraft     = inttegro.MessageTemplateStatusDraft
	StatusPublished = inttegro.MessageTemplateStatusPublished
	StatusArchived  = inttegro.MessageTemplateStatusArchived

	VariableTypeString   = inttegro.MessageTemplateVariableTypeString
	VariableTypeNumber   = inttegro.MessageTemplateVariableTypeNumber
	VariableTypeInteger  = inttegro.MessageTemplateVariableTypeInteger
	VariableTypeBoolean  = inttegro.MessageTemplateVariableTypeBoolean
	VariableTypeURL      = inttegro.MessageTemplateVariableTypeURL
	VariableTypeEmail    = inttegro.MessageTemplateVariableTypeEmail
	VariableTypePhone    = inttegro.MessageTemplateVariableTypePhone
	VariableTypeDate     = inttegro.MessageTemplateVariableTypeDate
	VariableTypeDatetime = inttegro.MessageTemplateVariableTypeDatetime
	VariableTypeArray    = inttegro.MessageTemplateVariableTypeArray

	VariableItemTypeString   = inttegro.MessageTemplateVariableItemTypeString
	VariableItemTypeNumber   = inttegro.MessageTemplateVariableItemTypeNumber
	VariableItemTypeInteger  = inttegro.MessageTemplateVariableItemTypeInteger
	VariableItemTypeBoolean  = inttegro.MessageTemplateVariableItemTypeBoolean
	VariableItemTypeURL      = inttegro.MessageTemplateVariableItemTypeURL
	VariableItemTypeEmail    = inttegro.MessageTemplateVariableItemTypeEmail
	VariableItemTypePhone    = inttegro.MessageTemplateVariableItemTypePhone
	VariableItemTypeDate     = inttegro.MessageTemplateVariableItemTypeDate
	VariableItemTypeDatetime = inttegro.MessageTemplateVariableItemTypeDatetime

	ContentSafetyStatusAllowed     = inttegro.ContentSafetyStatusAllowed
	ContentSafetyStatusRejected    = inttegro.ContentSafetyStatusRejected
	ContentSafetyStatusQuarantined = inttegro.ContentSafetyStatusQuarantined
)
