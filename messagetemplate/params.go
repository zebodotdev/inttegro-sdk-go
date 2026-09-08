package messagetemplate

type CreateParams struct {
	IdempotencyKey string        `json:"-"`
	About          string        `json:"about,omitempty"`
	Attachments    []string      `json:"attachments,omitempty"`
	Channel        Channel       `json:"channel"`
	Email          *EmailContent `json:"email,omitempty"`
	Locale         string        `json:"locale,omitempty"`
	Name           string        `json:"name"`
	Purpose        string        `json:"purpose"`
	SMS            *SMSContent   `json:"sms,omitempty"`
	Variables      []Variable    `json:"variables,omitempty"`
}

type UpdateParams struct {
	IdempotencyKey string        `json:"-"`
	ID             string        `json:"id"`
	About          string        `json:"about,omitempty"`
	Attachments    []string      `json:"attachments,omitempty"`
	Channel        Channel       `json:"channel,omitempty"`
	Email          *EmailContent `json:"email,omitempty"`
	Locale         string        `json:"locale,omitempty"`
	Name           string        `json:"name,omitempty"`
	Purpose        string        `json:"purpose,omitempty"`
	SMS            *SMSContent   `json:"sms,omitempty"`
	Variables      []Variable    `json:"variables,omitempty"`
}

type PageParams struct {
	Channel Channel `json:"channel,omitempty"`
	Locale  string  `json:"locale,omitempty"`
	Page    int     `json:"page,omitempty"`
	Purpose string  `json:"purpose,omitempty"`
	Size    int     `json:"size,omitempty"`
	Status  Status  `json:"status,omitempty"`
}

type RenderPreviewParams struct {
	MessageTemplate Reference `json:"message_template"`
}
