package messagetemplate

type Channel string

const (
	ChannelSMS   Channel = "sms"
	ChannelEmail Channel = "email"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type VariableType string

const (
	VariableTypeString   VariableType = "string"
	VariableTypeNumber   VariableType = "number"
	VariableTypeInteger  VariableType = "integer"
	VariableTypeBoolean  VariableType = "boolean"
	VariableTypeURL      VariableType = "url"
	VariableTypeEmail    VariableType = "email"
	VariableTypePhone    VariableType = "phone"
	VariableTypeDate     VariableType = "date"
	VariableTypeDatetime VariableType = "datetime"
	VariableTypeArray    VariableType = "array"
)

type VariableItemType string

const (
	VariableItemTypeString   VariableItemType = "string"
	VariableItemTypeNumber   VariableItemType = "number"
	VariableItemTypeInteger  VariableItemType = "integer"
	VariableItemTypeBoolean  VariableItemType = "boolean"
	VariableItemTypeURL      VariableItemType = "url"
	VariableItemTypeEmail    VariableItemType = "email"
	VariableItemTypePhone    VariableItemType = "phone"
	VariableItemTypeDate     VariableItemType = "date"
	VariableItemTypeDatetime VariableItemType = "datetime"
)

type ContentSafetyStatus string

const (
	ContentSafetyStatusAllowed     ContentSafetyStatus = "allowed"
	ContentSafetyStatusRejected    ContentSafetyStatus = "rejected"
	ContentSafetyStatusQuarantined ContentSafetyStatus = "quarantined"
)
